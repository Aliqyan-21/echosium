# Echosium 🎵

<p align="center">
  <img src="assets/logo.png" alt="Echosium Logo" width="250">
</p>

<p align="center">
  <em>Your coding companion, weaving productivity with sound.</em>
</p>

Echosium is an intelligent CLI music player that automatically adapts to your coding rhythm based on two states of a developer ("active coding state" and "thinking/reflection state"").

Whether you're deep in focused coding or taking a moment to reflect, Echosium sets the perfect soundtrack for your development journey.

## ✨ Features

- 🎯 **Smart State Detection**: Automatically detects whether you're actively coding or in a thinking/reflection state.
- 🎵 **Adaptive Music Selection**: Switches music based on your current state:
  - Coding state: Energetic, focus-enhancing tracks
  - Idle state: Relaxing, ambient tunes
- 🎼 **Mood-Based Playlists**: Access to vast library of tracks through Jamendo API
- 🔧 **Customizable**: Set your preferred music moods for both coding and idle states
- 💻 **Developer-Friendly**: Simple CLI interface with intuitive commands

## 🚀 Getting Started

### Prerequisites

- Go 1.22.9 or higher
- MPV player installed on your system
- Jamendo API client ID

### Installation

1. Install MPV player:

```bash
# For Ubuntu/Debian
sudo apt-get install mpv

# For macOS
brew install mpv

# For Windows (using Chocolatey)
choco install mpv
```

2. Clone the repository:

```bash
git clone https://github.com/aliqyan-21/echosium.git
cd echosium
```

3. Install dependencies:

```bash
go mod download
```

4. Create a config.json file in the root directory:

- For Linux

```bash
~/.config/echosium/config.json
```

```json
{
  "client_id": "your-jamendo-api-client-id"
  "idle_time": "15",
  "keypress_window" : "5",
  "min_key_presses" : "3"
}
```

###### here:

> **idle_time** = time in seconds to consider the developer to be idle and change in idle state (coding -> idle)

> **keypress_window** = time in seconds to check for min_key_presses in that time window

> **min_key_presses** = min key presses in keypress_window time to change to coding state (idle -> coding)

5. Build the project:

```bash
go build
```

## 🎮 Usage

### Auto Mode

Automatically switches music based on your coding activity:

```bash
# Start with default moods (relaxed for idle, focus for coding)
./echosium automode

# Customize moods
./echosium automode -i peaceful -c energetic
```

Options:

- `-i, --idle`: Specify mood for idle state (default: "relaxed")
- `-c, --coding`: Specify mood for coding state (default: "focus")

### Manual Mode

Play music with a specific mood without automatic switching:

```bash
# Start with default mood
./echosium start

# Specify a mood
./echosium start -m energetic
```

Options:

- `-m, --mood`: Specify the mood for tracks (default: "relaxed")

## 🎵 Example Moods

You can use various moods like:

- relaxed
- focus
- energetic
- peaceful
- chill
- ambient
- creative
- upbeat

> More Examples - [usage](usage_examples.md)

## ⚙️ How It Works

Echosium uses sophisticated state detection to determine your current activity:

- **Coding State**: Triggered when detecting 3 or more keypresses within a 5-second window (customizable)
- **Idle State**: Activated after 15 seconds of keyboard inactivity (customizable)
- **Music Transition**: Smooth transitions between states with appropriate mood-based tracks (customizable)

## 🔍 Technical Details

- Built with Go
- Uses Jamendo API for music streaming
- MPV player for music playback
- Implements goroutines for efficient state management
- Uses mutex locks for thread-safe operations
- [Diagrams](technical_diagrams.md)

## 🤝 Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Suggest features
- Submit pull requests

## 📝 License

This project is licensed under the Apache-2.0 License

## 🙏 Acknowledgments

- Powered by [Jamendo](https://www.jamendo.com/) API
- MPV player for audio playback
- All the amazing artists providing their music

---

Made with ❤️ for developers who code better with music.
