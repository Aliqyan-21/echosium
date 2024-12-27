package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/aliqyan-21/echosium/jamendo"
	hook "github.com/robotn/gohook"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

// initialization of important variables for better readablility and workflow
var (
	idleStateMood   string
	codingStateMood string
	currCmd         *exec.Cmd
	stateMutex      sync.Mutex
	lastKeyPress    time.Time
	idleTracks      []jamendo.Track
	codingTracks    []jamendo.Track
	currState       string
)

// initialization of some customizable constants in future
const checkInterval = 1 // every this seconds we will check states

var (
	idleThreshold             time.Duration = 15 // time without key presess to transition to idle state
	activeCodingWindow        time.Duration = 5  // active window to consider keypresses in
	minKeyPressForCodingState int           = 3  // min key pressed in activeCodingWindow to consider in coding state
)

// automodeCmd represents the autmode command
var automodeCmd = &cobra.Command{
	Use:   "automode",
	Short: "Changes music automatically according to two states of a developer - 'active coding' and 'thinking/reflecting'",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Failed to get the home directory: %v\n", err)
			return
		}

		confFile, err := os.Open(home + "/.config/echosium/config.json")
		if err != nil {
			fmt.Printf("Failed to open config file: %v\n", err)
			return
		}
		defer confFile.Close()

		var config map[string]string

		if err := json.NewDecoder(confFile).Decode(&config); err != nil {
			fmt.Printf("Failed to parse config: %v\n", err)
			return
		}

		clientID, exists := config["client_id"]
		if !exists {
			fmt.Printf("client_id not define\n")
			return
		} else if clientID == "" {
			fmt.Printf("client_id field is empty\n")
			return
		}

		// configuration loading
		tmp, exists := config["idle_time"]
		if exists && tmp != "" {
			conf, _ := strconv.Atoi(config["idle_time"])
			idleThreshold = time.Duration(conf)
		}

		tmp, exists = config["keypress_window"]
		if exists && tmp != "" {
			conf, _ := strconv.Atoi(config["keypress_window"])
			activeCodingWindow = time.Duration(conf)
		}

		tmp, exists = config["min_key_presses"]
		if exists && tmp != "" {
			conf, _ := strconv.Atoi(config["min_key_presses"])
			minKeyPressForCodingState = conf
		}

		// tracks for idle state mood
		fmt.Printf("Loading tracks for idle mode; Mood : %s\n", idleStateMood)
		idleTracks, err = jamendo.GetTracks(clientID, idleStateMood)
		if err != nil || len(idleTracks) == 0 {
			fmt.Printf("Failed to fetch the tracks for mood %s: %v\n", idleStateMood, err)
			return
		}

		// tracks for coding state mood
		fmt.Printf("Loading tracks for coding mode; Mood : %s\n", codingStateMood)
		codingTracks, err = jamendo.GetTracks(clientID, codingStateMood)
		if err != nil || len(codingTracks) == 0 {
			fmt.Printf("Failed to fetch the tracks for mood %s: %v\n", codingStateMood, err)
			return
		}

		lastKeyPress = time.Now()
		currState = "idle"
		fmt.Printf("Starting music in idle state\n")
		playTrack(idleTracks)

		events := hook.Start()
		defer hook.End()

		go observeKeyPress(events)
		changeStates()
	},
}

var keyPressCt int
var keyPressTimer time.Timer

// observeKeyPress is a goroutine that checks if any key is presed and updates lastKeyPress time
func observeKeyPress(events chan hook.Event) {
	for e := range events {
		if e.Kind == hook.KeyDown {
			stateMutex.Lock()
			// fmt.Println("A key was pressed")
			lastKeyPress = time.Now()
			keyPressCt++
			stateMutex.Unlock()

			if keyPressTimer.C != nil {
				keyPressTimer.Stop()
			}

			keyPressTimer = *time.AfterFunc(activeCodingWindow*time.Second, func() {
				stateMutex.Lock()
				keyPressCt = 0
				stateMutex.Unlock()
			})
		}
	}
}

// changeState is the function that changes the states according to keypress or idle
func changeStates() {
	ticker := time.NewTicker(checkInterval * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stateMutex.Lock()
		elapsed := time.Since(lastKeyPress)
		currKeyPressCt := keyPressCt
		stateMutex.Unlock()

		if elapsed >= idleThreshold*time.Second && currState != "idle" {
			updateState("idle")
		} else if elapsed < idleThreshold*time.Second && currKeyPressCt >= minKeyPressForCodingState && currState != "coding" { // time between key pressed should be less than a second
			updateState("coding")
		}
	}
}

// updateState function updates the currState and plays the tracks according to that states mood
func updateState(newState string) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if currState == newState {
		return
	}

	prevState := currState
	currState = newState

	fmt.Printf("\nState transition: %s -> %s at %s\n",
		prevState, currState, time.Now().Format("15:04:05"))

	if currState == "idle" {
		fmt.Printf("\nTransitioning to idle state; Mood : %s\n", idleStateMood)
		playTrack(idleTracks)
	} else {
		fmt.Printf("\nTransitioning to coding state; Mood : %s\n", codingStateMood)
		playTrack(codingTracks)
	}
}

// playTrack function plays the given track using mpv music player
func playTrack(tracks []jamendo.Track) {
	idx := rand.Intn(len(tracks)) // random int (0, length of tracks)
	track := tracks[idx]

	fmt.Printf("Now playing: %s by %s\n", track.Name, track.Artist)

	if currCmd != nil {
		_ = currCmd.Process.Kill()
		currCmd = nil
	}

	currCmd = exec.Command("mpv", "--no-audio-display", "--no-input-terminal", track.TrackUrl)
	currCmd.Stdout = os.Stdout
	currCmd.Stderr = os.Stderr

	go func() {
		err := currCmd.Run()

		if err != nil {
			fmt.Printf("Error occured while starting mpv: %v\n", err)
		}
	}()
}

// init initializes the command in cobra and sets the flags
func init() {
	automodeCmd.Flags().StringVarP(&idleStateMood, "idle", "i", "relaxed", "Specify the mood for your idle state (e.g., peaceful, relaxed, chill)")
	automodeCmd.Flags().StringVarP(&codingStateMood, "coding", "c", "focus", "Specify the mood for your coding state (e.g., focus, rock, energetic)")
	rootCmd.AddCommand(automodeCmd)
}
