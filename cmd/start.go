package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	"github.com/aliqyan-21/echosium/jamendo"
	"github.com/spf13/cobra"
)

var mood string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "plays the music based on the mood, default mood is focus",

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
		if !exists || clientID == "" {
			fmt.Println("Client ID not found in config.json.")
			return
		}

		fmt.Printf("Fetching tracks for mood: %s\n", mood)
		tracks, err := jamendo.GetTracks(clientID, mood)
		if err != nil {
			fmt.Printf("Failed to fetch tracks: %v\n", err)
			return
		}

		if len(tracks) == 0 {
			fmt.Printf("No tracks found for mood : %s.", mood)
		} else {
			randIdx := rand.Intn(len(tracks))
			track := tracks[randIdx]

			fmt.Printf("Now playing: %s by %s\n", track.Name, track.Artist)

			playMusic(track.TrackUrl)
		}
	},
}

func playMusic(trackUrl string) {
	_, err := exec.LookPath("mpv")
	if err != nil {
		fmt.Println("Error: mpv is not installed or not in your PATH.")
		return
	}

	cmd := exec.Command("mpv", "--no-audio-display", trackUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error playing music: %v\n", err)
	}
}

func init() {
	startCmd.Flags().StringVarP(&mood, "mood", "m", "relaxation", "Specify the mood for the music tracks (e.g., peaceful, upbeat, relaxed)")
	rootCmd.AddCommand(startCmd)
}
