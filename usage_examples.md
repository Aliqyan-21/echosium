# Echosium Usage Examples

## Basic Usage Examples

### 1. Auto Mode with Default Settings

```bash
./echosium automode
```

This will:

- Use "relaxed" music for idle state
- Use "focus" music for coding state
- Automatically switch between states based on your activity

### 2. Auto Mode with Custom Moods

```bash
./echosium automode -i ambient -c rock
```

This setup is perfect for:

- Ambient background music when thinking/planning
- Rock music when actively coding

### 3. Manual Mode with Different Moods

```bash
# Start with peaceful music
./echosium start -m peaceful

# Switch to energetic music
./echosium start -m energetic

# Try creative mood
./echosium start -m creative
```

## Advanced Usage Scenarios

### 1. Different Moods for Different Tasks

```bash
# For debugging sessions (high focus)
./echosium automode -i chill -c focus

# For creative coding sessions
./echosium automode -i ambient -c creative

# For documentation writing
./echosium automode -i peaceful -c relaxed
```

### 2. Project-Specific Configurations

Create different shell aliases for different types of work:

```bash
# Add to your .bashrc or .zshrc
alias echosium-debug='echosium automode -i chill -c focus'
alias echosium-create='echosium automode -i ambient -c creative'
alias echosium-docs='echosium automode -i peaceful -c relaxed'
```

## Configuration Examples

### 1. Basic config.json

```json
{
  "client_id": "your-jamendo-api-client-id"
  "idle_time": "15",
  "keypress_window" : "5",
  "min_key_presses" : "3"
}
```

### 2. Directory of config.json file

```bash
~/.config/echosium/config.json
```

## State Transition Examples

1. **Coding State Activation:**

   - Type code normally
   - After 3 keypresses within 5 seconds → Coding state music
   - Music changes to focus/coding mood

2. **Idle State Activation:**
   - Stop typing for 15 seconds
   - Automatic transition to idle state
   - Music changes to relaxed/idle mood

## Common Use Cases

1. **Long Coding Sessions:**

   ```bash
   ./echosium automode -i ambient -c focus
   ```

   - Ambient music during reading/thinking
   - Focus music during active coding
   - Automatic transitions based on activity

2. **Documentation Sprints:**

   ```bash
   ./echosium automode -i peaceful -c relaxed
   ```

   - Peaceful music during research
   - Relaxed music during writing
   - Gentle transitions for maintained flow

3. **Debug Sessions:**
   ```bash
   ./echosium automode -i chill -c energetic
   ```
   - Chill music during log reading
   - Energetic music during active debugging
   - Quick transitions for dynamic work

## Tips for Best Experience

1. **Mood Selection:**

   - Use calmer moods for idle state (`peaceful`, `ambient`, `chill`)
   - Use energetic moods for coding state (`focus`, `energetic`, `rock`)
   - Experiment with different combinations

2. **State Tuning:**

   - Default timings work well for most developers
   - 15-second idle threshold prevents frequent switching
   - 3 keypresses in 5 seconds accurately detects coding

3. **Music Flow:**
   - Let the automatic transitions work naturally
   - Focus on your work, let Echosium adapt to you
