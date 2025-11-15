# Huh Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Multi-field form with validation | `examples/burger/main.go` |
| Dynamic field generation based on input | `examples/conditional/main.go` |
| Autocomplete with context-dependent suggestions | `examples/dynamic/dynamic-suggestions/main.go` |
| File selection interface | `examples/filepicker/main.go` |
| Network-accessible forms | `examples/ssh-form/main.go` |
| Full bubbletea integration | `examples/bubbletea/main.go` |
| Multiple form pages/navigation | `examples/multiple-groups/main.go` |
| Column-based layout | `examples/layout/columns/main.go` |
| Custom theming | `examples/theme/main.go` |
| Accessibility support | `examples/accessibility/main.go` |
| Secure password input | `examples/accessibility-secure-input/main.go` |
| Field visibility controls | `examples/hide/main.go` |
| Conditional field skipping | `examples/skip/main.go` |
| Scrollable content | `examples/scroll/main.go` |
| Runtime-generated forms | `examples/dynamic/dynamic-all/main.go` |
| Multi-selection interface | `examples/stickers/main.go` |

## Examples by Capability

### Complex Multi-Step Forms
**Use burger when you need:**
- Multi-step workflow with validation at each stage
- Struct-based state management across form pages
- Custom enums with string conversion (Spice enum pattern)
- Spinner integration for processing steps
- Receipt/summary generation after completion
- Cross-field validation and constraints

**File**: `examples/burger/main.go`
**Key patterns**: Group-based pagination, method chaining for field configuration, validation functions returning errors

### Conditional Logic and Dynamic Fields
**Use conditional when you need:**
- Fields that appear based on previous selections
- Dynamic titles that update based on other field values
- OptionsFunc pattern for runtime option generation
- TitleFunc for context-dependent labels
- State-dependent form structure

**File**: `examples/conditional/main.go`
**Key patterns**: TitleFunc/OptionsFunc callbacks with field dependencies

### Context-Aware Autocomplete
**Use dynamic-suggestions when you need:**
- Suggestions that change based on another field
- PlaceholderFunc for dynamic hints
- SuggestionsFunc with field dependency tracking
- Organization/repository pattern (works for any hierarchical data)

**File**: `examples/dynamic/dynamic-suggestions/main.go`
**Key patterns**: Function-based placeholders and suggestions with watched dependencies

### File System Integration
**Use filepicker when you need:**
- File browser interface in terminal
- File type filtering (by extension)
- Path selection and validation
- Directory navigation

**File**: `examples/filepicker/main.go`
**Key patterns**: NewFilePicker with extension filters, path value binding

**Use filepicker-picking when you need:**
- Enhanced file selection with additional controls

**File**: `examples/filepicker-picking/main.go`

### Network and SSH Integration
**Use ssh-form when you need:**
- Forms accessible over SSH connections
- Remote form hosting on specific ports
- Wish middleware integration with bubbletea
- Session management for network clients
- Server-side form rendering

**File**: `examples/ssh-form/main.go`
**Key patterns**: Wish server setup, bubbletea.Middleware, signal handling for graceful shutdown

### Full Bubbletea Framework
**Use bubbletea when you need:**
- Full control over application state and rendering
- Custom styling with lipgloss
- Message-based state updates
- View composition and layout control
- Integration with other bubbletea components

**File**: `examples/bubbletea/main.go`
**Key patterns**: Model with Init/Update/View methods, message passing, theme integration

**Use bubbletea-options when you need:**
- Extended bubbletea options and configuration patterns

**File**: `examples/bubbletea-options/main.go`

### Multi-Page Navigation
**Use multiple-groups when you need:**
- Forms split across multiple pages/screens
- Navigation between form sections
- WithHeight for scrollable sections
- Group-based organization with independent submission

**File**: `examples/multiple-groups/main.go`
**Key patterns**: Multiple NewGroup calls, WithHeight for viewport control

