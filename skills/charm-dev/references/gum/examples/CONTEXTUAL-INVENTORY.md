# Gum Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|--------------|
| Multi-step git workflow with confirmations | `examples/commit.sh` |
| Interactive file staging and restoration | `examples/git-stage.sh` |
| Multi-branch operations with command selection | `examples/git-branch-manager.sh` |
| Multi-step user onboarding flow | `examples/demo.sh` |
| Interactive list filtering with key-value pairs | `examples/filter-key-value.sh` |
| System information display with adaptive layouts | `examples/diyfetch` |
| Card selection interface with dynamic rendering | `examples/magic.sh` |
| AI-powered content generation with clipboard integration | `examples/kaomoji.sh` |
| Key-value database browsing | `examples/skate.sh` |
| Video to GIF conversion workflow | `examples/convert-to-gif.sh` |
| POSIX shell single-choice selection | `examples/posix.sh` |
| Comprehensive component demonstration | `examples/test.sh` |
| Node.js process spawning and output capture | `examples/gum.js` |
| Python subprocess integration | `examples/gum.py` |
| Ruby backtick execution with hash mapping | `examples/gum.rb` |

## Examples by Capability

### Git Workflow Automation

**Use commit.sh when you need:**
- Conventional commit message composition
- Multi-step input collection with type selection
- Optional scope handling with conditional formatting
- Confirmation gates before git operations

**File**: `examples/commit.sh`
**Key patterns**: Type selection via `gum choose`, optional scope with test condition, pre-populated input values, confirmation workflow

**Use git-stage.sh when you need:**
- Interactive file staging from git status
- Toggle between add and restore operations
- Multi-file selection with no limit
- Piping selections to git commands via xargs

**File**: `examples/git-stage.sh`
**Key patterns**: Action selection, git status parsing with cut, no-limit choice, xargs integration

**Use git-branch-manager.sh when you need:**
- Multi-branch selection for batch operations
- Command selection (rebase, delete, update)
- Git repository validation
- Styled UI with custom colors
- Branch iteration with loop processing

**File**: `examples/git-branch-manager.sh`
**Key patterns**: Function-based branch selection with limit parameter, git validation, styled borders, color theming via variables, while-read loop for batch operations

### Interactive Workflows

**Use demo.sh when you need:**
- Multi-step user interaction flow
- Sequential input collection (name, secret, preferences)
- Conditional action processing based on selections
- Loading spinners for async operations
- Styled output with horizontal/vertical joins

**File**: `examples/demo.sh`
**Key patterns**: Sequential clear/prompt cycles, no-limit choice for multiple actions, grep-based action detection, spinner variations (line, pulse, monkey), horizontal join for side-by-side layouts

**Use magic.sh when you need:**
- Conditional confirmation gates
- Two-stage selection (category then detail)
- Dynamic styling based on data attributes
- Card/panel rendering with borders
- Vertical and horizontal text alignment

**File**: `examples/magic.sh`
**Key patterns**: Confirmation-based early exit, string extraction with cut/tr, vertical joins for corners, conditional color assignment, aligned text in bordered containers

**Use kaomoji.sh when you need:**
- External command integration (AI/API)
- Help text handling via case statement
- Variable output filtering and selection
- Clipboard operations with platform detection
- Error handling with empty output checks

**File**: `examples/kaomoji.sh`
**Key patterns**: Command existence checking, xsel/pbcopy detection, cut-based list parsing, variable null checks for error handling

### Data Filtering and Selection

**Use filter-key-value.sh when you need:**
- Key-value pair filtering
- Two-column data structure handling
- Filter-first then extract pattern
- Colon-delimited parsing
- Grep-based value extraction

**File**: `examples/filter-key-value.sh`
**Key patterns**: Heredoc for multi-line data, cut for column extraction, filter for key selection, grep for value lookup

**Use skate.sh when you need:**
- Database/collection browsing workflow
- Two-stage filter (collection then item)
- Key-only listing with filter
- Xargs-based value retrieval

**File**: `examples/skate.sh`
**Key patterns**: Database list selection, keys-only flag usage, filter for fuzzy search, xargs with -I for parameter substitution

### Display and Layout

**Use diyfetch when you need:**
- Terminal size-aware adaptive layouts
- Multiple layout fallback strategies
- ANSI color code rendering
- System information gathering
- Horizontal and vertical composition
- Height/width boundary detection

**File**: `examples/diyfetch`
**Key patterns**: stty size parsing, no_color regex for measurement, max-line-length checking, multiple layout attempts with fallback, centered alignment

**Use test.sh when you need:**
- Component capability reference
- Styling option examples (foreground, cursor, borders)
- Limit variations (single, multiple, unlimited)
- Input types (text, password, write)
- Format types (markdown, code, emoji, template)
- Spinner variations

**File**: `examples/test.sh`
**Key patterns**: Comprehensive flag demonstrations, placeholder text, cursor customization, border styles, color codes, height/width constraints

### Media Processing

**Use convert-to-gif.sh when you need:**
- File selection via filter
- Numeric input collection with defaults
- Basename extraction and extension handling
- Spinner for long-running processes
- FFmpeg integration with variable parameters

**File**: `examples/convert-to-gif.sh`
**Key patterns**: Placeholder prompts with default values, basename manipulation, parameter expansion for extension removal, spin with command execution

### Cross-Language Integration

**Use gum.js when you need:**
- Node.js child process spawning
- Stream piping to stderr
- Data event handling
- Stdout parsing and trimming

