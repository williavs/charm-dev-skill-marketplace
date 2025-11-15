# Component Composition in Bubbletea

Build complex TUIs by composing reusable components.

## Core Principle

Bubbletea components are models that implement `Init()`, `Update()`, and `View()`. Compose them by embedding in parent models and delegating appropriately.

## Pattern: Embedded Components

```go
type model struct {
    textInput  textinput.Model  // Component
    table      table.Model      // Component
    spinner    spinner.Model    // Component
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Delegate to components
    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    cmds = append(cmds, cmd)

    m.table, cmd = m.table.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}

func (m model) View() string {
    return m.textInput.View() + "\n" + m.table.View()
}
```

## Pattern: Focus Management

```go
type focusable interface {
    Focus() tea.Cmd
    Blur()
}

type model struct {
    inputs     []textinput.Model
    focusIndex int
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case tea.KeyMsg:
        if msg.String() == "tab" {
            // Blur current
            m.inputs[m.focusIndex].Blur()

            // Move focus
            m.focusIndex = (m.focusIndex + 1) % len(m.inputs)

            // Focus next
            return m, m.inputs[m.focusIndex].Focus()
        }

    // Delegate to focused component only
    var cmd tea.Cmd
    m.inputs[m.focusIndex], cmd = m.inputs[m.focusIndex].Update(msg)
    return m, cmd
}
```

## Pattern: View Switching

```go
type model struct {
    state      viewState
    listModel  list.Model
    detailView detailModel
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch m.state {
    case listView:
        return m.updateList(msg)
    case detailView:
        return m.updateDetail(msg)
    }
}

func (m model) View() string {
    switch m.state {
    case listView:
        return m.listModel.View()
    case detailView:
        return m.detailView.View()
    }
}
```

## Pattern: Conditional Delegation

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle app-level keys
    if key, ok := msg.(tea.KeyMsg); ok {
        if key.String() == "ctrl+c" {
            return m, tea.Quit
        }
    }

    // Delegate based on state
    var cmd tea.Cmd
    if m.inputMode {
        m.textInput, cmd = m.textInput.Update(msg)
    } else {
        m.list, cmd = m.list.Update(msg)
    }

    return m, cmd
}
```

## Available Charm Components

**Input Components:**
- `textinput` - Single-line text input
- `textarea` - Multi-line text area
- `filepicker` - File selection

**Display Components:**
- `list` - Selectable list with filtering
- `table` - Tabular data display
- `viewport` - Scrollable content
- `paginator` - Page navigation
- `progress` - Progress bars
- `spinner` - Loading indicators

**Time Components:**
- `timer` - Countdown timer
- `stopwatch` - Elapsed time

**UI Helpers:**
- `help` - Help menu system
- `key` - Key binding management

## Creating Custom Components

```go
// Component must implement tea.Model interface
type customWidget struct {
    value  string
    cursor int
}

func (w customWidget) Init() tea.Cmd {
    return nil
}

func (w customWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Handle keys
    }
    return w, nil
}

func (w customWidget) View() string {
    return lipgloss.NewStyle().Render(w.value)
}
```

## Best Practices

**1. Always batch commands from multiple components**

```go
var cmds []tea.Cmd
m.input, cmd = m.input.Update(msg)
cmds = append(cmds, cmd)
m.table, cmd = m.table.Update(msg)
cmds = append(cmds, cmd)
return m, tea.Batch(cmds...)
```

**2. Initialize components in model constructor**

```go
func initialModel() model {
    ti := textinput.New()
    ti.Focus()

    return model{
        textInput: ti,
        spinner:   spinner.New(),
    }
}
```

**3. Size components based on window dimensions**

```go
case tea.WindowSizeMsg:
    m.list.SetSize(msg.Width, msg.Height-4)
    m.viewport.Width = msg.Width
    m.viewport.Height = msg.Height - 10
```

## Examples

- Multi-step form: `references/bubbletea/examples/credit-card-form/`
- View composition: `references/bubbletea/examples/composable-views/`
- Split editors: `references/bubbletea/examples/split-editors/`
- Tab navigation: `references/bubbletea/examples/tabs/`
