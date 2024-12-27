package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/aliqyan-21/echosium/jamendo"
	"github.com/spf13/cobra"
)

// initialization of mood of song variable
var mood string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "plays the music based on the mood, default mood is focus",

	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Failed to get the home directory: %v\n", err)
			return
		}

		confFile, err := os.Open(home + "/.config/echosium/config.json")
		if err != nil {
			fmt.Printf("Failed to find config file:\n %v\n", err)
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
			for true {
				randIdx := rand.Intn(len(tracks))
				track := tracks[randIdx]

				const boxWidth = 50
				padding := func(content string) string {
					if len(content) > boxWidth-15 {
						return content[:boxWidth-18] + "..."
					}
					return content + strings.Repeat(" ", boxWidth-len(content)-15)
				}

				title := fmt.Sprintf("| Title    : %-*s |", boxWidth-15, padding(track.Name))
				artist := fmt.Sprintf("| Artist   : %-*s |", boxWidth-15, padding(track.Artist))
				album := fmt.Sprintf("| Album    : %-*s |", boxWidth-15, padding(track.Album))
				mood := fmt.Sprintf("| State    : %-*s |", boxWidth-15, padding(mood))
				info1 := fmt.Sprintf("| Info     : %-*s |", boxWidth-15, padding("Press Space to pause/play"))
				info2 := fmt.Sprintf("|          : %-*s |", boxWidth-15, padding("Arrow Keys to forward/backward"))

				fmt.Println(strings.Repeat("-", boxWidth))
				fmt.Println("| Now Playing" + strings.Repeat(" ", boxWidth-14) + "|")
				fmt.Println(strings.Repeat("-", boxWidth))
				fmt.Println(title)
				fmt.Println(artist)
				fmt.Println(album)
				fmt.Println(mood)
				fmt.Println(strings.Repeat("-", boxWidth))
				fmt.Println(info1)
				fmt.Println(info2)
				fmt.Println(strings.Repeat("-", boxWidth))

				playMusic(track.TrackUrl)
			}
		}
	},
}

// playMucis function plays the given track using mpv music player
func playMusic(trackUrl string) {
	_, err := exec.LookPath("mpv")
	if err != nil {
		fmt.Println("Error: mpv is not installed or not in your PATH.")
		return
	}

	cmd := exec.Command("mpv", "--no-audio-display", "--msg-level=all=no", trackUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error playing music: %v\n", err)
	}
}

func init() {
	startCmd.Flags().StringVarP(&mood, "mood", "m", "focus", "Specify the mood for the music tracks (e.g., peaceful, upbeat, relaxed)")
	rootCmd.AddCommand(startCmd)
}
