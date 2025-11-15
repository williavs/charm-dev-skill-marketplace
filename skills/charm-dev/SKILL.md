---
name: charm-dev
description: Build terminal UIs using Bubbletea and the Charm ecosystem (Lipgloss, Huh, Gum). Use when creating TUIs, interactive CLIs, terminal applications, or when the user mentions Bubbletea, Charm libraries, terminal styling, or TUI development.
---

# Charm Development

Build production-quality terminal user interfaces using Bubbletea and the Charm ecosystem.

## Primary Workflow

**1. Understand the architecture** → `references/architecture/ELM-ARCHITECTURE.md`

Read this first to understand The Elm Architecture (TEA) pattern that Bubbletea implements.

**2. Find your pattern** → `references/bubbletea/examples/CONTEXTUAL-INVENTORY.md`

Quick reference table (lines 5-45) maps use cases to example files. Find what you need, go to the example, copy the pattern.

**3. Check API details** → `references/api/BUBBLETEA_API_FEATURES.md`

Comprehensive API reference for advanced features (message filtering, focus reporting, bracketed paste, testing, etc.).

**4. Style your UI** → `references/lipgloss/examples/`

Use Lipgloss for all styling. Examples show layouts, colors, borders, alignment, and responsive design.

## Architecture Quick Reference

Read these to understand Bubbletea's design:

- **`references/architecture/ELM-ARCHITECTURE.md`** - The Elm Architecture (Model-View-Update pattern)
- **`references/architecture/MVC-IN-BUBBLETEA.md`** - How TEA maps to MVC if you're familiar with that pattern
- **`references/architecture/COMPONENTS.md`** - Component composition and focus management
- **`references/architecture/BEST-PRACTICES.md`** - Do's and don'ts, common anti-patterns

## Component Library Quick Reference

**Find examples in:** `references/bubbletea/examples/CONTEXTUAL-INVENTORY.md`

**Common patterns:**

| Need | Example File |
|------|--------------|
| Multi-step form with validation | `credit-card-form/main.go:83` |
| Window resize handling | `window-size/main.go:360` |
| External command execution | `exec/main.go:288` |
| Real-time event handling | `realtime/main.go:305` |
| HTTP requests | `http/main.go:279` |
| Multiple view switching | `composable-views/main.go:228` |
| File selection | `file-picker/main.go:172` |
| Selectable lists | `list-default/main.go:148` |
| Progress tracking | `package-manager/main.go:47` |

Full table with 45+ patterns in CONTEXTUAL-INVENTORY.md lines 5-45.

## Available Libraries

**Bubbletea** (`references/bubbletea/`) - Core TUI framework with 47 examples
- Forms, lists, tables, progress, spinners, timers
- Real-time updates, HTTP, external commands
- Mouse, keyboard, window events

**Lipgloss** (`references/lipgloss/`) - Styling and layout with 20 examples
- Colors, borders, padding, margins
- Layouts (horizontal, vertical, tables)
- Responsive design

**Huh** (`references/huh/`) - Interactive forms with 40 examples
- Input fields, selects, confirms
- Multi-step forms, validation
- Dynamic forms, SSH integration

**Gum** (`references/gum/`) - CLI components with 31 examples
- One-line CLI tools
- Filters, pickers, confirmations
- Quick script integration

## Quick Start Pattern

```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

// Model - All application state
type model struct {
    cursor int
    choices []string
}

// Init - Initial command
func (m model) Init() tea.Cmd {
    return nil
}

// Update - Handle messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "up":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

// View - Render UI
func (m model) View() string {
    s := "Choose:\n\n"
    for i, choice := range m.choices {
        cursor := " "
        if i == m.cursor {
            cursor = ">"
        }
        s += fmt.Sprintf("%s %s\n", cursor, choice)
    }
    return s
}

func main() {
    p := tea.NewProgram(model{
        choices: []string{"Option 1", "Option 2", "Option 3"},
    })
    p.Run()
}
```

## Development Approach

1. **Find the closest example** in CONTEXTUAL-INVENTORY.md
2. **Read the complete example code** - Don't skip this
3. **Copy the pattern** - Don't reinvent, adapt
4. **Compose components** - Use `composable-views` pattern for complex UIs
5. **Style with Lipgloss** - Keep styling separate from logic

## Common Use Cases

**Multi-step wizards** → `references/bubbletea/examples/credit-card-form/`

**Data tables** → `references/bubbletea/examples/table/`

**File browsers** → `references/bubbletea/examples/file-picker/`

**Progress indicators** → `references/bubbletea/examples/package-manager/`

**Chat interfaces** → `references/bubbletea/examples/chat/`

**External processes** → `references/bubbletea/examples/exec/`

**Real-time updates** → `references/bubbletea/examples/realtime/`

## Key Principles

**Model = State** - All state in one struct

**View = Pure** - (Model → String), no side effects

**Update = Pure** - ((Model, Msg) → (Model, Cmd)), no I/O

**Commands = Side Effects** - Async operations return messages

**Always handle WindowSizeMsg** - Make UIs responsive

**Use Lipgloss for styling** - Don't hand-craft ANSI codes

**Batch commands** - Use `tea.Batch()` for concurrent operations

## Resources Structure

- **`references/architecture/`** - TEA, MVC, components, best practices
- **`references/bubbletea/`** - 47 Bubbletea examples + CONTEXTUAL-INVENTORY.md
- **`references/lipgloss/`** - 20 styling examples
- **`references/huh/`** - 40 form examples
- **`references/gum/`** - 31 CLI component examples
- **`references/api/`** - BUBBLETEA_API_FEATURES.md (comprehensive API reference)

## Getting Help

**For specific use case:** Check CONTEXTUAL-INVENTORY.md lines 5-45 first

**For architecture questions:** Read ELM-ARCHITECTURE.md

**For component composition:** Read COMPONENTS.md

**For common mistakes:** Read BEST-PRACTICES.md

**For API details:** Check BUBBLETEA_API_FEATURES.md
