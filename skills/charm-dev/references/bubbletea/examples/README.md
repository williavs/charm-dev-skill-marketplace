# Bubbletea Examples

Reference implementations for Bubbletea TUI patterns. Each example demonstrates specific capabilities and implementation approaches.

## Quick Start

All examples are runnable Go programs. To run any example:

```bash
cd examples/[example-name]
go run main.go
```

## Finding Examples

For contextual reference (use X when you need Y), see [CONTEXTUAL-INVENTORY.md](CONTEXTUAL-INVENTORY.md).

For original descriptions, see [obsolete/README.md.original](obsolete/README.md.original).

## Categories

### Input
- **Single field**: textinput
- **Multiple fields**: textinputs, credit-card-form, split-editors
- **Multi-line**: textarea, chat
- **File selection**: file-picker
- **Autocomplete**: autocomplete
- **Validation**: credit-card-form

### Display
- **Lists**: list-default, list-simple, list-fancy
- **Tables**: table, table-resize
- **Text content**: pager, glamour
- **Pagination**: paginator

### Progress and Status
- **Progress bars**: progress-static, progress-animated, progress-download
- **Spinners**: spinner, spinners
- **Timers**: timer, stopwatch
- **Installation flow**: package-manager

### Layout
- **Multiple views**: views, composable-views
- **Tabs**: tabs
- **Responsive**: window-size, table-resize

### Integration
- **HTTP**: http
- **External commands**: exec
- **Pipes**: pipe
- **Real-time events**: realtime
- **Daemon mode**: tui-daemon-combo

### Terminal Control
- **Screen modes**: altscreen-toggle, fullscreen
- **Mouse**: mouse
- **Focus events**: focus-blur
- **Suspend**: suspend
- **Window title**: set-window-title

### Advanced Patterns
- **Custom messages**: send-msg
- **Command sequencing**: sequence
- **Debouncing**: debounce
- **Quit protection**: prevent-quit
- **Help system**: help
- **Cell buffer**: cellbuffer

### Learning
- **Basic structure**: simple

## Running Examples

Each example directory contains:
- `main.go` - Primary implementation
- Supporting `.go` files as needed
- Some include GIF demonstrations

Dependencies are managed via the shared `go.mod` in the examples root.

## Documentation

- **CONTEXTUAL-INVENTORY.md** - Detailed capability reference with implementation patterns
- **INVENTORY.md** - Existing inventory documentation
- **obsolete/README.md.original** - Original prescriptive descriptions

## File Paths

All examples use the pattern:
```
examples/[example-name]/main.go
```

Multi-file examples include additional `.go` files in the same directory.
