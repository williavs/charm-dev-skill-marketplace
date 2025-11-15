# Bubbletea Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|--------------|
| Installation progress tracking | `package-manager/main.go` |
| Multi-step form with validation | `credit-card-form/main.go` |
| Multi-line text input | `chat/main.go`, `textarea/main.go` |
| Single text input | `textinput/main.go` |
| Multiple text inputs with focus management | `textinputs/main.go`, `split-editors/main.go` |
| File selection interface | `file-picker/main.go` |
| Tabular data display | `table/main.go`, `table-resize/main.go` |
| Selectable list items | `list-default/main.go`, `list-fancy/main.go`, `list-simple/main.go` |
| Paginated content | `paginator/main.go` |
| Long-form text viewing | `pager/main.go` |
| Markdown rendering | `glamour/main.go` |
| Progress bars (animated) | `progress-animated/main.go` |
| Progress bars (static) | `progress-static/main.go` |
| Download progress | `progress-download/main.go` |
| Loading indicators | `spinner/main.go`, `spinners/main.go` |
| Time tracking | `timer/main.go`, `stopwatch/main.go` |
| Help menu system | `help/main.go` |
| HTTP requests | `http/main.go` |
| External command execution | `exec/main.go` |
| Real-time event handling | `realtime/main.go` |
| Multiple view switching | `views/main.go`, `composable-views/main.go` |
| Tab navigation | `tabs/main.go` |
| Stdin/stdout piping | `pipe/main.go` |
| Background daemon mode | `tui-daemon-combo/main.go` |
| Alt screen buffer toggle | `altscreen-toggle/main.go` |
| Fullscreen mode | `fullscreen/main.go` |
| Mouse input | `mouse/main.go` |
| Window resize handling | `window-size/main.go` |
| Focus/blur events | `focus-blur/main.go` |
| Prevent accidental quit | `prevent-quit/main.go` |
| Process suspension | `suspend/main.go` |
| Window title setting | `set-window-title/main.go` |
| Input debouncing | `debounce/main.go` |
| Custom message passing | `send-msg/main.go` |
| Sequential command execution | `sequence/main.go` |
| Choice menu | `result/main.go` |
| Autocomplete | `autocomplete/main.go` |
| Cell buffer rendering | `cellbuffer/main.go` |
| Basic counter | `simple/main.go` |

## Examples by Capability

### Installation and Progress Tracking

**Use package-manager when you need:**
- Progress tracking across multiple sequential tasks
- Combined spinner and progress bar display
- State management for multi-step processes
- Tea.Println for persistent output above TUI

**File**: `examples/package-manager/main.go`
**Key patterns**: Progress bar updates, spinner integration, sequential command batching, custom message types for completion events

**Use progress-animated when you need:**
- Smooth animated progress visualization
- Progress bar with gradient styling
- Indeterminate or determinate progress states

**File**: `examples/progress-animated/main.go`
**Key patterns**: Progress.FrameMsg handling, animated updates

**Use progress-static when you need:**
- Manual progress incrementation
- Static progress bar updates

**File**: `examples/progress-static/main.go`
**Key patterns**: SetPercent calls, manual progress control

**Use progress-download when you need:**
- File download with progress tracking
- Network operation progress visualization

**File**: `examples/progress-download/main.go`
**Key patterns**: HTTP client integration, progress callbacks

### Form Input and Validation

**Use credit-card-form when you need:**
- Multi-step form flow with navigation
- Input validation with immediate feedback
- Focus management across multiple fields
- Custom validators for text inputs

**File**: `examples/credit-card-form/main.go`
**Key patterns**: textinput.Model array, focus cycling, validator functions, conditional styling

**Use textinput when you need:**
- Single-line text input
- Basic input capture

**File**: `examples/textinput/main.go`
**Key patterns**: textinput.Model usage, input events

**Use textinputs when you need:**
- Multiple independent text fields
- Focus switching between inputs
- Cursor mode changes

**File**: `examples/textinputs/main.go`
**Key patterns**: Input array management, focus state tracking

**Use textarea when you need:**
- Multi-line text editing
- Text area with scrolling