### Layout Systems
**Use layout/columns when you need:**
- Side-by-side field arrangement
- Multi-column form layouts
- LayoutColumns configuration

**File**: `examples/layout/columns/main.go`
**Key patterns**: WithLayout(huh.LayoutColumns(n))

**Use layout/stack when you need:**
- Vertical stacking of form sections

**File**: `examples/layout/stack/main.go`

**Use layout/grid when you need:**
- Grid-based component positioning

**File**: `examples/layout/grid/main.go`

**Use layout/default when you need:**
- Standard single-column layout patterns

**File**: `examples/layout/default/main.go`

### Theming and Styling
**Use theme when you need:**
- Custom color schemes and styling
- Theme switching at runtime
- Pre-built themes (Base, Dracula, Base16, Charm, Catppuccin)
- Theme preview and selection interface

**File**: `examples/theme/main.go`
**Key patterns**: huh.Theme* constructors, theme application to forms

### Accessibility Features
**Use accessibility when you need:**
- ACCESSIBLE environment variable support
- Enhanced keyboard navigation
- Screen reader compatibility patterns

**File**: `examples/accessibility/main.go`

**Use accessibility-secure-input when you need:**
- Accessible secure password/sensitive input fields
- EchoMode configuration with accessibility support

**File**: `examples/accessibility-secure-input/main.go`

### Field Visibility and Control
**Use hide when you need:**
- Conditional field visibility based on logic
- Runtime show/hide of form elements

**File**: `examples/hide/main.go`
**Key patterns**: Field visibility predicates

**Use skip when you need:**
- Programmatic field skipping
- Conditional field inclusion in form flow

**File**: `examples/skip/main.go`

### Scrolling and Large Content
**Use scroll when you need:**
- Scrollable option lists
- Large form sections with viewport management
- Height-constrained selection menus

**File**: `examples/scroll/main.go`
**Key patterns**: Height configuration for scrollable regions

### Dynamic Form Generation
**Use dynamic-all when you need:**
- Complete runtime form generation
- Forms built from data structures
- Programmatic field creation

**File**: `examples/dynamic/dynamic-all/main.go`

**Use dynamic-bubbletea when you need:**
- Dynamic forms within bubbletea framework

**File**: `examples/dynamic/dynamic-bubbletea/main.go`

**Use dynamic-count when you need:**
- Counter-based field generation
- Variable-length field arrays

**File**: `examples/dynamic/dynamic-count/main.go`

**Use dynamic-country when you need:**
- Geographic/hierarchical data selection
- Country/region selection patterns

**File**: `examples/dynamic/dynamic-country/main.go`

**Use dynamic-increment when you need:**
- Incremental field addition
- User-controlled field count

**File**: `examples/dynamic/dynamic-increment/main.go`

**Use dynamic-markdown when you need:**
- Markdown content rendering in forms
- Rich text display within form context

**File**: `examples/dynamic/dynamic-markdown/main.go`

**Use dynamic-name when you need:**
- Name-based dynamic field generation
- Field creation based on string inputs

**File**: `examples/dynamic/dynamic-name/main.go`

### Multi-Selection Interfaces
**Use stickers when you need:**
- Multiple item selection with visual feedback
- Checkbox-style multi-select
- Selected state tracking

**File**: `examples/stickers/main.go`
**Key patterns**: NewMultiSelect with Option configurations

### Time-Based Operations
**Use timer when you need:**
- Time limits on form completion
- Timeout handling in forms
- Time-based field behavior

**File**: `examples/timer/main.go`

### Help Systems
**Use help when you need:**
- Integrated help text display
- Keyboard shortcut documentation
- Context-sensitive help

**File**: `examples/help/main.go`

### Git and GitHub Integration
**Use gh/create when you need:**
- GitHub workflow interfaces
- Repository creation/PR flows
- Push/fork/PR decision logic

**File**: `examples/gh/create.go`

**Use git when you need:**
- Git operation command builders
- Repository interaction forms

