# MVC Pattern in Bubbletea

Bubbletea implements The Elm Architecture (TEA), which maps cleanly to the Model-View-Controller (MVC) pattern familiar to many developers.

## MVC → TEA Mapping

```
MVC Pattern          TEA in Bubbletea
───────────          ────────────────
Model                Model (struct)
View                 View() function
Controller           Update() function
```

## The Model - Application State

The Model in Bubbletea serves the same role as the Model in MVC: it holds all application state.

```go
type model struct {
    // Domain data
    users      []User
    currentID  int

    // UI state
    cursor     int
    selected   map[int]struct{}

    // Component state
    textInput  textinput.Model
    spinner    spinner.Model

    // Application state
    loading    bool
    err        error
    quitting   bool
}
```

**Key differences from traditional MVC:**
- Immutable - Updates return a new model
- Complete - All state lives here (no hidden state)
- Serializable - Can be marshaled/unmarshaled

## The View - Rendering

The View function is a pure transformation: Model → String

```go
func (m model) View() string {
    if m.loading {
        return m.spinner.View() + " Loading...\n"
    }

    if m.err != nil {
        return fmt.Sprintf("Error: %v\n", m.err)
    }

    // Render based on model state
    var s strings.Builder

    s.WriteString("Users:\n\n")
    for i, user := range m.users {
        cursor := " "
        if i == m.cursor {
            cursor = ">"
        }
        s.WriteString(fmt.Sprintf("%s %s\n", cursor, user.Name))
    }

    return s.String()
}
```

**Key differences from traditional MVC:**
- Pure function - No side effects
- Returns string - Not HTML/templates
- Declarative - Describes what, not how

## The Controller - Update Logic

The Update function handles all user input and state transitions.

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        return m.handleKeyPress(msg)

    case usersLoadedMsg:
        m.users = msg.users
        m.loading = false
        return m, nil

    case errorMsg:
        m.err = msg.err
        m.loading = false
        return m, nil
    }

    return m, nil
}

func (m model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "q", "ctrl+c":
        m.quitting = true
        return m, tea.Quit

    case "up":
        if m.cursor > 0 {
            m.cursor--
        }

    case "down":
        if m.cursor < len(m.users)-1 {
            m.cursor++
        }

    case "enter":
        return m, m.loadUserDetails(m.users[m.cursor].ID)
    }

    return m, nil
}
```

**Key differences from traditional MVC:**
- Message-based - Not HTTP requests
- Returns commands - For async operations
- Single event loop - All messages processed sequentially

## Component Composition (Sub-MVCs)

Bubbletea models can contain other models, similar to nested MVC structures.

```go
type model struct {
    // Sub-components (each is its own MVC)
    textInput  textinput.Model  // Has its own Update/View
    spinner    spinner.Model
    table      table.Model

    // Application state
    activeView string
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Delegate to sub-components
    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    cmds = append(cmds, cmd)

    m.table, cmd = m.table.Update(msg)
    cmds = append(cmds, cmd)

    // Handle app-level logic
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "tab" {
            m.activeView = m.nextView()
        }
    }

    return m, tea.Batch(cmds...)
}

func (m model) View() string {
    switch m.activeView {
    case "input":
        return m.textInput.View()
    case "table":
        return m.table.View()
    default:
        return "Unknown view"
    }
}
```

## Common MVC Patterns in Bubbletea

### 1. Loading Data (Controller)

```go
// In traditional MVC: Controller method fetches data
// In Bubbletea: Command fetches data, returns message

func fetchUsers() tea.Cmd {
    return func() tea.Msg {
        users, err := api.GetUsers()
        if err != nil {
            return errorMsg{err}
        }
        return usersLoadedMsg{users}
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case usersLoadedMsg:
        m.users = msg.users
        m.loading = false
        return m, nil
}
```

### 2. Form Handling

```go
// Model
type model struct {
    inputs     []textinput.Model
    focusIndex int
}

// View
func (m model) View() string {
    var s strings.Builder
    for i, input := range m.inputs {
        s.WriteString(input.View())
        if i < len(m.inputs)-1 {
            s.WriteString("\n")
        }
    }
    return s.String()
}

// Controller
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case tea.KeyMsg:
        switch msg.String() {
        case "tab":
            m.focusIndex = (m.focusIndex + 1) % len(m.inputs)
            return m, m.updateFocus()

        case "enter":
            if m.allValid() {
                return m, m.submitForm()
            }
        }

    // Delegate to active input
    var cmd tea.Cmd
    m.inputs[m.focusIndex], cmd = m.inputs[m.focusIndex].Update(msg)
    return m, cmd
}
```

### 3. Navigation (Multiple Views)

```go
type viewState int

const (
    listView viewState = iota
    detailView
    editView
)

type model struct {
    currentView viewState

    // View-specific state
    list       []Item
    selected   *Item
}

func (m model) View() string {
    switch m.currentView {
    case listView:
        return m.renderList()
    case detailView:
        return m.renderDetail()
    case editView:
        return m.renderEdit()
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            if m.currentView == listView {
                m.selected = &m.list[m.cursor]
                m.currentView = detailView
            }
        case "e":
            if m.currentView == detailView {
                m.currentView = editView
            }
        case "esc":
            m.currentView = listView
        }
}
```

## Best Practices

### Keep Model Pure

```go
// Good - Pure data
type model struct {
    users []User
    cursor int
}

// Bad - Side effects
type model struct {
    db *sql.DB  // Don't store connections
    file *os.File  // Don't store file handles
}
```

### Keep View Pure

```go
// Good - Pure rendering
func (m model) View() string {
    return fmt.Sprintf("Count: %d", m.count)
}

// Bad - Side effects in View
func (m model) View() string {
    m.count++  // Don't mutate state
    log.Println("Rendering")  // Don't do I/O
    return fmt.Sprintf("Count: %d", m.count)
}
```

### Keep Update Pure (Use Commands for Side Effects)

```go
// Good - Returns command
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case saveMsg:
        return m, saveToFile(m.data)  // Command handles I/O
}

func saveToFile(data string) tea.Cmd {
    return func() tea.Msg {
        err := os.WriteFile("data.txt", []byte(data), 0644)
        if err != nil {
            return errorMsg{err}
        }
        return savedMsg{}
    }
}

// Bad - Side effects in Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case saveMsg:
        os.WriteFile("data.txt", []byte(m.data), 0644)  // Don't do I/O here
        return m, nil
}
```

## Comparison: Traditional MVC vs Bubbletea

| Aspect | Traditional MVC | Bubbletea |
|--------|----------------|-----------|
| Request/Event | HTTP Request | tea.Msg |
| Controller | Method on Controller class | Update function |
| Model Mutation | Direct mutation | Return new model |
| View Rendering | Template engine | View function |
| Async Operations | Callbacks/Promises | Commands |
| State Management | Distributed across components | Centralized in Model |
| Testing | Mock HTTP/DB | Pure functions |

## Migration Guide: From MVC to Bubbletea

If you're coming from traditional MVC:

1. **Model**: Move all state into a single struct
2. **View**: Convert templates into View() function
3. **Controller**: Split into Update() (synchronous) and Commands (async)
4. **Routes**: Become message types
5. **Middleware**: Becomes message filtering or Update logic
6. **Sessions**: Store in Model
7. **Database**: Access via Commands, not directly

## References

- ELM Architecture: `references/architecture/ELM-ARCHITECTURE.md`
- Component Patterns: `references/architecture/COMPONENTS.md`
- Bubbletea Examples: `references/bubbletea/examples/`
