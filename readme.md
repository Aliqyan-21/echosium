# Echosium

<p align="center">
  <img src="assets/logo.png" alt="Echosium Logo" width="250">
</p>

<p align="center">
  <em>Your coding companion, weaving productivity with sound.</em>
</p>

<p align="center">
  <a href="https://github.com/aliqyan-21/echosium/stargazers">
    <img src="https://img.shields.io/github/stars/aliqyan-21/echosium?style=flat-square&color=yellow" alt="Stars">
  </a>
  <a href="https://github.com/aliqyan-21/echosium/network/members">
    <img src="https://img.shields.io/github/forks/aliqyan-21/echosium?style=flat-square&color=orange" alt="Forks">
  </a>
  <a href="https://github.com/aliqyan-21/echosium/issues">
    <img src="https://img.shields.io/github/issues/aliqyan-21/echosium?style=flat-square&color=red" alt="Issues">
  </a>
  <a href="https://github.com/aliqyan-21/echosium/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/aliqyan-21/echosium?style=flat-square&color=blue" alt="License">
  </a>
  <br>
  <img src="https://img.shields.io/badge/Go-1.22.9-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go Version">
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey?style=flat-square" alt="Platform Support">
  <a href="https://goreportcard.com/report/github.com/aliqyan-21/echosium">
    <img src="https://goreportcard.com/badge/github.com/aliqyan-21/echosium?style=flat-square" alt="Go Report Card">
  </a>
  <br>
  <img src="https://img.shields.io/badge/maintenance-actively--developed-brightgreen.svg?style=flat-square" alt="Maintenance">
</p>

Echosium is an intelligent CLI music player that syncs with your natural development rhythm. By detecting your coding patterns, it automatically transitions between energizing tracks during active development and calming melodies during reflection phases, creating the perfect acoustic environment for your workflow.

## Demo

https://github.com/user-attachments/assets/52ac6986-668b-4465-9c66-338005dffed5

## Contents

- [Quick Start](#-quick-start)
- [Features](#-features)
- [Getting Started](#-getting-started)
- [Usage](#-usage)
- [How It Works](#how-it-works)
- [Technical Details](#technical-details)
- [Contributing](#-contributing)
- [License](#license)
- [Acknowledgments](#-acknowledgments)

## Features

- **Intelligent State Detection**: Automatically identifies your programming state.

  - Active coding: When you're in the flow of writing code
  - Reflection: During code review and problem-solving moments

- **Dynamic Music Adaptation**: Switches music based on your current state:

  - Coding state: Energetic, focus-enhancing tracks
  - Idle state: Relaxing, ambient tracks

- **Rich Music Library**: Access to vast library of tracks through Jamendo API

- **Zero-Config Setup**: Works right out of the box with smart defaults

- **Full Customization**: Personalize your preferences through config file, moods from terminal params

- **Developer-Centric CLI**: Clean, intuitive command-line interface

## Getting Started

### Prerequisites

- Go 1.22.9 or higher
- MPV player installed on your system
- [Jamendo API client ID](jamendo_key.md)

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

4. Install app:

```bash
go build
go install
```

5. Create a config.json file in the root directory:

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

## Example Moods

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

## How It Works

Echosium uses sophisticated state detection to determine your current activity:

- **Coding State**: Triggered when detecting 3 or more keypresses within a 5-second window (customizable)
- **Idle State**: Activated after 15 seconds of keyboard inactivity (customizable)
- **Music Transition**: Smooth transitions between states with appropriate mood-based tracks (customizable)

## Technical Details

- Built with Go
- Uses Jamendo API for music streaming
- MPV player for music playback
- Implements goroutines for efficient state management
- Uses mutex locks for thread-safe operations
- [Diagrams](technical_diagrams.md)

## 🤝 Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Discussing the current state of the code
- Submitting a fix
- Proposing new features

## License

This project is licensed under the Apache-2.0 License

## 🙏 Acknowledgments

- Powered by [Jamendo](https://www.jamendo.com/) API
- MPV player for audio playback
- All the amazing artists providing their music

---

Made with ❤️ for developers who code better with music.
