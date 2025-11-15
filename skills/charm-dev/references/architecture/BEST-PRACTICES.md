# Bubbletea Best Practices

## Architecture

**Use TEA pattern correctly:**
- Model = all state
- View = pure function (Model → String)
- Update = pure function ((Model, Msg) → (Model, Cmd))
- Commands = side effects

**Keep functions pure:**
```go
// Good - Pure Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.counter++
    return m, saveData(m.counter)  // Command handles I/O
}

// Bad - Side effects in Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.counter++
    os.WriteFile("data.txt", ...)  // Don't do I/O here
    return m, nil
}
```

## State Management

**All state in Model:**
```go
type model struct {
    // Domain data
    users []User

    // UI state
    cursor int

    // Component state
    textInput textinput.Model

    // Flags
    loading bool
    err error
}
```

**Don't store connections/handles:**
```go
// Bad
type model struct {
    db *sql.DB      // Don't
    file *os.File   // Don't
}
```

## Message Handling

**Use typed messages:**
```go
type userLoadedMsg struct { users []User }
type errorMsg struct { err error }

// Not: interface{} or any
```

**Handle all message types:**
```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        return m.handleKeys(msg)
    case tea.WindowSizeMsg:
        return m.handleResize(msg)
    case userLoadedMsg:
        return m.handleUserLoaded(msg)
    }
    return m, nil  // Always return model
}
```

## Commands

**Batch concurrent operations:**
```go
func (m model) Init() tea.Cmd {
    return tea.Batch(
        fetchUsers(),
        fetchPosts(),
        tick(),
    )
}
```

**Sequence dependent operations:**
```go
case submitMsg:
    return m, tea.Sequence(
        saveData(m.data),
        showSuccess,
        tea.Quit,
    )
```

## Components

**Initialize in constructor:**
```go
func initialModel() model {
    ti := textinput.New()
    ti.Focus()
    ti.CharLimit = 50

    return model{textInput: ti}
}
```

**Delegate appropriately:**
```go
// Delegate to active component only
if m.inputMode {
    m.textInput, cmd = m.textInput.Update(msg)
} else {
    m.list, cmd = m.list.Update(msg)
}
```

## Rendering

**Use Lipgloss for all styling:**
```go
import "github.com/charmbracelet/lipgloss"

var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("170"))

    helpStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("241"))
)

func (m model) View() string {
    return titleStyle.Render("Title") + "\n" +
           m.content + "\n\n" +
           helpStyle.Render("Press q to quit")
}
```

**Handle window resize:**
```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
    m.viewport.Width = msg.Width
    m.viewport.Height = msg.Height - 4  // Reserve space for header/footer
```

## Error Handling

**Return errors as messages:**
```go
func fetchData() tea.Cmd {
    return func() tea.Msg {
        data, err := api.Get()
        if err != nil {
            return errorMsg{err}
        }
        return dataMsg{data}
    }
}
```

**Display errors in View:**
```go
func (m model) View() string {
    if m.err != nil {
        return fmt.Sprintf("Error: %v\n\nPress q to quit", m.err)
    }
    // Normal view
}
```

## Performance

**Limit FPS for battery savings:**
```go
p := tea.NewProgram(
    initialModel(),
    tea.WithMaxFPS(30),  // Default is adaptive
)
```

**Avoid expensive operations in View:**
```go
// Bad - Expensive in View
func (m model) View() string {
    users := fetchUsers()  // Don't
    sorted := sortExpensive(m.data)  // Don't
    return render(sorted)
}

// Good - Compute in Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case dataLoadedMsg:
        m.data = msg.data
        m.sorted = sort(m.data)  // Pre-compute
        return m, nil
}

func (m model) View() string {
    return render(m.sorted)  // Just use
}
```

## Common Program Options

```go
p := tea.NewProgram(
    initialModel(),
    tea.WithAltScreen(),           // Fullscreen mode
    tea.WithMouseCellMotion(),     // Mouse support
    tea.WithReportFocus(),         // Focus events
)
```

## Testing

Use `teatest` for testing TUIs:

```go
import "github.com/charmbracelet/x/exp/teatest"

func TestApp(t *testing.T) {
    tm := teatest.NewTestModel(
        t, initialModel(),
        teatest.WithInitialTermSize(80, 24),
    )

    // Send input
    tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})

    // Wait for quit
    tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))
}
```

## Anti-Patterns

**Don't mutate and return original:**
```go
// Bad
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, nil  // Returns pointer to mutated model
}

// Good
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.counter++
    return m, nil  // Returns value (immutable)
}
```

**Don't skip window size:**
```go
// Bad - Ignores resize
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        // handle keys
    }
    return m, nil
}

// Good - Handles resize
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width, m.height = msg.Width, msg.Height
    case tea.KeyMsg:
        // handle keys
    }
    return m, nil
}
```

## References

- TEA Architecture: `references/architecture/ELM-ARCHITECTURE.md`
- MVC Mapping: `references/architecture/MVC-IN-BUBBLETEA.md`
- Components: `references/architecture/COMPONENTS.md`
- Examples: `references/bubbletea/examples/CONTEXTUAL-INVENTORY.md`