**File**: `examples/textarea/main.go`
**Key patterns**: textarea.Model usage, line wrapping

**Use chat when you need:**
- Multi-line input with display area
- Chat-like interface patterns

**File**: `examples/chat/main.go`
**Key patterns**: Combined viewport and textarea

**Use split-editors when you need:**
- Multiple text areas in single view
- Editor pane focus management
- Side-by-side text editing

**File**: `examples/split-editors/main.go`
**Key patterns**: Multiple textarea instances, focus switching logic

### Data Display and Selection

**Use table when you need:**
- Row-based data presentation
- Keyboard navigation of rows
- Cell-based layout with headers

**File**: `examples/table/main.go`
**Key patterns**: table.Model, column definitions, row selection

**Use table-resize when you need:**
- Dynamic table sizing based on window
- Responsive table layouts

**File**: `examples/table-resize/main.go`
**Key patterns**: WindowSizeMsg handling, dynamic column width

**Use list-default when you need:**
- Standard list interface
- Filtering, pagination, help text
- Item selection with defaults

**File**: `examples/list-default/main.go`
**Key patterns**: list.Model, list.Item interface, default styling

**Use list-fancy when you need:**
- Custom list item delegates
- Advanced list customization
- Custom rendering per item
- Dynamic item insertion/removal

**File**: `examples/list-fancy/main.go`
**Key patterns**: Custom ItemDelegate, status messages, toggle features

**Use list-simple when you need:**
- Minimal list appearance
- Compact list presentation

**File**: `examples/list-simple/main.go`
**Key patterns**: Simplified list styling

**Use file-picker when you need:**
- File system navigation
- File type filtering
- File selection interface

**File**: `examples/file-picker/main.go`
**Key patterns**: filepicker.Model, allowed types, directory traversal

**Use result when you need:**
- Choice selection from options
- Simple menu interface

**File**: `examples/result/main.go`
**Key patterns**: Choice rendering, selection handling

**Use autocomplete when you need:**
- Text input with suggestions
- Completion dropdown

**File**: `examples/autocomplete/main.go`
**Key patterns**: Suggestion matching, completion logic

### Content Viewing

**Use pager when you need:**
- Scrollable text content viewing
- Less-like functionality
- Document reader interface

**File**: `examples/pager/main.go`
**Key patterns**: viewport.Model, content scrolling, keyboard navigation

**Use paginator when you need:**
- Page-based navigation
- Content chunking across pages
- Numeric page controls

**File**: `examples/paginator/main.go`
**Key patterns**: paginator.Model, page calculation, item slicing

**Use glamour when you need:**
- Markdown rendering in viewport
- Styled markdown display

**File**: `examples/glamour/main.go`
**Key patterns**: glamour.Render integration, markdown to ANSI

### View Management and Navigation

**Use views when you need:**
- Multiple distinct screens
- View state machine
- Screen transitions with custom progress

**File**: `examples/views/main.go`
**Key patterns**: State enum, view switching logic, progress bar from scratch

**Use composable-views when you need:**
- Composing multiple bubble models together
- Switching between different model types
- Model delegation pattern

**File**: `examples/composable-views/main.go`
**Key patterns**: Embedded models (timer, spinner), state-based rendering, Init command batching

**Use tabs when you need:**
- Tab-based navigation
- Multiple content sections
- Horizontal navigation pattern

**File**: `examples/tabs/main.go`
**Key patterns**: Tab state management, lipgloss tab styling

### Loading and Status Indicators

**Use spinner when you need:**
- Loading state indication
- Indefinite wait feedback

**File**: `examples/spinner/main.go`
**Key patterns**: spinner.Model, Tick command

**Use spinners when you need:**
- Demonstration of available spinner types
- Spinner selection interface

**File**: `examples/spinners/main.go`
**Key patterns**: Spinner type enumeration, style switching

### Time-Based Operations

**Use timer when you need:**
- Countdown timer
- Time-limited operations
- Timeout handling

**File**: `examples/timer/main.go`
**Key patterns**: timer.Model, timer.TickMsg, timeout events

