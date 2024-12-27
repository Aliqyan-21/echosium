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

// automodeCmd represents the autmode command
var automodeCmd = &cobra.Command{
	Use:   "automode",
	Short: "Changes music automatically according to two states of a developer - 'active coding' and 'thinking/reflecting'",
	Run: func(cmd *cobra.Command, args []string) {
		confFile, err := os.Open("config.json")
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

// observeKeyPress is a goroutine that checks if any key is presed and updates lastKeyPress time
func observeKeyPress(events chan hook.Event) {
	for e := range events {
		if e.Kind == hook.KeyDown {
			fmt.Println("A key was pressed")
			lastKeyPress = time.Now()
		}
	}
}

// changeState is the function that changes the states according to keypress or idle
func changeStates() {
	for {
		time.Sleep(30 * time.Second)

		elapsed := time.Since(lastKeyPress)

		if elapsed >= 20*time.Second && currState != "idle" {
			updateState("idle")
		} else if elapsed <= 1*time.Second && currState != "coding" { // time between key pressed should be less than a second
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

	currState = newState

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
	automodeCmd.Flags().StringVarP(&codingStateMood, "coding", "c", "focused", "Specify the mood for your coding state (e.g., focus, rock, energetic)")
	rootCmd.AddCommand(automodeCmd)
}
