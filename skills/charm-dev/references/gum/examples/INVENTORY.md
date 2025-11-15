# Gum Examples Inventory

## Overview
This directory contains 31 files demonstrating Gum TUI component usage across shell scripts, automation demos, and multi-language integrations. The examples span 783 total lines of code implementing interactive terminal interfaces for common development workflows.

## Examples Catalog

### Core Workflow Scripts
- **demo.sh** (47 lines): Interactive user onboarding flow with input collection, choice selection, and multi-step confirmation processes
- **commit.sh** (28 lines): Conventional commit message composer with type selection, scope input, and confirmation workflow for git operations
- **git-branch-manager.sh** (64 lines): Git branch operations manager supporting rebase, delete, and update commands with multi-branch selection
- **git-stage.sh** (12 lines): Interactive git staging workflow with file selection and commit confirmation

### System Information Display
- **diyfetch** (99 lines): Customizable system information display tool with adaptive terminal sizing and multiple layout configurations
- **test.sh** (52 lines): Comprehensive component testing suite demonstrating all Gum features with various styling options

### Interactive Applications
- **magic.sh** (45 lines): Card selection game with dynamic card rendering using colored borders and suit symbols
- **kaomoji.sh** (47 lines): AI-powered emoji generator with clipboard integration using external mods command
- **filter-key-value.sh** (13 lines): Key-value pair filtering example for data selection workflows

### Multi-Language Integration
- **gum.js** (13 lines): Node.js spawn integration for choice selection with stderr/stdout handling
- **gum.py** (7 lines): Python subprocess wrapper for command execution and output processing
- **gum.rb** (29 lines): Ruby implementation with hash-based color mapping and styled output

### Automation and Scripting
- **skate.sh** (7 lines): Database entry creation using skate key-value store
- **posix.sh** (5 lines): POSIX shell compatibility demonstration
- **convert-to-gif.sh** (15 lines): VHS tape file to GIF conversion utility

### Demo and Documentation Files
- **10 .tape files**: VHS animation scripts for generating demo GIFs (choose, commit, confirm, customize, demo, file, input, pager, spin, write)
- **README.md** (36 lines): Markdown content for glamour formatting demonstration
- **format.ansi** (12 lines): ANSI escape sequence examples for terminal formatting
- **Data files**: fav.txt, flavors.txt, story.txt - sample data for filtering and display operations

## Implementation Patterns

### Input Collection Patterns
- Single value input with placeholder text and validation
- Multi-choice selection with limit controls (single, multiple, unlimited)
- Password input with masked characters
- Multi-line text composition with editor-like interface

### Display and Styling Patterns
- Bordered containers with custom colors and padding
- Horizontal and vertical layout composition using join operations
- Conditional styling based on data content (card colors, git status)
- Terminal size-aware adaptive layouts

### Workflow Control Patterns
- Confirmation gates before destructive operations
- Sequential step processing with user feedback
- Error handling with fallback options
- Integration with external tools (git, mods, skate)

### Cross-Language Integration
- Process spawning and output capture in Node.js, Python, Ruby
- Error handling through stderr redirection
- Text processing and formatting across language boundaries

## Practical Applications

### NextJS Scaffolder with Auth/ShadCN
- **commit.sh** provides template for project setup confirmation workflows
- **choose.sh** patterns applicable for framework selection (NextJS versions, auth providers)
- **input.sh** patterns for project configuration (name, description, API keys)
- **confirm.sh** patterns for installation confirmations and destructive operations
- Multi-step wizard patterns from **demo.sh** for guided setup processes

### Inter-Agent Communication Systems
- **filter-key-value.sh** demonstrates agent discovery and selection mechanisms
- **git-branch-manager.sh** provides templates for distributed workflow coordination
- **magic.sh** shows dynamic interface generation based on system state
- Cross-language examples show integration patterns for polyglot agent systems

### GitHub Search Tools
- **filter-key-value.sh** provides exact pattern for repository/issue filtering
- **git-branch-manager.sh** demonstrates branch and remote management interfaces
- **test.sh** filter examples show search result presentation patterns
- Input validation patterns applicable to search query building

### TUI for App Distribution/Package Management
- **diyfetch** provides system information gathering patterns for compatibility checking
- **confirm.sh** patterns for installation/removal confirmations
- **choose.sh** demonstrates package selection interfaces with multiple choice support
- **spin.sh** patterns for long-running download/installation operations
- Error handling patterns from workflow scripts applicable to package resolution

### File Structure
```
Total files: 31
Shell scripts: 11 (.sh files)
Animation demos: 10 (.tape files)
Language examples: 3 (JavaScript, Python, Ruby)
Data files: 4 (txt, ansi)
Documentation: 2 (README.md, .gitignore)
Configuration: 1 (.gitignore)
```

Languages represented: Bash/Shell, JavaScript (Node.js), Python, Ruby
External dependencies: git, mods, xsel/pbcopy, vhs, skate