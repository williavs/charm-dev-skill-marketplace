# Bubble Tea API Features Documentation

Comprehensive reference for advanced Bubble Tea features with official documentation links, code examples, and best practices from the Charm team.

**Research Date:** October 4, 2025
**Bubble Tea Versions Covered:** v1.x and v2.0 Beta

---

## Table of Contents

1. [Program Control](#program-control)
2. [Message Filtering](#message-filtering)
3. [Terminal Suspension](#terminal-suspension)
4. [Bracketed Paste](#bracketed-paste)
5. [Focus Reporting](#focus-reporting)
6. [Window Management](#window-management)
7. [Performance Control](#performance-control)
8. [Environment Configuration](#environment-configuration)
9. [Testing with Teatest](#testing-with-teatest)
10. [Version Information](#version-information)

---

## Program Control

### tea.Quit vs tea.Interrupt

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea

#### tea.Quit()

Signals the program to exit cleanly.

```go
func Quit() Msg
```

**Returns:** `QuitMsg`

**Usage:**
```go
case tea.KeyMsg:
    if msg.String() == "q" {
        return m, tea.Quit
    }
```

**Version:** Available since v0.1.0

---

#### tea.Interrupt()

Signals the program to interrupt, typically when Ctrl+C is pressed. Returns `ErrInterrupted` from `Program.Run()`.

```go
func Interrupt() Msg
```

**Returns:** `InterruptMsg`

**Usage:**
```go
case tea.KeyMsg:
    if msg.Type == tea.KeyCtrlC {
        return m, tea.Interrupt
    }
```

**Best Practice:**
> "You can switch any tea.Quit command with tea.Interrupt for seamless interruptions."
> — Bubble Tea v1.3.0 Release Notes

**Version:** Introduced in v1.3.0 (refined interrupt handling)

**Documentation:** https://github.com/charmbracelet/bubbletea/releases/tag/v1.3.0

---

## Message Filtering

### tea.WithFilter()

Supplies an event filter invoked before Bubble Tea processes messages. Allows intercepting and potentially modifying or blocking messages.

```go
func WithFilter(filter func(Model, Msg) Msg) ProgramOption
```

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea#WithFilter

**Parameters:**
- `filter`: Function that receives the current model and incoming message
- Returns: The message to process (or `nil` to ignore)

**Example - Prevent Quit with Unsaved Changes:**
```go
func filter(m tea.Model, msg tea.Msg) tea.Msg {
    if _, ok := msg.(tea.QuitMsg); !ok {
        return msg
    }

    model := m.(myModel)
    if model.hasChanges {
        return nil  // Block the quit message
    }

    return msg
}

p := tea.NewProgram(Model{}, tea.WithFilter(filter))
```

**Use Cases:**
- Preventing shutdown with unsaved changes
- Message transformation and routing
- Conditional message blocking
- Cross-cutting concerns (logging, analytics)

**Version:** Merged April 17, 2023 via [PR #536](https://github.com/charmbracelet/bubbletea/pull/536)

---

## Terminal Suspension

### tea.Suspend() and tea.Resume

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea

#### tea.Suspend()

Signals the program should suspend, typically when Ctrl+Z is pressed.

```go
func Suspend() Msg
```

**Returns:** `SuspendMsg`

**Note:**
> "Since bubbletea puts the terminal in raw mode, it needs to be handled on a per-program basis."

---

#### ResumeMsg

Sent when a program resumes from suspend state.

```go
type ResumeMsg struct{}
```

**Example:**
```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.Type == tea.KeyCtrlZ {
            return m, tea.Suspend
        }
    case tea.ResumeMsg:
        // Do something when resuming
        m.status = "Resumed from suspend"
    }
    return m, nil
}
```

**Version:** Available since early v0.x releases

---

## Bracketed Paste

### v1.x Implementation

#### EnableBracketedPaste()

Tells the Bubble Tea program to accept bracketed paste input. Automatically disabled when program quits.

```go
func EnableBracketedPaste() Msg
```

#### DisableBracketedPaste()

Tells the Bubble Tea program to stop processing bracketed paste input.

```go
func DisableBracketedPaste() Msg
```

**v1 Example:**
```go
func (m model) Init() tea.Cmd {
    return tea.EnableBracketedPaste
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.Paste {
            // Handle pasted content in v1
            m.content += msg.String()
        }
    }
    return m, nil
}
```

---

### v2.0 Implementation

In Bubble Tea v2, paste events are sent as dedicated message types.

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea/v2

#### Message Types

```go
type PasteMsg []byte
type PasteStartMsg struct{}
type PasteEndMsg struct{}
```

#### WithoutBracketedPaste()

Program option to start with bracketed paste disabled.

```go
func WithoutBracketedPaste() ProgramOption
```

**v2 Example:**
```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.PasteStartMsg:
        // User started pasting
        m.isPasting = true
    case tea.PasteMsg:
        // Handle paste content
        m.text += string(msg)
    case tea.PasteEndMsg:
        // User stopped pasting
        m.isPasting = false
    }
    return m, nil
}
```

**Migration Note:**
> In v2, bracketed-paste has its own message type. Use `tea.PasteMsg` to match against paste events instead of checking `KeyMsg.Paste`.

**Version:** v2.0.0-alpha.1 and later

---

## Focus Reporting

### tea.WithReportFocus()

Enables reporting when the terminal gains and loses focus.

```go
func WithReportFocus() ProgramOption
```

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea#WithReportFocus

**Message Types:**
```go
type FocusMsg struct{}  // Terminal gained focus
type BlurMsg struct{}   // Terminal lost focus
```

**Example:**
```go
p := tea.NewProgram(model{}, tea.WithReportFocus())

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.FocusMsg:
        m.hasFocus = true
        m.statusBar = "Window focused"
    case tea.BlurMsg:
        m.hasFocus = false
        m.statusBar = "Window not focused"
    }
    return m, nil
}
```

**Important Notes:**
- Most terminals and multiplexers support focus reporting
- tmux requires configuration: `set-option -g focus-events on`
- Some terminals may not support this feature

**Version:** Introduced in v1.1.0 "Let's Focus"

**Release Notes:** https://github.com/charmbracelet/bubbletea/releases/tag/v1.1.0

---

## Window Management

### tea.WindowSizeMsg

Reports terminal size changes. Sent initially and on every resize.

```go
type WindowSizeMsg struct {
    Width  int
    Height int
}
```

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea#WindowSizeMsg

**Example:**
```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        // Update component sizes based on new dimensions
        m.viewport.Width = msg.Width
        m.viewport.Height = msg.Height - 5
    }
    return m, nil
}
```

**Best Practice:**
> "Your model should record the dimensions and use them when rendering to calculate the sizes of widgets."

---

### tea.SetWindowTitle()

Produces a command that sets the terminal title.

```go
func SetWindowTitle(title string) Cmd
```

**Example:**
```go
func (m model) Init() tea.Cmd {
    return tea.SetWindowTitle("My Awesome App")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if m.currentPage == "editor" {
        return m, tea.SetWindowTitle("Editor - " + m.fileName)
    }
    return m, nil
}
```

**Version:** Available in v1.x and v2.x

---

## Performance Control

### tea.WithMaxFPS()

Sets a custom maximum frames per second for rendering.

```go
func WithMaxFPS(fps int) ProgramOption
```

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea#WithMaxFPS

**Example:**
```go
// Limit to 30 FPS for battery saving
p := tea.NewProgram(model{}, tea.WithMaxFPS(30))

// Or increase for smoother animations
p := tea.NewProgram(model{}, tea.WithMaxFPS(120))
```

**Use Cases:**
- Battery saving on mobile/laptop (lower FPS)
- Smooth animations (higher FPS)
- Resource-constrained environments
- Matching monitor refresh rates

**Default:** Bubble Tea uses an adaptive frame rate by default

**Version:** Available in v1.x and v2.x

---

## Environment Configuration

### tea.WithEnvironment()

Sets environment variables for the program, useful in remote sessions (SSH).

```go
func WithEnvironment(env []string) ProgramOption
```

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/bubbletea#WithEnvironment

**Parameters:**
- `env`: String slice in `os.Environ()` format

**Example - SSH Session:**
```go
import "github.com/charmbracelet/ssh"

var sess ssh.Session
pty, _, _ := sess.Pty()
environ := append(sess.Environ(), "TERM="+pty.Term)
p := tea.NewProgram(model, tea.WithEnvironment(environ))
```

**Use Cases:**
- SSH session environment passthrough
- Custom TERM values for color support
- Integration with `tea.WithColorProfile()`
- Testing with controlled environments

**Default:** Uses `os.Environ()` if not set

**Version:** Added in PR #1063

---

## Testing with Teatest

### Package Overview

**Official Documentation:** https://pkg.go.dev/github.com/charmbracelet/x/exp/teatest

**Blog Post:** https://charm.land/blog/teatest/

Package teatest provides helper functions to test `tea.Model`s.

**Installation:**
```bash
go get github.com/charmbracelet/x/exp/teatest@latest
```

**Status:** Experimental work-in-progress library

**Import:**
```go
import "github.com/charmbracelet/x/exp/teatest"
```

---

### Core Testing Functions

#### NewTestModel()

Creates a new TestModel for testing Bubble Tea programs.

```go
tm := teatest.NewTestModel(
    t,
    model{},
    teatest.WithInitialTermSize(300, 100),
)
```

---

#### WaitFor()

Reads from output until a condition is met.

```go
teatest.WaitFor(
    t,
    tm.Output(),
    func(bts []byte) bool {
        return bytes.Contains(bts, []byte("expected text"))
    },
    teatest.WithCheckInterval(time.Millisecond*100),
    teatest.WithDuration(time.Second*3),
)
```

---

#### RequireEqualOutput()

Compares output against golden files.

```go
out, err := io.ReadAll(tm.FinalOutput(t))
teatest.RequireEqualOutput(t, out)
```

**Golden Files:**
- Use `-update` flag to create/update golden files
- Stored alongside test files

---

### Testing Patterns

#### Full Output Test

```go
func TestFullOutput(t *testing.T) {
    m := initialModel(time.Second)
    tm := teatest.NewTestModel(
        t, m,
        teatest.WithInitialTermSize(300, 100),
    )
    out, err := io.ReadAll(tm.FinalOutput(t))
    teatest.RequireEqualOutput(t, out)
}
```

---

#### Final Model State Test

```go
func TestFinalModel(t *testing.T) {
    tm := teatest.NewTestModel(
        t,
        initialModel(time.Second),
        teatest.WithInitialTermSize(300, 100),
    )
    fm := tm.FinalModel(t)

    // Assert on model state
    if fm.counter != expectedValue {
        t.Errorf("expected %d, got %d", expectedValue, fm.counter)
    }
}
```

---

#### Interaction Test

```go
func TestInteraction(t *testing.T) {
    tm := teatest.NewTestModel(
        t,
        initialModel(10*time.Second),
        teatest.WithInitialTermSize(300, 100),
    )

    // Wait for specific output
    teatest.WaitFor(
        t, tm.Output(),
        func(bts []byte) bool {
            return bytes.Contains(bts, []byte("ready"))
        },
        teatest.WithCheckInterval(time.Millisecond*100),
        teatest.WithDuration(time.Second*3),
    )

    // Send key press
    tm.Send(tea.KeyMsg{
        Type:  tea.KeyRunes,
        Runes: []rune("q"),
    })

    // Wait for program to finish
    tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))
}
```

---

### Available Options

- `WithInitialTermSize(width, height)` - Set terminal dimensions
- `WithFinalTimeout(duration)` - Set timeout for program completion
- `WithCheckInterval(duration)` - Configure waiting interval during tests

---

### Best Practices

1. **Test Full Output:** Use golden files for regression testing
2. **Test Model State:** Verify internal state after interactions
3. **Test Interactions:** Simulate user input and verify behavior
4. **Use Timeouts:** Always set reasonable timeouts to prevent hanging tests
5. **Isolate Tests:** Each test should be independent

**Repository:** https://github.com/charmbracelet/x/tree/main/exp/teatest

---

## Version Information

### Feature Introduction Timeline

| Feature | Version | Release Date | Notes |
|---------|---------|--------------|-------|
| `tea.Quit()` | v0.1.0 | Early 2020 | Core functionality |
| `tea.Suspend()` / `ResumeMsg` | v0.x | Early 2020 | Core functionality |
| `EnableBracketedPaste()` | v0.x | 2021 | Issue #404 tracking |
| `tea.SetWindowTitle()` | v1.x | - | Window management |
| `tea.WithFilter()` | v1.x | April 17, 2023 | PR #536 |
| `tea.WithReportFocus()` | v1.1.0 | - | "Let's Focus" release |
| `tea.Interrupt()` | v1.3.0 | - | Refined interrupt handling |
| `tea.WithMaxFPS()` | v1.x | - | Performance control |
| `tea.WithEnvironment()` | v1.x | - | PR #1063 |
| `PasteMsg` / v2 paste types | v2.0.0-alpha.1 | - | Breaking change from v1 |
| Mouse API improvements | v2.0.0-alpha.1 | - | Split into specific types |

---

### Current Versions

**Stable (v1.x):**
- Latest: v1.3.x series
- Documentation: https://pkg.go.dev/github.com/charmbracelet/bubbletea

**Beta (v2.x):**
- Latest: v2.0.0-beta.4
- Documentation: https://pkg.go.dev/github.com/charmbracelet/bubbletea/v2
- Install: `go get github.com/charmbracelet/bubbletea/v2@v2.0.0-beta.4`

**Breaking Changes in v2:**
- Bracketed paste uses dedicated message types (`PasteMsg`, `PasteStartMsg`, `PaseEndMsg`)
- Mouse events split into specific types (`MouseClickMsg`, `MouseReleaseMsg`, `MouseWheelMsg`, `MouseMotionMsg`)
- More intuitive key message handling

---

## Additional Resources

### Official Documentation

- **Main Repository:** https://github.com/charmbracelet/bubbletea
- **Go Packages (v1):** https://pkg.go.dev/github.com/charmbracelet/bubbletea
- **Go Packages (v2):** https://pkg.go.dev/github.com/charmbracelet/bubbletea/v2
- **Release Notes:** https://github.com/charmbracelet/bubbletea/releases

### Learning Resources

- **Official Examples:** https://github.com/charmbracelet/bubbletea/tree/main/examples
- **Teatest Guide:** https://charm.land/blog/teatest/
- **Building Bubble Tea Programs:** https://leg100.github.io/en/posts/building-bubbletea-programs/
- **Rapidly Building CLIs:** https://www.inngest.com/blog/interactive-clis-with-bubbletea

### Community

- **v2 Alpha Discussion:** https://github.com/charmbracelet/bubbletea/discussions/1156
- **v2 Release Checklist:** https://github.com/charmbracelet/bubbletea/issues/1386

---

## Usage Notes

### Terminal Compatibility

- **Focus Reporting:** Not supported by all terminals
- **Bracketed Paste:** Widely supported but not universal
- **Sixel Graphics:** Terminal-dependent (iTerm2, Kitty support varies)
- **Mouse Events:** Generally well-supported

### tmux Configuration

For full feature support in tmux:

```bash
# Enable focus events
set-option -g focus-events on

# Enable mouse support
set-option -g mouse on
```

---

## Common Program Options

```go
p := tea.NewProgram(
    model{},
    tea.WithAltScreen(),              // Use full terminal
    tea.WithMouseCellMotion(),        // Enable mouse support
    tea.WithReportFocus(),            // Enable focus events
    tea.WithFilter(filterFunc),       // Add message filter
    tea.WithMaxFPS(60),               // Set frame rate
    tea.WithEnvironment(customEnv),   // Custom environment
)
```

---

**Generated:** October 4, 2025
**Based on:** Official Bubble Tea documentation, source code, and release notes
**Maintained by:** texamples research
**Last Updated:** 2025-10-04
