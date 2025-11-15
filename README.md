# Charm Development Skills Marketplace

Skills for building terminal user interfaces using Bubbletea and the Charm ecosystem.

## Available Skills

### charm-dev

Comprehensive skill for building TUIs with Bubbletea, Lipgloss, Huh, and Gum.

**Includes:**
- 47 Bubbletea examples with contextual inventory
- Architecture guides: ELM/TEA, MVC patterns, component composition
- 20 Lipgloss styling examples
- 40 Huh interactive form examples
- 31 Gum CLI component examples
- API reference documentation

## Installation

Add this marketplace:

```bash
/plugin marketplace add yourusername/charm-dev-skill-marketplace
```

Install the charm-dev skill:

```bash
/plugin install charm-dev@charm-dev-skills
```

Browse and install interactively:

```bash
/plugin
```

## Contents

### Architecture Documentation
- ELM-ARCHITECTURE.md - The Elm Architecture (TEA) pattern
- MVC-IN-BUBBLETEA.md - TEA to Model-View-Controller mapping
- COMPONENTS.md - Component composition and focus management
- BEST-PRACTICES.md - Patterns and anti-patterns

### Example Libraries
- Bubbletea - 47 examples (forms, lists, tables, progress, timers, HTTP, external processes)
- Lipgloss - 20 examples (colors, borders, layouts, responsive design)
- Huh - 40 examples (multi-step forms, validation, dynamic forms)
- Gum - 31 examples (filters, pickers, confirmations)

### Reference Materials
- CONTEXTUAL-INVENTORY.md - Use case to example mapping
- BUBBLETEA_API_FEATURES.md - Complete API reference

## Structure

```
charm-dev/
├── SKILL.md
└── references/
    ├── architecture/
    ├── api/
    ├── bubbletea/
    ├── lipgloss/
    ├── huh/
    └── gum/
```

## License

MIT
