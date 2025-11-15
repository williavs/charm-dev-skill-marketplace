# Lipgloss Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|--------------|
| Tabular data with borders and styling | `examples/table/pokemon/main.go` |
| Filesystem tree rendering | `examples/tree/files/main.go` |
| Hierarchical data structures | `examples/tree/simple/main.go` |
| Multi-pane layout composition | `examples/layout/main.go` |
| Checklist with conditional rendering | `examples/list/grocery/main.go` |
| Nested list structures | `examples/list/simple/main.go` |
| SSH terminal integration | `examples/ssh/main.go` |
| Custom enumerators | `examples/list/duckduckgoose/main.go` |
| Conditional row styling | `examples/table/pokemon/main.go` |
| Color grids and gradients | `examples/layout/main.go` |
| Tab navigation interfaces | `examples/layout/main.go` |
| Status bars | `examples/layout/main.go` |
| Modal dialogs | `examples/layout/main.go` |

## Examples by Capability

### Table Rendering

**Use table/pokemon when you need:**
- Data display with structured columns and rows
- Conditional styling based on cell content
- Alternating row colors for readability
- Type-based color coding
- Row highlighting based on data values
- Border customization and header styling

**File**: `examples/table/pokemon/main.go`
**Key patterns**: StyleFunc for dynamic per-cell styling, color maps for content-based coloring, header transformation functions

**Use table/ansi when you need:**
- ANSI color and formatting testing
- Terminal capability verification
- Color palette display

**File**: `examples/table/ansi/main.go`
**Key patterns**: ANSI escape sequence rendering

**Use table/chess when you need:**
- Grid layouts with alternating background colors
- Board-like representations
- Position-based styling

**File**: `examples/table/chess/main.go`
**Key patterns**: Position-based color alternation

**Use table/languages when you need:**
- Multi-language text display
- Right-to-left text support
- International character rendering

**File**: `examples/table/languages/main.go`
**Key patterns**: Unicode handling, bidirectional text

**Use table/mindy when you need:**
- Character styling demonstrations
- Font attribute testing

**File**: `examples/table/mindy/main.go`
**Key patterns**: Multiple style attribute combinations

### Tree Rendering

**Use tree/files when you need:**
- Recursive directory traversal
- Filesystem visualization
- Hierarchical file structure display
- Dot-file filtering

**File**: `examples/tree/files/main.go`
**Key patterns**: Recursive tree building, os.ReadDir integration, child node attachment

**Use tree/simple when you need:**
- Static hierarchical data visualization
- Operating system or category trees
- Nested relationship display

**File**: `examples/tree/simple/main.go`
**Key patterns**: Declarative tree construction, nested tree.New() calls

**Use tree/background when you need:**
- Tree styling with background colors
- Visual separation through backgrounds

**File**: `examples/tree/background/main.go`
**Key patterns**: Background color application to tree nodes

**Use tree/makeup when you need:**
- Cosmetic tree variations
- Custom styling patterns

**File**: `examples/tree/makeup/main.go`
**Key patterns**: Style composition for tree elements

**Use tree/rounded when you need:**
- Rounded border tree styling
- Softer visual appearance

**File**: `examples/tree/rounded/main.go`
**Key patterns**: Rounded border application

**Use tree/styles when you need:**
- Multiple tree styling approaches
- Style comparison
- Custom tree aesthetics

**File**: `examples/tree/styles/main.go`
**Key patterns**: Various style configuration methods

**Use tree/toggle when you need:**
- Interactive-style expandable trees
- Collapsible state representation
- Node expansion indicators

**File**: `examples/tree/toggle/main.go`
**Key patterns**: Toggle state visualization

### List Rendering

**Use list/grocery when you need:**
- Item completion tracking
- Checkmark indicators
- Strikethrough for completed items
- Custom enumerators based on state
- Style functions that vary by item status

**File**: `examples/list/grocery/main.go`
**Key patterns**: Custom enumerator functions, EnumeratorStyleFunc, ItemStyleFunc for state-dependent rendering

**Use list/simple when you need:**
- Basic nested lists
- Roman numeral enumeration
- Hierarchical list structures

**File**: `examples/list/simple/main.go`
**Key patterns**: list.New() with nested list.New(), list.Roman enumerator

**Use list/duckduckgoose when you need:**
- Custom enumeration patterns
- Non-standard list markers

**File**: `examples/list/duckduckgoose/main.go`
**Key patterns**: Custom enumerator implementation

**Use list/glow when you need:**
- Markdown-styled list formatting
- Document-style rendering

**File**: `examples/list/glow/main.go`
**Key patterns**: Markdown-compatible styling

**Use list/roman when you need:**
- Roman numeral enumeration
- Classical list styling

**File**: `examples/list/roman/main.go`
**Key patterns**: list.Roman enumerator

**Use list/sublist when you need:**
- Multiple nesting levels
- Complex hierarchical structures

**File**: `examples/list/sublist/main.go`
**Key patterns**: Deep nesting patterns

### Layout Composition

**Use layout when you need:**
- Multi-section terminal interfaces
- Tab navigation rendering
- Color grids and gradients
- Dialog box placement
- Status bar construction
- Horizontal and vertical content joining
- Centered content placement
- Complex multi-column layouts