**File**: `examples/gum.js`
**Key patterns**: spawn with argument spreading, stderr.pipe, stdout.on('data'), toString().trim()

**Use gum.py when you need:**
- Python subprocess execution
- Stdout capture with text mode
- Result object stdout property access

**File**: `examples/gum.py`
**Key patterns**: subprocess.run with stdout=PIPE, text=True flag, stdout.strip() for cleanup

**Use gum.rb when you need:**
- Ruby backtick command execution
- Hash-based value mapping
- Multi-item selection with limit
- Conditional output based on array length
- Chomp for newline removal

**File**: `examples/gum.rb`
**Key patterns**: Backtick execution with chomp, hash.freeze for constants, split("\n") for arrays, conditional length-based formatting

### Minimal Examples

**Use posix.sh when you need:**
- POSIX shell compatibility
- Single-choice selection
- Basic choose implementation

**File**: `examples/posix.sh`
**Key patterns**: #!/bin/sh shebang, basic choose with string arguments

## Implementation Patterns

### Input Collection
- **Single text**: `gum input --placeholder "text" --prompt "label: "`
- **Password**: `gum input --password`
- **Multi-line**: `gum write --placeholder "text"`
- **Pre-populated**: `gum input --value "existing text"`

### Choice Selection
- **Single**: `gum choose item1 item2 item3`
- **Multiple limited**: `gum choose --limit N item1 item2`
- **Unlimited**: `gum choose --no-limit item1 item2`
- **Custom prefixes**: `--cursor-prefix "(•)" --selected-prefix "(x)"`

### Filtering
- **Stdin pipe**: `echo -e "line1\nline2" | gum filter`
- **File input**: `cat file.txt | gum filter`
- **Placeholder**: `gum filter --placeholder "search..."`
- **Custom indicator**: `gum filter --indicator ">"`

### Confirmation
- **Default yes**: `gum confirm "message"`
- **Default no**: `gum confirm "message" --default=false`
- **Custom labels**: `gum confirm --affirmative "OK" --negative "Cancel"`

### Styling
- **Borders**: `gum style --border normal|double|rounded`
- **Colors**: `--foreground N --border-foreground N` (256-color codes)
- **Spacing**: `--padding "V H" --margin N`
- **Dimensions**: `--width N --height N`
- **Alignment**: `--align left|center|right`

### Layout Composition
- **Horizontal**: `gum join --horizontal item1 item2`
- **Vertical**: `gum join --vertical item1 item2`
- **Centered**: `gum join --horizontal --align center item1 item2`
- **Mixed**: Nest joins for complex layouts

### Process Indicators
- **Basic**: `gum spin -- command`
- **Custom spinner**: `gum spin --spinner minidot|pulse|monkey`
- **Title**: `gum spin --title "Loading..."`
- **Show output**: `gum spin --show-output -- command`

### String Manipulation
- **Column extraction**: `cut -d':' -f1` (key from key:value)
- **Character removal**: `tr -d '()'`
- **Newline splitting**: `split("\n")` or `tr " " "\n"`
- **Basename**: `basename "$path"`, extension: `${var%%.*}`

### External Integration
- **Piping**: `command | gum filter | command`
- **Xargs**: `gum choose | xargs command`
- **Xargs with placeholder**: `gum filter | xargs -I {} command {}`
- **Variable capture**: `VAR=$(gum choose ...)`

### Error Handling
- **Empty check**: `[ -z "$var" ] && exit 1`
- **Command existence**: `command -v cmd &> /dev/null`
- **Confirmation gate**: `gum confirm "?" && command || exit`
- **Git validation**: `git rev-parse --git-dir > /dev/null 2>&1`

### Conditional Logic
- **String match**: `[[ "$var" == "value" ]]`
- **Grep in variable**: `grep -q "pattern" <<< "$var"`
- **Case statement**: `case "$1" in "-h"|"--help") ... esac`
- **Color by content**: `if [[ condition ]]; then COLOR=N; else COLOR=M; fi`

### Data Sources
- **Heredoc**: `export VAR=$(cat <<END\nline1\nline2\nEND\n)`
- **Sequence**: `echo {1..500} | sed 's/ /\n/g'`
- **Command output**: `git branch --format="%(refname:short)"`
- **File lines**: `cat file.txt`

### Cross-Platform
- **Clipboard (X11)**: `xclip -sel clip` or `xsel`
- **Clipboard (macOS)**: `pbcopy`
- **Detection**: `command -v xsel &> /dev/null`
- **Terminal size**: `stty size` → `height=${size% *}` `width=${size#* }`

## VHS Demo Files

Animation scripts for documentation (10 .tape files):
- `choose.tape` - Choice selection demos
- `commit.tape` - Commit workflow animation
- `confirm.tape` - Confirmation dialog demos
- `customize.tape` - Customization examples
- `demo.tape` - Full demo walkthrough
- `file.tape` - File picker demos
- `input.tape` - Input field examples
- `pager.tape` - Pager functionality
- `spin.tape` - Spinner animations
- `write.tape` - Multi-line write demos

**Pattern**: VHS tape scripts for generating documentation GIFs, not runtime examples

## Data Files

- `fav.txt` - Sample favorite items list
- `flavors.txt` - Sample flavor options
- `story.txt` - Sample story text
- `format.ansi` - ANSI escape sequence examples

**Pattern**: Test data for filtering and display demonstrations