**Use stopwatch when you need:**
- Elapsed time tracking
- Uptime display
- Time measurement

**File**: `examples/stopwatch/main.go`
**Key patterns**: Custom tick implementation, time accumulation

### Network and External Operations

**Use http when you need:**
- HTTP requests within TUI
- Network operation handling
- Async HTTP call patterns

**File**: `examples/http/main.go`
**Key patterns**: tea.Cmd for async requests, custom message types (statusMsg, errMsg), HTTP client usage

**Use exec when you need:**
- Launching external commands
- Editor integration
- Process execution during TUI runtime

**File**: `examples/exec/main.go`
**Key patterns**: tea.ExecProcess, os/exec integration, editor spawning

**Use tui-daemon-combo when you need:**
- Dual TUI/daemon mode operation
- Background process with optional UI
- Service with interactive mode

**File**: `examples/tui-daemon-combo/main.go`
**Key patterns**: Mode switching, daemon process patterns

### Real-Time and Event Handling

**Use realtime when you need:**
- Channel-based event streams
- Asynchronous activity handling
- Real-time data updates

**File**: `examples/realtime/main.go`
**Key patterns**: Channel subscriptions, waitForActivity pattern, event listener commands

**Use send-msg when you need:**
- Custom message types
- Application-specific events
- Message-driven architecture

**File**: `examples/send-msg/main.go`
**Key patterns**: Custom tea.Msg types, message dispatching

**Use sequence when you need:**
- Sequential command execution
- Ordered operation chains
- tea.Sequence usage

**File**: `examples/sequence/main.go`
**Key patterns**: tea.Sequence command, ordered execution

**Use debounce when you need:**
- Input throttling
- Preventing rapid event handling
- Delayed action execution

**File**: `examples/debounce/main.go`
**Key patterns**: Debounce timing, event coalescing

### Screen and Terminal Management

**Use altscreen-toggle when you need:**
- Switching between alt screen and normal buffer
- Preserving terminal content
- Screen buffer management

**File**: `examples/altscreen-toggle/main.go`
**Key patterns**: tea.EnterAltScreen, tea.ExitAltScreen commands

**Use fullscreen when you need:**
- Fullscreen TUI mode
- Maximum viewport usage

**File**: `examples/fullscreen/main.go`
**Key patterns**: Fullscreen initialization

**Use window-size when you need:**
- Responsive layout based on terminal size
- Window resize handling
- Dynamic content sizing

**File**: `examples/window-size/main.go`
**Key patterns**: tea.WindowSizeMsg handling, dynamic layout

**Use set-window-title when you need:**
- Terminal window title updates
- Title bar customization

**File**: `examples/set-window-title/main.go`
**Key patterns**: Window title escape sequences

**Use cellbuffer when you need:**
- Direct cell buffer manipulation
- Low-level rendering control

**File**: `examples/cellbuffer/main.go`
**Key patterns**: Cell buffer operations

### Input and Interaction

**Use mouse when you need:**
- Mouse click handling
- Mouse motion tracking
- Cursor position detection

**File**: `examples/mouse/main.go`
**Key patterns**: tea.MouseMsg handling, mouse event types

**Use focus-blur when you need:**
- Window focus state detection
- Focus event handling

**File**: `examples/focus-blur/main.go`
**Key patterns**: Focus/blur message handling

**Use prevent-quit when you need:**
- Confirmation before exit
- Prevent accidental termination
- Quit protection logic

**File**: `examples/prevent-quit/main.go`
**Key patterns**: Quit confirmation state, conditional exit

**Use suspend when you need:**
- Process suspension support
- Ctrl+Z handling
- Background/foreground transitions

**File**: `examples/suspend/main.go`
**Key patterns**: Suspend signal handling

### Help and Documentation

**Use help when you need:**
- Contextual help display
- Keyboard shortcut documentation
- Help menu toggle

**File**: `examples/help/main.go`
**Key patterns**: help.Model, key.Binding definitions, help text rendering

### I/O Integration

**Use pipe when you need:**
- Reading from stdin
- Writing to stdout
- Shell pipeline integration
- TUI with pipe mode support