**File**: `examples/layout/main.go`
**Key patterns**: lipgloss.JoinHorizontal, lipgloss.JoinVertical, lipgloss.Place for positioning, adaptive color schemes, border composition, terminal width detection

### SSH Integration

**Use ssh when you need:**
- Remote terminal session handling
- Per-client renderer customization
- Terminal capability detection per session
- Adaptive color profiles for remote clients
- PTY integration
- Custom SSH server middleware

**File**: `examples/ssh/main.go`
**Key patterns**: Custom renderer creation with lipgloss.NewRenderer(session), termenv.Output integration, SSH session to terminal output bridging, client-specific style initialization

## Implementation Patterns

### Dynamic Styling

**StyleFunc Pattern**: Apply styles dynamically based on row, column, or content
```go
StyleFunc(func(row, col int) lipgloss.Style {
    if row == table.HeaderRow {
        return headerStyle
    }
    if data[row][col] == "specific_value" {
        return specialStyle
    }
    return baseStyle
})
```

**State-Based Rendering**: Conditional styling based on item state
```go
func itemStyleFunc(items list.Items, i int) lipgloss.Style {
    if isCompleted(items.At(i)) {
        return completedStyle
    }
    return normalStyle
}
```

### Layout Techniques

**Horizontal Joining**: Combine elements side-by-side
```go
row := lipgloss.JoinHorizontal(lipgloss.Top, elem1, elem2, elem3)
```

**Vertical Joining**: Stack elements vertically
```go
column := lipgloss.JoinVertical(lipgloss.Left, elem1, elem2, elem3)
```

**Centered Placement**: Position content within defined bounds
```go
centered := lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
```

### Color Management

**Adaptive Colors**: Respond to terminal light/dark mode
```go
adaptive := lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}
```

**Color Maps**: Associate colors with data values
```go
colorMap := map[string]lipgloss.Color{
    "type1": lipgloss.Color("#FF0000"),
    "type2": lipgloss.Color("#00FF00"),
}
```

**Color Gradients**: Generate color ranges
```go
blends := gamut.Blends(startColor, endColor, steps)
```

### Border Styling

**Standard Borders**: Apply built-in border styles
```go
style.Border(lipgloss.NormalBorder())
style.Border(lipgloss.RoundedBorder())
style.Border(lipgloss.ThickBorder())
```

**Custom Borders**: Define specific border characters
```go
customBorder := lipgloss.Border{
    Top: "─", Bottom: "─",
    Left: "│", Right: "│",
    TopLeft: "╭", TopRight: "╮",
    BottomLeft: "┘", BottomRight: "└",
}
```

### Renderer Customization

**Custom Output Targets**: Direct rendering to specific outputs
```go
renderer := lipgloss.NewRenderer(outputWriter)
renderer.SetOutput(termenvOutput)
```

**Client-Specific Rendering**: Adapt to terminal capabilities
```go
renderer.HasDarkBackground()
renderer.Output().BackgroundColor()
```

### Tree Construction

**Declarative Trees**: Build static hierarchies
```go
t := tree.Root("root").
    Child("child1").
    Child(tree.New().Root("child2").Child("grandchild"))
```

**Dynamic Trees**: Populate from data sources
```go
func addBranches(root *tree.Tree, path string) error {
    items, _ := os.ReadDir(path)
    for _, item := range items {
        if item.IsDir() {
            branch := tree.Root(item.Name())
            root.Child(branch)
            addBranches(branch, filepath.Join(path, item.Name()))
        } else {
            root.Child(item.Name())
        }
    }
}
```

### List Construction

**Simple Lists**: Flat item lists
```go
l := list.New("item1", "item2", "item3")
```

**Nested Lists**: Hierarchical structures
```go
l := list.New(
    "item1",
    list.New("subitem1", "subitem2"),
    "item2",
)
```

**Custom Enumerators**: Define enumeration logic
```go
func customEnum(items list.Items, i int) string {
    if condition(items.At(i)) {
        return "✓"
    }
    return "•"
}
l := list.New(items...).Enumerator(customEnum)
```

## Technical Specifications

### Common Imports
- `github.com/charmbracelet/lipgloss` - Core styling
- `github.com/charmbracelet/lipgloss/table` - Table rendering
- `github.com/charmbracelet/lipgloss/tree` - Tree rendering
- `github.com/charmbracelet/lipgloss/list` - List rendering
- `github.com/muesli/termenv` - Terminal environment detection
- `github.com/charmbracelet/ssh` - SSH server integration
- `golang.org/x/term` - Terminal size detection

### Style Chain Methods
- `.Foreground()`, `.Background()` - Color application
- `.Bold()`, `.Italic()`, `.Underline()`, `.Strikethrough()` - Text attributes
- `.Padding()`, `.Margin()` - Spacing control
- `.Border()`, `.BorderForeground()` - Border configuration
- `.Width()`, `.Height()`, `.MaxWidth()`, `.MaxHeight()` - Dimension constraints
- `.Align()` - Content alignment

### Terminal Integration
- Terminal width detection via `term.GetSize()`
- Adaptive color schemes for light/dark backgrounds
- Color profile detection per client session
- PTY management for SSH sessions

### Data Flow Patterns
1. Define base styles with `lipgloss.NewStyle()`
2. Create style variations through method chaining
3. Apply styles via render functions, StyleFunc, or style methods
4. Join or place rendered elements for layout composition
5. Output final string with `fmt.Println()` or session writer