**File**: `examples/git/main.go`

### CLI Tool Integration
**Use gum when you need:**
- Gum CLI tool styling patterns
- Integration with gum commands

**File**: `examples/gum/main.go`

### Documentation Examples
The `examples/readme/` directory contains minimal examples for documentation:

- `examples/readme/input/main.go` - Basic input field
- `examples/readme/main/main.go` - Primary example
- `examples/readme/confirm/main.go` - Confirmation dialogs
- `examples/readme/multiselect/main.go` - Multi-selection
- `examples/readme/note/main.go` - Note/info display
- `examples/readme/select/main.go` - Selection menus
- `examples/readme/select/scroll/main.go` - Scrollable selections
- `examples/readme/text/main.go` - Text input and display

## Implementation Patterns

### Form Construction
```go
// Group-based organization
huh.NewForm(
    huh.NewGroup(field1, field2, field3),
    huh.NewGroup(field4, field5),
)

// Method chaining for field configuration
huh.NewInput().
    Title("Label").
    Placeholder("hint").
    Validate(validatorFunc).
    Value(&variable)
```

### Validation
```go
// Validation function signature
func(s string) error {
    if !valid(s) {
        return errors.New("validation message")
    }
    return nil
}

// Attach to field
.Validate(validationFunc)
```

### Dynamic Content Generation
```go
// Function-based placeholders
.PlaceholderFunc(func() string {
    // Return placeholder based on state
}, &dependencyVar)

// Dynamic suggestions
.SuggestionsFunc(func() []string {
    // Return suggestions based on state
}, &dependencyVar)

// Dynamic options
.OptionsFunc(func() []huh.Option[T] {
    // Return options based on state
}, &dependencyVar)

// Dynamic titles
.TitleFunc(func() string {
    // Return title based on state
}, &dependencyVar)
```

### State Management
```go
// Struct-based form state
type FormState struct {
    Field1 string
    Field2 int
    Field3 []string
}

var state FormState

// Bind fields to struct members
huh.NewInput().Value(&state.Field1)
```

### Styling Integration
```go
// Lipgloss styling
renderer := lipgloss.NewRenderer(os.Stdout)
styles := styleDefinition(renderer)

// Apply theme
form.WithTheme(huh.ThemeDracula())
```

### Network Integration
```go
// SSH server middleware
wish.NewServer(
    wish.WithAddress(host + ":" + port),
    wish.WithMiddleware(
        bubbletea.Middleware(handlerFunc),
    ),
)
```

### Layout Configuration
```go
// Column layout
.WithLayout(huh.LayoutColumns(n))

// Height constraints
.WithHeight(lines)
```

## File Structure

Total examples: 40 Go source files across 43 subdirectories

### Core Examples (root level)
- accessibility, accessibility-secure-input
- bubbletea, bubbletea-options
- burger (complex multi-step)
- conditional
- filepicker, filepicker-picking
- gh, git, gum
- help, hide
- multiple-groups
- scroll, skip
- ssh-form
- stickers
- theme
- timer

### Dynamic Examples (dynamic/ subdirectory)
- dynamic-all, dynamic-bubbletea
- dynamic-count, dynamic-country
- dynamic-increment, dynamic-markdown
- dynamic-name, dynamic-suggestions

### Layout Examples (layout/ subdirectory)
- columns, default, grid, stack

### Documentation Examples (readme/ subdirectory)
- confirm, input, main, multiselect
- note, select, text
- select/scroll

## Dependencies

External packages from charmbracelet ecosystem:
- github.com/charmbracelet/huh (core form library)
- github.com/charmbracelet/bubbletea (TUI framework)
- github.com/charmbracelet/lipgloss (styling)
- github.com/charmbracelet/ssh (SSH server)
- github.com/charmbracelet/wish (SSH middleware)
- github.com/charmbracelet/log (logging)

Build system: Go modules with local replace directives
