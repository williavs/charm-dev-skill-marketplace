# The Elm Architecture (TEA) in Bubbletea

Bubbletea implements The Elm Architecture (TEA), a pattern for building interactive applications with a clear separation of concerns.

## Core Concept

The Elm Architecture follows a simple, predictable loop:

```
┌─────────────────────────────────────────────────┐
│                                                 │
│  ┌──────────┐      ┌──────────┐      ┌──────┐  │
│  │  Model   │─────▶│  View    │─────▶│ UI   │  │
│  └──────────┘      └──────────┘      └──────┘  │
│       ▲                                  │      │
│       │                                  │      │
│       │                                  ▼      │
│  ┌──────────┐                      ┌─────────┐ │
│  │  Update  │◀─────────────────────│   Msg   │ │
│  └──────────┘                      └─────────┘ │
│                                                 │
└─────────────────────────────────────────────────┘
```

## The Three Core Functions

### 1. Model - The Application State

The Model is a single source of truth for all application state.

```go
type model struct {
    // State
    counter     int
    textInput   textinput.Model
    choices     []string
    cursor      int
    selected    map[int]struct{}

    // Status
    quitting    bool
    err         error
}
```

**Key Principles:**
- Immutable - Updates return a new model
- Complete - Contains all state needed to render the view
- Serializable - Can be saved/loaded

### 2. View - Rendering the UI

The View is a pure function: `Model → String`

```go
func (m model) View() string {
    if m.quitting {
        return "Goodbye!\n"
    }

    // Build UI from model state
    s := fmt.Sprintf("Counter: %d\n\n", m.counter)
    s += m.textInput.View()

    return s
}
```

**Key Principles:**
- Pure function - Same model always produces same view
- No side effects - Only reads from model, doesn't modify
- Declarative - Describes what to show, not how to update

### 3. Update - State Transitions

The Update function handles all state changes: `(Model, Msg) → (Model, Cmd)`

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            m.quitting = true
            return m, tea.Quit

        case "up":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        }

    case tickMsg:
        m.counter++
        return m, tick() // Schedule next tick
    }

    return m, nil
}
```

**Key Principles:**
- Pure function - No side effects
- Returns new model - Doesn't mutate
- Returns commands - For async operations

## Messages (Msg)

Messages are the only way to trigger state changes. They can be:

### 1. User Input Messages
```go
case tea.KeyMsg:      // Keyboard input
case tea.MouseMsg:    // Mouse input
case tea.WindowSizeMsg: // Terminal resize
```

### 2. Custom Messages
```go
type tickMsg time.Time
type httpResponseMsg struct {
    data []byte
    err  error
}
type frameMsg image.Image
```

### 3. System Messages
```go
case tea.QuitMsg:     // Quit signal
case tea.FocusMsg:    // Window focus
case tea.BlurMsg:     // Window blur
```

## Commands (Cmd)

Commands represent side effects - things that happen asynchronously.

```go
// Command that waits and sends a message
func tick() tea.Cmd {
    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

// Command that makes an HTTP request
func fetchData(url string) tea.Cmd {
    return func() tea.Msg {
        resp, err := http.Get(url)
        if err != nil {
            return httpResponseMsg{err: err}
        }
        data, _ := io.ReadAll(resp.Body)
        return httpResponseMsg{data: data}
    }
}
```

**Command Patterns:**
- Return `nil` for no command
- Return `tea.Quit` to exit
- Return `tea.Batch(cmd1, cmd2)` for multiple commands
- Return `tea.Sequence(cmd1, cmd2)` for sequential execution

## Init - Bootstrapping

The Init function returns the initial command(s) to run when starting.

```go
func (m model) Init() tea.Cmd {
    return tea.Batch(
        tick(),                     // Start timer
        m.textInput.Focus(),        // Focus input
        fetchData("https://..."),   // Load data
    )
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "time"
    tea "github.com/charmbracelet/bubbletea"
)

// MODEL
type model struct {
    counter  int
    quitting bool
}

func initialModel() model {
    return model{counter: 0}
}

// MESSAGES
type tickMsg time.Time

// UPDATE
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            m.quitting = true
            return m, tea.Quit
        }

    case tickMsg:
        m.counter++
        return m, tick()
    }

    return m, nil
}

// VIEW
func (m model) View() string {
    if m.quitting {
        return "Goodbye!\n"
    }
    return fmt.Sprintf("Counter: %d\n\nPress q to quit\n", m.counter)
}

// INIT
func (m model) Init() tea.Cmd {
    return tick()
}

// COMMANDS
func tick() tea.Cmd {
    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
    }
}
```

## Why TEA?

### Benefits

1. **Predictable** - State changes only through Update
2. **Testable** - Pure functions are easy to test
3. **Debuggable** - Message flow is traceable
4. **Composable** - Models can contain other models
5. **Time-travel debugging** - Can replay messages

### Constraints

1. **No direct mutation** - Must return new model
2. **Async via commands** - Can't do side effects in Update
3. **Single event loop** - All messages processed sequentially

## Best Practices

1. **Keep Update pure** - No I/O, no side effects
2. **Use typed messages** - Easier to debug than `interface{}`
3. **Batch commands** - Use `tea.Batch()` for concurrent operations
4. **Handle all cases** - Every message type should be handled
5. **Return updated model** - Always return `m`, even if unchanged

## Common Patterns

### Delegating to Sub-Models

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    // Delegate to input component
    m.textInput, cmd = m.textInput.Update(msg)

    return m, cmd
}
```

### Chaining Commands

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case submitMsg:
        return m, tea.Sequence(
            saveData(m.data),
            showSuccess,
            tea.Quit,
        )
}
```

### Batching Commands

```go
func (m model) Init() tea.Cmd {
    return tea.Batch(
        fetchUsers(),
        fetchPosts(),
        tick(),
    )
}
```

## References

- Official Bubbletea Tutorial: https://github.com/charmbracelet/bubbletea/tree/main/tutorials
- The Elm Architecture Guide: https://guide.elm-lang.org/architecture/
- Bubbletea examples: `references/bubbletea/examples/`