**File**: `examples/pipe/main.go`
**Key patterns**: Stdin detection, pipe vs interactive mode

### Basic Patterns

**Use simple when you need:**
- Minimal TUI structure
- Counter or basic state example
- Learning basic Update/View pattern

**File**: `examples/simple/main.go`
**Key patterns**: Minimal Model implementation, basic key handling, simple state management

## Implementation Patterns

### Progress Tracking Pattern
Combine spinner and progress bar models in a single view with tea.Batch for concurrent updates. Use custom message types to signal completion of individual steps. Track index and total items for progress calculation.

**Example**: `package-manager/main.go`

### Multi-Step Form Pattern
Use array of input models with focus index. Implement next/prev navigation functions. Apply validators to each input. Manage focus state with Focus()/Blur() calls on active input.

**Example**: `credit-card-form/main.go`

### View Composition Pattern
Embed multiple bubble models in main model. Use state enum to track active view. Delegate Update calls to active model based on state. Return composed views using lipgloss layout functions.

**Example**: `composable-views/main.go`

### List Customization Pattern
Implement custom ItemDelegate for full control over item rendering. Use status messages for feedback. Provide custom key bindings beyond defaults. Toggle list features dynamically.

**Example**: `list-fancy/main.go`

### Async Operation Pattern
Define custom message type for operation results. Return tea.Cmd that performs async work and returns custom message. Handle custom message in Update to process results. Use for HTTP, file I/O, or any blocking operation.

**Example**: `http/main.go`

### Real-Time Event Pattern
Create channel for event stream. Launch goroutine command to generate events. Implement wait command that listens on channel. Return events as messages to Update function. Use for live data feeds, websockets, or monitoring.

**Example**: `realtime/main.go`

### External Process Pattern
Use tea.ExecProcess for commands that need full terminal control. Handle process completion with custom message. Manage screen state around process execution. Works for editors, shells, or interactive tools.

**Example**: `exec/main.go`

### Responsive Layout Pattern
Handle tea.WindowSizeMsg to get terminal dimensions. Calculate component sizes based on available space. Update model dimensions in response to resize. Re-render with adjusted constraints.

**Example**: `table-resize/main.go`, `window-size/main.go`

### Input Validation Pattern
Define validator function with error return. Assign to textinput.Validate field. Validation runs on each keystroke. Display error state visually in View.

**Example**: `credit-card-form/main.go`

### Sequential Task Pattern
Use tea.Sequence to chain commands in order. Each command executes after previous completes. Final command can trigger quit or next phase. Useful for setup sequences or multi-stage operations.

**Example**: `sequence/main.go`, `package-manager/main.go` (in completion handler)

### Help System Pattern
Define key.Binding for each command. Implement key.KeyMap interface. Pass to help.Model. Toggle help visibility with dedicated keybinding. Supports short and full help modes.

**Example**: `help/main.go`

### Pagination Pattern
Use paginator.Model to manage page state. Calculate visible items based on page and per-page count. Slice data array for current page. Render paginator controls separately.

**Example**: `paginator/main.go`

### Tab Navigation Pattern
Track active tab index. Render tab headers with active styling. Switch content based on tab state. Use lipgloss for visual tab appearance.

**Example**: `tabs/main.go`

### Dual Mode Pattern
Implement flag-based mode selection. Separate TUI and headless code paths. Allow same binary to run as daemon or interactive. Useful for tools that need both modes.

**Example**: `tui-daemon-combo/main.go`

### Debounce Pattern
Track last input time. Delay action until quiet period. Cancel pending actions on new input. Prevents excessive processing of rapid events.

**Example**: `debounce/main.go`

### Focus Management Pattern
Track focus index across multiple inputs or components. Implement focus ring navigation. Call Focus()/Blur() methods when switching. Visual indication of focused component.

**Example**: `textinputs/main.go`, `split-editors/main.go`

### File Selection Pattern
Use filepicker.Model for navigation. Set AllowedTypes for filtering. Handle DidSelectFile message. Display error for invalid selections.

**Example**: `file-picker/main.go`
