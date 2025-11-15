# Charmbracelet Examples Inventory

Comprehensive collection of examples from the entire Charmbracelet ecosystem, extracted and organized for TUI development reference.

## Overview

This repository contains **762 files** across **18 example directories** from Charmbracelet's TUI ecosystem:
- Bubbletea (47 examples) - Core TUI framework
- Gum (31 examples) - CLI components
- VHS (106+ examples) - Terminal recording
- Huh (40 examples) - Interactive forms
- Lipgloss (20 examples) - Styling system
- Wish (15 examples) - SSH framework
- And 12 more specialized tools

Each directory contains an `INVENTORY.md` with practical assessments focused on real-world applications.

## Target Applications

These examples provide patterns for:
- **NextJS Scaffolder** with auth/shadcn templating
- **Inter-Agent Communication** systems for coding agents
- **GitHub Search Tools** with TUI interfaces
- **Package Distribution TUIs** for brew/npm/bun deployment

---

## anthropic-sdk-go



---

## bubbletea



---

## colorprofile



---

## glamour



---

## glamour



---

## go-genai



---

## gum



---

## harmonica



---

## huh



---

## lipgloss



---

## log



---

## pop



---

## ultraviolet



---

## vhs-action



---

## vhs



---

## vhs



---

## wish



---

## x



---


## Development Reference Structure

```
texamples/
├── anthropic-sdk-go/examples/    # AI SDK integration patterns
├── bubbletea/examples/           # Core TUI framework (47 examples)
├── gum/examples/                 # CLI components (31 examples)
├── huh/examples/                 # Interactive forms (40 examples)
├── lipgloss/examples/            # Styling system (20 examples)
├── vhs/examples/                 # Terminal recording (106+ examples)
├── wish/examples/               # SSH framework (15 examples)
└── [12 more directories]        # Additional specialized tools
```

## Usage

1. **Browse Examples**: Each directory contains working code examples
2. **Reference Patterns**: Use INVENTORY.md files for implementation guidance
3. **Copy Templates**: Examples are ready-to-use starting points
4. **Adapt for Projects**: Practical applications mapped to common use cases

## Contributing

This is a curated snapshot of Charmbracelet examples. For updates to source examples, contribute to the original repositories at [github.com/charmbracelet](https://github.com/charmbracelet).

Generated from Charmbracelet ecosystem repositories as of September 2025.

## anthropic-sdk-go

## Examples Catalog

### Core Messaging
- **message** - Basic message API implementation with simple text input/output
- **message-streaming** - Streaming message responses with real-time token processing
- **message-mcp-streaming** - Model Context Protocol integration with external server communication

### Tool Integration
- **tools** - Function calling with three example tools (coordinates, weather, temperature units)
- **tools-streaming** - Streaming function calls with real-time tool execution feedback
- **tools-streaming-jsonschema** - Tool calling with automatic JSON schema generation from Go structs

### Multimodal & File Handling
- **multimodal** - Image analysis with base64 encoding for PNG input processing
- **file-upload** - Document upload via beta Files API with text analysis capabilities

### Platform Integrations
- **bedrock** - AWS Bedrock integration with standard message processing
- **bedrock-streaming** - AWS Bedrock with streaming response handling
- **vertex** - Google Vertex AI integration for cloud deployment
- **vertex-streaming** - Google Vertex AI with streaming capabilities

## Implementation Patterns

### Common Structures
- Context-based request handling using `context.TODO()`
- Anthropic client initialization with optional configuration
- Message parameter construction using builder patterns
- Error handling with panic-based failure management
- Color-coded terminal output for user/assistant distinction

### Tool Implementation
- Interface-based tool parameter definition
- JSON marshaling/unmarshaling for tool inputs/outputs
- Loop-based conversation flow with tool result injection
- Hardcoded mock responses for demonstration purposes

### Stream Processing
- Event-based streaming with type switching
- Message accumulation for complete response reconstruction
- Real-time output rendering during response generation
- Error propagation through stream error checking

### Platform Adaptations
- AWS SDK v2 integration for Bedrock access
- Model name variations for platform-specific endpoints
- Configuration loading for cloud service authentication
- Service-specific parameter formatting


---

## bubbletea

## Examples Catalog

### Form and Input Components
- **autocomplete** - GitHub API integration with text input autocomplete using real-time HTTP requests
- **credit-card-form** - Multi-field form with validation, field navigation, and input constraints
- **textinput** - Single text input field with basic validation
- **textinputs** - Multiple text inputs with focus management and cursor mode switching
- **textarea** - Multi-line text area for longer content input
- **split-editors** - Multiple textarea components with focus switching

### Data Display and Navigation
- **list-default** - Basic list component with item selection and filtering
- **list-fancy** - Enhanced list with custom styling, item management, and multiple UI toggles
- **list-simple** - Minimal list implementation with compact appearance
- **table** - Tabular data display with row selection and keyboard navigation
- **table-resize** - Table component with dynamic column resizing
- **tabs** - Tab navigation system with styled borders and content switching

### File and System Operations
- **file-picker** - Directory and file browser with type filtering and selection
- **exec** - External command execution integration within TUI
- **pipe** - Shell pipe communication for data input/output

### Progress and Status Indicators
- **package-manager** - Progress tracking for multi-step installation processes with spinners
- **progress-animated** - Animated progress bar with dynamic updates
- **progress-download** - File download with real-time progress indication
- **progress-static** - Static progress bar with manual increment control
- **spinner** - Single loading spinner implementation
- **spinners** - Multiple spinner types and styles demonstration

### Network and Communication
- **http** - HTTP request integration with status reporting
- **send-msg** - External message injection into running TUI applications
- **realtime** - Go channel communication for live data updates
- **chat** - Multi-line chat interface with viewport and input areas

### Utility and Control
- **help** - Context-sensitive help system with key binding documentation
- **timer** - Countdown timer with start/stop controls
- **stopwatch** - Elapsed time tracking with lap functionality
- **debounce** - Input throttling to prevent excessive processing
- **mouse** - Mouse event handling and cursor interaction
- **result** - Selection menu with choice confirmation

### Layout and Display
- **altscreen-toggle** - Alternative screen buffer switching
- **fullscreen** - Full terminal screen utilization
- **views** - Multiple view management with switching between different interfaces
- **composable-views** - Component composition with spinner and timer integration
- **tabs** - Tab-based navigation with content areas
- **window-size** - Dynamic window size handling and responsive layout

### Advanced Patterns
- **tui-daemon-combo** - Hybrid TUI/daemon mode operation
- **glamour** - Markdown rendering within viewport component
- **pager** - Text pagination similar to less command
- **paginator** - List pagination with page navigation
- **cellbuffer** - Low-level terminal cell manipulation
- **sequence** - Command sequencing and batch operations
- **prevent-quit** - Application quit prevention and confirmation
- **suspend** - Terminal suspension and restoration handling
- **focus-blur** - Focus state management across components
- **set-window-title** - Terminal window title manipulation
- **simple** - Minimal TUI implementation template

## Implementation Patterns

### State Management
- Model-View-Update architecture with tea.Model interface
- Message passing for event handling and state updates
- Command batching for multiple simultaneous operations

### Input Handling
- Keyboard event mapping with tea.KeyMsg processing
- Mouse interaction support for click and drag operations
- Custom key binding systems with help documentation

### Component Integration
- Bubble composition for complex multi-component interfaces
- Focus management between multiple interactive elements
- External process integration with command execution

### Styling and Layout
- Lipgloss styling framework for colors, borders, and positioning
- Responsive layout handling for terminal resizing
- Theme and style customization patterns

### Network Integration
- HTTP request handling with progress indication
- Real-time data updates through goroutines and channels
- API integration with JSON parsing and error handling

## Bubble Tea API Features

**Recent Updates (v0.23.0 - v2.0 Beta)** - See [BUBBLETEA_API_FEATURES.md](./BUBBLETEA_API_FEATURES.md) for complete documentation

### Program Control

#### tea.Interrupt vs tea.Quit (v1.3.0)
Graceful interrupt handling for production applications.

```go
case tea.KeyMsg:
    if msg.Type == tea.KeyCtrlC {
        return m, tea.Interrupt  // Graceful shutdown
    }
```

### Message Filtering (v0.24.0)

#### tea.WithFilter()
Filter messages before processing - essential for modal dialogs and preventing unwanted quits.

```go
filter := func(m tea.Model, msg tea.Msg) tea.Msg {
    if model.modalOpen {
        if _, ok := msg.(tea.QuitMsg); ok {
            return nil  // Block quit when modal is open
        }
    }
    return msg
}
p := tea.NewProgram(model, tea.WithFilter(filter))
```

### Focus/Blur Events (v1.1.0)

Track window focus for better UX. Example: [`focus-blur`](./bubbletea/examples/focus-blur/)

```go
p := tea.NewProgram(model, tea.WithReportFocus())

case tea.FocusMsg:
    m.hasFocus = true
case tea.BlurMsg:
    m.hasFocus = false
```

**Note:** Requires tmux config: `set-option -g focus-events on`

### Suspend/Resume (v0.27.0)

Handle Ctrl+Z programmatically. Example: [`suspend`](./bubbletea/examples/suspend/)

```go
case tea.KeyMsg:
    if msg.Type == tea.KeyCtrlZ {
        return m, tea.Suspend
    }
case tea.ResumeMsg:
    m.status = "Resumed from suspend"
```

### Bracketed Paste (v0.26.0)

**Performance boost for pasting large text** (SQL queries, code blocks, etc.)

```go
// v1.x - Check KeyMsg.Paste
case tea.KeyMsg:
    if msg.Paste {
        m.content += msg.String()
    }

// v2.x - Dedicated message types
case tea.PasteMsg:
    m.text += string(msg)
```

**Disable for sensitive input:**
```go
return m, tea.DisableBracketedPaste
```

### Testing with teatest (v0.24.0)

Official testing package for Bubble Tea applications.

```go
import "github.com/charmbracelet/x/exp/teatest"

func TestApp(t *testing.T) {
    tm := teatest.NewTestModel(t, model{},
        teatest.WithInitialTermSize(300, 100))

    teatest.WaitFor(t, tm.Output(),
        func(bts []byte) bool {
            return bytes.Contains(bts, []byte("ready"))
        })

    tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
    tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))
}
```

### Additional Features

- **Window Title:** `tea.SetWindowTitle("My App")` - Set terminal title
- **FPS Control:** `tea.WithMaxFPS(30)` - Control rendering framerate
- **Environment Vars:** `tea.WithEnvironment(env)` - Essential for SSH/Wish
- **Custom Outputs:** `tea.WithOutput(os.Stderr)` - Render to stderr/buffers
- **Extended Mouse:** Mouse tracking across entire window (v0.25.0)
- **Windows Improvements:** Better input handling and resize events (v0.26.0)

### v2.0 Beta Features

**Split keyboard enhancements** - More granular control:
- `WithKeyReleases()` - Key release events
- `WithUniformKeyLayout()` - Uniform layout format
- `RequestKeyDisambiguation()` - Disambiguated key events

**Debug mode:** Set `TEA_DEBUG=1` to dump panic traces

---

## colorprofile

## Examples Catalog

### Profile Example (`profile/main.go`)
- **File count**: 1 Go source file (76 lines)
- **Language**: Go 1.23.1
- **Functionality**: Demonstrates color profile detection, color conversion between different terminal capabilities, and automatic ANSI color degradation
- **Key features**:
  - Detects terminal color capabilities (TrueColor, ANSI256, ANSI, ASCII, NoTTY)
  - Converts RGB colors to appropriate terminal color formats
  - Shows color degradation from true color to 4-bit ANSI to ASCII-only
  - Interactive demonstration with explanatory text output

### Writer Example (`writer/writer.go`)
- **File count**: 1 Go source file (24 lines)
- **Language**: Go 1.23.1
- **Functionality**: Provides a command-line utility that reads ANSI-colored input from stdin and outputs degraded colors based on terminal capabilities
- **Key features**:
  - Pipes stdin to stdout with automatic color conversion
  - Acts as a filter for ANSI color sequences
  - Handles color degradation transparently

## Implementation Patterns

### Color Profile Detection
- Uses environment variable inspection combined with stdout analysis
- Provides fallback mechanisms for unsupported terminals
- Implements automatic capability detection without user configuration

### Color Conversion Pipeline
- RGB to terminal-specific color format conversion
- Graceful degradation from high-color to low-color terminals
- ANSI escape sequence processing and modification

### Writer Interface Implementation
- Standard Go io.Writer pattern for color processing
- Environment-aware initialization
- Stream processing for real-time color conversion


---

## glamour

## Examples Catalog

### helloworld (20 lines)
Basic markdown rendering demonstration. Takes hardcoded markdown string and renders to terminal using dark theme. Shows fundamental Glamour usage pattern of `glamour.Render(input, style)` for simple markdown-to-terminal output.

### custom_renderer (22 lines)
Terminal renderer configuration example. Demonstrates custom renderer creation with word wrapping at 40 characters and dark style theme. Uses `glamour.NewTermRenderer()` with configuration options for controlled output formatting.

### stdin-stdout (46 lines)
Command-line markdown processor. Reads markdown from stdin, renders through Glamour with auto-style detection and 80-character word wrap, outputs to stdout. Includes error handling and can be used as unix pipeline component.

### stdin-stdout-custom-styles (46 lines)
Alternative stdin processor using environment-based configuration. Similar to stdin-stdout but uses `glamour.WithEnvironmentConfig()` instead of `glamour.WithAutoStyle()` for theme detection from terminal environment.

### artichokes (57 lines)
Color profile demonstration with embedded content. Uses Go embed directive to include artichokes.md file, demonstrates color profile detection and downsampling through colorprofile package. Shows terminal color capability detection and adaptive rendering.

## Implementation Patterns

### Rendering Approaches
- Simple string rendering: `glamour.Render(string, style)`
- Configured renderer: `glamour.NewTermRenderer(options...)`
- Byte-based processing: `renderer.RenderBytes([]byte)`

### Configuration Options
- Style themes: "dark", auto-detection, environment-based
- Word wrapping: Fixed width (40, 80 characters)
- Color handling: Auto-detection, manual profile specification

### Input Sources
- Hardcoded strings in source
- Standard input streams (stdin)
- Embedded files using `//go:embed`
- File system reads with `io.ReadAll()`

### Error Handling
- Stderr output for processing errors
- Exit codes for failure conditions
- Graceful degradation for unsupported features


---

## glamour

## Examples Catalog

### Text Formatting
- **block_quote**: Demonstrates quote block rendering with indentation and styling
- **code**: Inline code snippet formatting with background highlighting
- **code_block**: Multi-line code block rendering with syntax highlighting support
- **emph**: Italic text emphasis rendering
- **strong**: Bold text emphasis rendering
- **strikethrough**: Text strikethrough formatting

### Document Structure
- **heading**: Hierarchical heading rendering (h1, h2, h3) with prefix/suffix customization
- **hr**: Horizontal rule rendering for document separation
- **paragraph**: Basic paragraph text formatting

### Lists and Tasks
- **enumeration**: Numbered list rendering with customizable numbering styles
- **list**: Unordered list rendering with bullet point customization
- **ordered_list**: Sequential numbered list implementation
- **task**: Checkbox task list rendering for TODO items

### Links and Media
- **emoji**: Emoji character rendering in terminal output
- **image**: Image placeholder rendering with alt-text display
- **link**: Hyperlink rendering with URL highlighting

### Table Rendering (13 variations)
- **table**: Basic table structure with headers and rows
- **table_align**: Table column alignment demonstration
- **table_truncate**: Table content truncation for width management
- **table_wrap**: Table text wrapping within cells
- **table_with_footer_auto_links**: Auto-generated footer links
- **table_with_footer_auto_links_short**: Shortened footer link format
- **table_with_footer_images**: Image references in table footers
- **table_with_footer_images_no_alt**: Image handling without alt text
- **table_with_footer_images_same_alt**: Duplicate alt text handling
- **table_with_footer_links**: Reference links in table footers
- **table_with_footer_links_complex**: Multi-developer profile table with links and images
- **table_with_footer_links_huge**: Large reference link collections
- **table_with_footer_links_no_color**: Colorless link rendering
- **table_with_footer_links_repeated**: Duplicate link reference handling
- **table_with_inline_links**: Direct inline link embedding

## Implementation Patterns

### Style Configuration Pattern
Each example follows a consistent structure:
- `.md` file: Contains markdown content to render
- `.style` file: JSON configuration defining colors, margins, prefixes, and formatting
- `.png` file: Screenshot showing expected terminal output (when available)

### Style Properties
Style files use JSON objects with element-specific configuration:
- Color values (numeric color codes)
- Background colors
- Margins and indentation
- Text prefixes and suffixes
- Font styling (bold, italic)

### Table Rendering Complexity
Tables demonstrate the most complex rendering scenarios with:
- Footer reference management
- Image placeholder handling
- Link processing and deduplication
- Content truncation and wrapping
- Multi-column alignment


---

## go-genai

## Examples Catalog

### Chat Systems (`chats/`)
- **chat.go**: Basic chat session implementation with temperature configuration
- **chat_stream.go**: Streaming chat with real-time response processing
- **Files**: 2 Go files, 1 README
- **Function**: Creates persistent chat sessions, handles streaming responses

### Live Streaming (`live/`)
- **live_streaming_server.go**: WebSocket server for real-time AI interactions (6,881 lines)
- **live_streaming.html**: Frontend interface with WebSocket client implementation
- **Files**: 1 Go file, 1 HTML file, 1 README
- **Function**: Real-time bidirectional communication between browser and AI model via WebSockets

### File Management (`files/`)
- **upload_file.go**: File upload and processing pipeline
- **list_download.go**: File enumeration and retrieval operations
- **Files**: 2 Go files, 1 README
- **Function**: Handles file uploads to AI service, downloads processed results

### Tool Integration (`mcptoolbox/`)
- **mcp_toolbox.go**: MCP Toolbox SDK integration for database operations
- **go.mod/go.sum**: Dependencies for MCP integration
- **Files**: 1 Go file, 2 module files, 1 README
- **Function**: Converts MCP tools to GenAI-compatible format, handles database tool invocation

### Content Generation (`models/generate_content/`)
- **function_declaration_json_schema.go**: JSON schema-based tool definitions
- **function_declaration_schema.go**: Programmatic tool schema creation
- **text.go**: Basic text generation
- **text_stream.go**: Streaming text generation
- **image_modality.go**: Image analysis and processing
- **image_modality_stream.go**: Streaming image processing
- **audio_modality.go**: Audio content analysis
- **Files**: 7 Go files, 1 README
- **Function**: Demonstrates multimodal content generation with tool calling

### Token Operations (`models/count_tokens/`, `models/compute_tokens/`)
- **text_tokens.go**: Token counting and computation utilities
- **Files**: 1 Go file per directory, 2 READMEs
- **Function**: Token usage analysis and cost estimation

### Image Processing (`models/edit_image/`, `models/recontext_image/`, `models/segment_image/`, `models/upscale_image/`)
- **image.go**: Core image manipulation functions
- **ingredients.go**: Recipe-based image editing
- **segment_image.go**: Image segmentation operations
- **Files**: 4-6 Go files across directories
- **Function**: Image editing, enhancement, segmentation, and upscaling

### Media Generation (`models/generate_images/`, `models/generate_videos/`)
- **images.go**: Image generation from text prompts
- **videos.go**: Video generation from text descriptions
- **Files**: 1 Go file per directory, 2 READMEs
- **Function**: Creates visual media from text descriptions

### Content Processing (`models/embed_content/`)
- **text.go**: Text embedding generation for semantic analysis
- **Files**: 1 Go file, 1 README
- **Function**: Converts text to vector embeddings

### Batch Operations (`batches/`)
- **create_get_cancel.go**: Batch job management operations
- **Files**: 1 Go file, 1 README
- **Function**: Handles asynchronous batch processing tasks

### Caching System (`caches/`)
- **create_get_delete.go**: Cache management operations
- **list.go**: Cache enumeration functionality
- **Files**: 2 Go files, 1 README
- **Function**: Content caching for performance optimization

### Model Tuning (`tunings/`)
- **create_get.go**: Model fine-tuning operations
- **list.go**: Tuning job listing and management
- **Files**: 2 Go files, 1 README
- **Function**: Custom model training and management

### Resource Management (`models/resource_methods/`)
- **get.go**: Resource retrieval operations
- **list.go**: Resource enumeration
- **Files**: 2 Go files, 1 README
- **Function**: API resource management and inspection

## Implementation Patterns

### Authentication Configuration
- Environment-based backend selection (Vertex AI vs Gemini API)
- API key and project configuration through environment variables
- Consistent authentication pattern across all examples

### Client Initialization
- Standard `genai.NewClient(ctx, config)` pattern
- Backend detection and appropriate initialization
- Error handling with fatal logging

### Content Generation
- Temperature-based response control
- Streaming vs non-streaming response handling
- Multimodal input processing (text, image, audio)

### Tool Integration
- JSON schema-based tool definitions
- Function declaration conversion utilities
- Tool execution and response handling

### WebSocket Communication
- Gorilla WebSocket for real-time communication
- Bidirectional message proxying
- Embedded HTML templates

### Error Handling
- Consistent error checking with `log.Fatal()`
- Context-based cancellation support
- Resource cleanup patterns


---

## gum

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


---

## harmonica

## Examples Catalog

### Particle Example (`particle/main.go`)
- **Type**: Terminal User Interface (TUI) animation
- **Functionality**: Projectile physics simulation with gravity
- **Implementation**: 88 lines of Go code using BubbleTea framework
- **Physics**: Demonstrates terminal gravity acceleration on a moving projectile
- **Interaction**: Terminates on any key press or when projectile reaches maximum height
- **Output**: ASCII coordinates display showing real-time position updates at 60 FPS

### Spring TUI Example (`spring/tui/main.go`)
- **Type**: Terminal User Interface animation with styled output
- **Functionality**: Spring physics simulation with damping
- **Implementation**: 111 lines of Go code using BubbleTea and Lipgloss libraries
- **Physics**: Spring oscillation with configurable frequency (7.0) and damping (0.15)
- **Visual**: Colored sprite block that moves horizontally toward target position
- **Interaction**: Automatic termination when spring settles near target position

### Spring OpenGL Example (`spring/opengl/main.go`)
- **Type**: Interactive OpenGL graphics application
- **Functionality**: Real-time interactive spring physics with mouse control
- **Implementation**: 325 lines of Go code using Pixel graphics library and gg rendering
- **Physics**: User-adjustable spring frequency and damping parameters
- **Interaction**: Mouse click targets, keyboard controls for parameter adjustment
- **Visual**: Circular sprites with smooth spring animations and help text overlay
- **Dependencies**: Requires JetBrains Mono font file for text rendering

### Support Files
- **go.mod**: Module definition with dependencies for BubbleTea, Harmonica, Lipgloss, Pixel, and gg libraries
- **go.sum**: Dependency checksums for package verification
- **JetBrainsMono-Regular.ttf**: TrueType font file (18 tables, digitally signed)

## Implementation Patterns

### Common Architecture
- **Framework**: BubbleTea Model-View-Update pattern for TUI applications
- **Animation**: Consistent 60 FPS frame timing using time-based message passing
- **Physics Integration**: Harmonica library for spring mechanics and projectile calculations
- **State Management**: Immutable state updates through message handlers

### Rendering Approaches
- **ASCII Text**: Position-based spacing for simple coordinate display
- **Styled Terminal**: Lipgloss for colored blocks and formatted text output
- **OpenGL Graphics**: Pixel library with gg context for hardware-accelerated rendering

### Input Handling
- **Keyboard**: Universal quit-on-keypress pattern for TUI examples
- **Mouse**: Click-to-target positioning in OpenGL implementation
- **Real-time**: Arrow key parameter adjustment with shift modifier support


---

## huh

## Examples Catalog

### Core Form Components
- **accessibility**: Basic form with accessibility settings via ACCESSIBLE environment variable
- **accessibility-secure-input**: Secure input handling with accessibility features
- **bubbletea**: Full bubbletea integration with custom styling and state management for job application form
- **bubbletea-options**: Extended bubbletea options demonstration
- **filepicker**: File selection interface with type filtering (.png, .jpeg, .webp, .gif)
- **filepicker-picking**: Enhanced file picker with selection capabilities
- **skip**: Form field skipping logic demonstration
- **timer**: Time-based form operations
- **help**: Help system integration and display
- **hide**: Conditional field visibility controls
- **scroll**: Scrollable content handling in forms

### Multi-Component Forms
- **burger**: Complex multi-step ordering system with validation, spinners, and receipt generation
- **multiple-groups**: Form group management and navigation
- **conditional**: Field dependency and conditional display logic
- **stickers**: Multi-selection interface with visual feedback

### Layout Systems
- **layout/columns**: Column-based form layouts
- **layout/stack**: Vertical stacking arrangements
- **layout/grid**: Grid-based component positioning
- **layout/default**: Standard layout patterns

### Dynamic Content Generation
- **dynamic/dynamic-all**: Complete dynamic form generation
- **dynamic/dynamic-bubbletea**: Dynamic content within bubbletea framework
- **dynamic/dynamic-count**: Counter-based dynamic field generation
- **dynamic/dynamic-country**: Geographic data selection with dynamic options
- **dynamic/dynamic-increment**: Incremental field addition
- **dynamic/dynamic-markdown**: Markdown content rendering in forms
- **dynamic/dynamic-name**: Name-based dynamic field generation
- **dynamic/dynamic-suggestions**: Repository search with organization-based autocomplete

### Network and Authentication
- **ssh-form**: SSH server hosting forms over network connections (localhost:2222)
- **gh/create**: GitHub workflow simulation with push/fork/PR creation flow
- **git**: Git operation interfaces
- **gum**: Integration with gum CLI tool styling

### Theming and Styling
- **theme**: Custom theme application and styling examples

### Documentation Examples
- **readme/input**: Input field demonstrations
- **readme/main**: Primary readme example implementation
- **readme/confirm**: Confirmation dialog patterns
- **readme/multiselect**: Multiple selection interfaces
- **readme/note**: Note and information display
- **readme/select**: Selection menu examples
- **readme/text**: Text input and display handling

## Implementation Patterns

### Form Structure
- Group-based organization with huh.NewGroup()
- Field chaining with method calls
- Validation functions attached to fields
- State management through struct binding

### Input Validation
- Custom validation functions with error returns
- Field-specific constraints (length, format, selection limits)
- Cross-field dependency validation

### Dynamic Content
- Function-based placeholder generation
- Context-dependent suggestion systems
- Real-time option filtering based on other field values

### Network Integration
- SSH server middleware for remote form access
- Bubbletea program hosting over network
- Session management and authentication flows

### Styling Architecture
- Lipgloss renderer integration
- Adaptive color schemes (light/dark)
- Component-specific styling with inheritance
- Theme system with customizable components


---

## lipgloss

## Examples Catalog

### Layout (1 example)
- **layout** - Complex demo showing tabs, color grids, dialogs, status bars, and multi-column layouts. Demonstrates advanced styling patterns including adaptive colors, borders, and terminal-aware rendering.

### List (6 examples)
- **list/simple** - Basic nested list with Roman numeral enumerators
- **list/duckduckgoose** - Custom enumerator demonstration
- **list/glow** - Markdown-styled list formatting
- **list/grocery** - Interactive-style list with checkmarks and strikethrough for completed items
- **list/roman** - Roman numeral enumeration patterns
- **list/sublist** - Hierarchical list structures with multiple nesting levels

### SSH (1 example)
- **ssh** - SSH server integration showing per-client renderer customization, terminal capability detection, and adaptive color profiles for remote sessions

### Table (5 examples)
- **table/ansi** - ANSI color and formatting demonstration in tabular format
- **table/chess** - Chess game board representation with alternating colors
- **table/languages** - International text display with right-to-left language support
- **table/mindy** - Character styling and formatting showcase
- **table/pokemon** - Data table with conditional styling, row highlighting, and type-based color coding

### Tree (7 examples)
- **tree/simple** - Basic hierarchical tree structure for operating systems
- **tree/files** - Filesystem tree renderer that recursively displays directory contents
- **tree/background** - Tree styling with background color variations
- **tree/makeup** - Cosmetic styling variations for tree components
- **tree/rounded** - Rounded border styling for tree elements
- **tree/styles** - Multiple tree styling patterns and customizations
- **tree/toggle** - Interactive-style tree with expandable/collapsible behavior

## Implementation Patterns

### Common Patterns
- Style definition through `lipgloss.NewStyle()` chains
- Adaptive color schemes for light/dark terminal detection
- Conditional styling based on row/column position or data content
- Terminal width detection and responsive layout adjustment
- Custom renderer creation for specific output targets

### Technical Implementations
- **Border Styles**: Normal, thick, rounded borders with custom foreground colors
- **Color Management**: Hex colors, adaptive colors, color blending, and gradients
- **Layout Techniques**: Horizontal/vertical joining, centering, padding, margins
- **Data Rendering**: Dynamic styling functions based on content or position
- **Terminal Integration**: SSH session handling, PTY management, and environment detection


---

## log

## Examples Catalog

### Basic Logging
- **log/log.go** (7 lines): Minimal usage example showing basic Print function
- **cookie/cookie.go** (8 lines): Simple Debug and Info level logging demonstration
- **new/new.go** (12 lines): Creating custom logger instances with output redirection

### Error Handling
- **error/error.go** (12 lines): Error logging with structured key-value pairs
- **oven/oven.go** (13 lines): Helper function usage and caller reporting

### Configuration and Options
- **options/options.go** (19 lines): Logger configuration including prefix, time format, timestamps, and caller reporting
- **batch2/batch2.go** (15 lines): Logger context creation with persistent key-value pairs
- **chocolate-chips/chocolate-chips.go** (18 lines): Context logger creation and configuration chaining

### Styling and Formatting
- **styles/styles.go** (25 lines): Custom styling with lipgloss integration, error level customization, and key-value styling
- **format/format.go** (15 lines): Loop-based logging for progress tracking scenarios
- **temperature/temperature.go** (15 lines): Incremental value logging with formatting

### Integration
- **slog/main.go** (26 lines): Standard library slog integration with timestamp configuration and UTC time functions
- **app/app.go** (62 lines): Comprehensive example with custom types, helper functions, multiple log levels, and conditional logging

### Recording Files
12 .tape files (7-10 lines each) for generating terminal output recordings using VHS or similar tools.

## Implementation Patterns

### Logger Configuration
- Custom output destinations (os.Stderr redirection)
- Time format customization (time.Kitchen, time.RFC3339)
- Log level management (DebugLevel, InfoLevel, etc.)
- Caller reporting toggle
- Timestamp reporting control

### Structured Logging
- Key-value pair logging pattern throughout examples
- Context logger creation with With() method
- Persistent context values across multiple log calls
- Custom type String() method implementation for logging

### Styling Integration
- Lipgloss styling library integration
- Custom level styling (background, foreground, padding)
- Key and value styling customization
- Error highlighting patterns

### Helper Functions
- log.Helper() usage for accurate caller reporting
- Function-scoped logging contexts
- Error condition logging patterns


---

## pop

## Examples Catalog

### Demonstration Scripts (.tape files)
- **cli.tape** (387 bytes): Command-line email sending with file attachment. Shows piping markdown content to pop with explicit from/to/subject parameters and README.md attachment.
- **demo.tape** (596 bytes): Interactive TUI workflow demonstration. Records step-by-step email composition using pop's form interface for sender, recipient, subject, and message body.
- **gum-example.tape** (404 bytes): Integration with gum selection tools. Demonstrates dynamic recipient selection from contact list and sender choice using gum's choose and filter commands.
- **invoice-example.tape** (366 bytes): Business workflow integration. Shows PDF invoice generation followed by email attachment using pop with custom body text.
- **mods-example.tape** (469 bytes): AI integration demonstration. Uses mods (AI text generation) to create email content and sends via pop with preview mode enabled.

### Supporting Files
- **contacts.txt** (148 bytes): Sample email address list containing 10 @charm.sh addresses for use with contact selection examples.
- **message.md** (63 bytes): Template markdown email content for demonstrations. Contains basic formatting with bold and italic text.
- **README.md** (6 bytes): Minimal documentation containing only "# Pop" header.

## Implementation Patterns

### Email Workflow Patterns
- **Pipeline Integration**: Input redirection and command chaining for automated email sending
- **Interactive Forms**: TUI-based step-through email composition with tab navigation
- **External Tool Integration**: Combining pop with gum for contact selection and mods for content generation
- **File Attachment**: Direct file attachment capability with various document types

### Demonstration Patterns
- **VHS Recording Configuration**: Consistent framerate (28-30 fps), padding (40-70px), and resolution settings across all demos
- **Environment Setup**: API key management using pass command for secure credential handling
- **User Interaction Simulation**: Keyboard input timing, navigation delays, and visual presentation sequences


---

## ultraviolet

## Examples Catalog

### altscreen
- **Files**: 1 Go file (1,873 bytes)
- **Function**: Demonstrates alternate screen buffer switching with inline mode toggle
- **Key features**: Screen mode switching, cursor visibility control, viewport management
- **Input handling**: Space to toggle modes, Ctrl+C/q to exit
- **Application**: Foundation for applications requiring different display modes

### draw
- **Files**: 1 Go file (4,682 bytes)
- **Function**: Interactive drawing application with mouse and keyboard input
- **Key features**: Mouse event handling, character drawing, altscreen buffer, focus events
- **Input handling**: Mouse drawing, keyboard input, character rendering
- **Application**: Base for interactive drawing tools or text editors with mouse support

### helloworld
- **Files**: 1 Go file (2,354 bytes)
- **Function**: Basic terminal application displaying styled text in a fixed area
- **Key features**: Terminal initialization, event loop, styled string rendering, suspend/resume
- **Input handling**: q/Ctrl+C to quit, Ctrl+Z for suspend
- **Application**: Template for basic TUI applications with event handling

### image
- **Files**: 1 Go file (11,095 bytes) + 1 JPEG file (6,300 bytes)
- **Function**: Image rendering with multiple encoding formats
- **Key features**: Supports blocks, sixel, iTerm2, and Kitty image protocols
- **Input handling**: Command-line flags for encoding selection
- **Application**: Image viewers, media applications, terminal graphics

### layout
- **Files**: 1 Go file (13,636 bytes)
- **Function**: Complex layout system with tabs, columns, and styled components
- **Key features**: Lipgloss styling, responsive design, color theme detection, tabbed interface
- **Input handling**: Not specified in examined portion
- **Application**: Dashboard layouts, multi-panel applications, styled interfaces

### panic
- **Files**: 1 Go file (1,491 bytes)
- **Function**: Demonstrates panic recovery and graceful terminal restoration
- **Key features**: Panic handling, countdown timer, terminal cleanup
- **Input handling**: q/Ctrl+C to exit before panic
- **Application**: Error handling patterns, graceful shutdown mechanisms

### prependline
- **Files**: 1 Go file (3,881 bytes)
- **Function**: Line prepending functionality with color cycling
- **Key features**: Dynamic line addition, color backgrounds, frame height management
- **Input handling**: Event-driven line addition
- **Application**: Log viewers, streaming text applications, status displays

### space
- **Files**: 1 Go file (3,161 bytes)
- **Function**: Animated space-like visualization with randomized grayscale patterns
- **Key features**: Color gradients, animation loops, mathematical pattern generation
- **Input handling**: Standard quit commands
- **Application**: Screensavers, visual effects, background animations

### tv
- **Files**: 1 Go file (3,680 bytes)
- **Function**: TV test pattern generator with color bars
- **Key features**: Color bar rendering, predefined color palette, test pattern display
- **Input handling**: Standard terminal events
- **Application**: Display testing, color calibration, retro aesthetics

## Implementation Patterns

### Common Patterns
- **Terminal initialization**: All examples use `uv.DefaultTerminal()` or `uv.NewTerminal()`
- **Event handling**: Context-based cancellation with event channels
- **Error handling**: Consistent error checking with log.Fatalf()
- **Cleanup**: Deferred terminal restoration and shutdown procedures

### Input Management
- **Standard quit**: q, Ctrl+C patterns across examples
- **Context cancellation**: Proper goroutine management
- **Event streaming**: Channel-based event distribution

### Display Techniques
- **Alternate screen**: Most examples use altscreen buffer for full-screen applications
- **Styled rendering**: Integration with Lipgloss for advanced styling
- **Dynamic sizing**: Responsive to terminal size changes

### Architecture
- **Single-file applications**: Each example is self-contained
- **Module structure**: Shared dependencies via go.mod
- **Logging**: Debug logging to separate files


---

## vhs-action

## Examples Catalog

### auto-commit.yml
- **File Count**: 1 YAML workflow file
- **Language**: GitHub Actions YAML
- **Functionality**: Triggers on push events when `vhs.tape` files change, executes VHS recording, commits generated GIF files back to repository
- **Dependencies**: charmbracelet/vhs-action@v1, stefanzweifel/git-auto-commit-action@v4
- **Trigger Pattern**: Push to main branch with changes to vhs.tape files
- **Output**: Automatically commits generated *.gif files to repository

### comment-pr.yml
- **File Count**: 1 YAML workflow file
- **Language**: GitHub Actions YAML
- **Functionality**: Triggers on pull requests with vhs.tape changes, generates GIF, uploads to Imgur, posts GIF link as PR comment
- **Dependencies**: charmbracelet/vhs-action@v1, devicons/public-upload-to-imgur@v2.2.2, github-actions-up-and-running/pr-comment@v1.0.1
- **Trigger Pattern**: Pull request events with vhs.tape file modifications
- **Output**: Imgur-hosted GIF link posted as pull request comment

## Implementation Patterns

### Conditional Workflow Triggers
Both examples use path-based triggering that activates only when specific files (vhs.tape) are modified, preventing unnecessary workflow runs.

### Multi-Step Action Chains
Sequential action execution pattern: checkout → VHS generation → external service integration (Git commit or image hosting + commenting).

### Secret Management
Both workflows require GitHub secrets (GITHUB_TOKEN, IMGUR_CLIENT_ID) for external service authentication.

### File Pattern Matching
Auto-commit example uses glob pattern `*.gif` to target generated files, while comment example targets specific `./vhs.gif` output.


---

## vhs

## Examples Catalog

### Bubble Tea (36 examples)
TUI framework demonstrations in Go showing interactive application patterns:
- **Package manager simulation** - TUI for software distribution interfaces
- **Form components** - Credit card forms, text inputs, text areas
- **List interfaces** - Default, fancy, and simple list implementations
- **Communication patterns** - Send message functionality, inter-component messaging
- **Data presentation** - Tables, tabs, pagination, progress indicators
- **Real-time updates** - Live data feeds, timers, stopwatch functionality
- **Layout management** - Split editors, composable views, fullscreen modes
- **Integration patterns** - HTTP clients, daemon combinations, external process execution

### GitHub CLI (2 examples)
GitHub API interaction demonstrations:
- **Issue management** - Listing and viewing GitHub issues
- **Pull request workflows** - Listing pull requests across all states

### Gum (3 examples)
Shell scripting TUI components for automation:
- **File navigation** - Directory browsing with selection interface
- **Document viewing** - Paged content display with navigation
- **Data tables** - CSV data presentation with column formatting

### CLI UI (7 examples)
Ruby-based CLI interface components:
- **Text formatting** - Color and style output formatting
- **Progress tracking** - Progress bars and spinner animations
- **User interaction** - Text prompts with validation and defaults
- **Layout structures** - Nested frames and hierarchical displays
- **Status indicators** - Real-time status widgets and symbols

### Glow (3 examples)
Markdown viewer and editor demonstrations:
- **Document browsing** - File navigation and content viewing
- **Search functionality** - Content search and filtering
- **Content editing** - In-terminal markdown editing workflows

### Command Demonstrations (13 examples)
Basic terminal interaction recordings:
- **Input methods** - Keyboard input simulation (arrows, backspace, clipboard)
- **Command execution** - Type commands, enter, tab completion
- **Display control** - Hide/show output, comments, clear screen

### Settings Configuration (24 examples)
VHS recording configuration demonstrations:
- **Visual styling** - Font size, family, spacing, themes
- **Layout control** - Width, height, margins, padding, borders
- **Shell compatibility** - Various shell environments (bash, zsh, fish, nu, cmd, pwsh)
- **Recording behavior** - Typing speed, cursor appearance, loop settings

### Additional Categories
- **Error handling** (3 examples) - Parser errors, dimension validation, requirement checks
- **Publishing** (3 examples) - Output format generation and distribution
- **Environment** (1 example) - Environment variable usage
- **Decorations** (1 example) - Visual enhancement techniques
- **Slides** (1 example) - Presentation mode demonstrations
- **Split layouts** (1 example) - Multi-pane interface recordings
- **Neofetch** (6 examples) - System information display with ASCII art colorization
- **jqp** (1 example) - JSON query and processing interface

## Implementation Patterns

### Terminal Recording Framework
- **Tape file structure** - Command sequences with timing and output control
- **Multi-format output** - GIF, MP4, WebM, ASCII text generation
- **Interactive simulation** - Keyboard input, mouse actions, command execution
- **Visual customization** - Fonts, colors, dimensions, styling

### TUI Component Architecture
- **Model-view-update pattern** - State management and event handling (Bubble Tea)
- **Component composition** - Reusable UI elements and layout systems
- **Event-driven interfaces** - User input handling and response patterns
- **Real-time data integration** - Live updates and external data sources

### CLI Tool Integration
- **GitHub API workflows** - Issue and PR management interfaces
- **File system navigation** - Directory browsing and file selection
- **Data visualization** - Tables, progress indicators, formatted output
- **Shell scripting enhancement** - Interactive components for automation


---

## vhs

## Examples Catalog

### Bubble Tea Framework (37 files)
- **Interactive Forms**: Credit card forms, text inputs, multi-field forms, textarea components
- **List Management**: Default lists, fancy lists, simple lists with navigation
- **Progress Indicators**: Animated progress bars, static progress, spinners, timers
- **Layout Components**: Split editors, composable views, fullscreen modes, help systems
- **Communication**: Chat interfaces, message sending, daemon combinations
- **Data Display**: Tables, pagination, results display
- **Process Management**: External command execution, HTTP requests, realtime updates

### CLI UI Framework (9 files)
- **Text Formatting**: Color formatting, symbol rendering, nested frames
- **Interactive Elements**: Text prompts, spinners, progress bars, status widgets
- **Output Control**: Stdout routing, frame management
- **Ruby Integration**: IRB-based examples for CLI UI gem

### GitHub CLI Integration (2 files)
- **Issue Management**: List issues, view issue details
- **Pull Request Operations**: List pull requests with state filtering

### Gum Interactive Components (3 files)
- **File Navigation**: Directory browsing with gum file selector
- **Content Display**: Pager functionality for README viewing
- **Data Presentation**: Table display with CSV data formatting

### Glow Markdown Viewer (4 files)
- **Document Rendering**: Markdown file display and navigation
- **Search Functionality**: Text search within documents
- **Interactive Editing**: Document modification workflows

### VHS Command Reference (22 files)
- **Input Commands**: Type, arrow keys, backspace, tab, space, enter
- **Control Keys**: Ctrl combinations, alt keys, escape sequences
- **Display Control**: Hide/show commands, clipboard operations
- **Comments**: Documentation within tape files

### VHS Settings Configuration (44 files)
- **Visual Settings**: Font size (10, 20, 40px), font family, letter spacing, line height
- **Layout Settings**: Width, height, padding, margin, border radius
- **Shell Configuration**: Bash, zsh, fish, PowerShell, cmd, custom shells
- **Appearance**: Themes, window bars, cursor settings, typing speed
- **Output Control**: Loop offset for GIF animations

### Additional Tools
- **jqp**: JSON processing with interactive interface (1 file)
- **neofetch**: System information display (1 file)
- **Decorations**: Terminal decoration examples (1 file)
- **Error Handling**: Parser errors, dimension validation, requirement checks (3 files)

## Implementation Patterns

### Recording Structure
- Hide/Show commands for clean output control
- Build-execute-cleanup workflow for compiled examples
- Sleep commands for natural timing
- Standard output formatting (GIF 600px width)

### Interactive Navigation
- Arrow key sequences for menu traversal
- Enter confirmations for selections
- Escape sequences for cancellation
- Ctrl+C for process termination

### TUI Component Patterns
- List-based navigation interfaces
- Form input with validation
- Progress indication for long operations
- Multi-pane layouts with split views
- Modal dialogs and help systems

### Command Line Integration
- External tool demonstration (gh, gum, glow)
- Package management workflows
- File system operations
- Git workflow integration


---

## wish

## Examples Catalog

### Basic SSH Servers
- **simple** - Minimal SSH server that prints "Hello, world!" to connecting clients. Uses ED25519 key authentication and basic middleware. Single file implementation.
- **graceful-shutdown** - SSH server with proper signal handling for clean shutdown. Demonstrates context-based shutdown with 30-second timeout.
- **banner** - SSH server that displays a custom banner from banner.txt file before allowing user interaction. Shows middleware for pre-session content.

### Authentication and Identity
- **identity** - SSH server that identifies users by their public keys. Maintains a hardcoded map of known users and greets them by name. Demonstrates key-based user identification.
- **multi-auth** - SSH server supporting multiple authentication methods. Shows implementation of both public key and password authentication mechanisms.

### Interactive Applications
- **bubbletea** - SSH server hosting a Bubble Tea TUI application. Displays terminal information (size, color profile, background) in real-time. Uses alt screen mode.
- **bubbleteaprogram** - SSH server that runs complete Bubble Tea programs over SSH connections. Demonstrates program lifecycle management.
- **bubbletea-exec** - Combines Bubble Tea interface with PTY allocation for running external programs. Shows interaction between TUI and shell execution.
- **multichat** - Multi-user chat application over SSH. Supports concurrent users in shared chat rooms with real-time message broadcasting. Uses textarea and viewport components.

### Command Line Integration
- **cobra** - SSH server that executes spf13/cobra CLI commands. Implements an echo command with reverse flag option. Shows integration of CLI frameworks over SSH.

### File Transfer and Repository Management
- **scp** - SSH server with SCP and SFTP support. Serves files from testdata directory with read-only filesystem handler. Implements secure file transfer protocols.
- **git** - SSH server for Git repository hosting. Provides repository listing interface and handles Git operations (clone, push, fetch). Stores repositories in .repos directory.

### System Integration
- **forward** - SSH server with reverse port forwarding capabilities. Enables network tunneling through SSH connections.
- **pty** - SSH server that allocates pseudo-terminals for interactive shell sessions. Demonstrates PTY handling for terminal applications.
- **exec** - SSH server that executes shell commands and scripts. Includes example.sh script for testing command execution capabilities.

## Implementation Patterns

### Common Architecture
- All examples use net.JoinHostPort for address binding
- Standard ED25519 key generation and management
- Middleware-based request processing pipeline
- Context-based graceful shutdown with signal handling
- Structured logging using charmbracelet/log

### Authentication Patterns
- Public key authentication with ED25519 keys
- User identification through key fingerprinting
- Access control through hardcoded user maps
- Session-based authentication state management

### TUI Integration
- Bubble Tea model-view-update architecture
- Session-specific renderer creation for color profiles
- Alt screen mode for fullscreen applications
- Real-time message passing between concurrent sessions

### File System Operations
- Filesystem abstraction for SCP/SFTP handlers
- Directory traversal with security boundaries
- Read-only and read-write access patterns
- Git repository management through filesystem


---

## x

## Examples Catalog

### cellbuf (3 files)
Interactive terminal screen buffer demonstration with mouse and keyboard input handling. Shows raw terminal mode, alternative screen buffer usage, and real-time rendering with resize support. Useful for building terminal applications requiring direct screen manipulation.

### charmtone (2 files)
Color palette utility for the CharmTone design system. Generates CSS, SCSS, and Vim colorscheme configurations. Includes terminal color guide rendering with gradient blending. Applicable for design system tooling and color palette management.

### faketty (1 file)
PTY (pseudo-terminal) wrapper for running commands with controlled terminal dimensions. Creates fake TTY environment for testing terminal applications with specific row/column configurations. Relevant for terminal emulation and command execution testing.

### img2term (1 file)
Image to terminal converter using Sixel graphics protocol. Reads JPEG/PNG images and outputs terminal-displayable graphics. Useful for terminal-based image viewers and media applications.

### layout (1 file)
Comprehensive Lip Gloss styling demonstration showing advanced layout features, color detection, and responsive terminal UI design. Contains patterns for column layouts, styling, and dark/light theme adaptation.

### mosaic (2 files)
Image-to-ASCII art converter that renders JPEG images as terminal output using character-based graphics. Demonstrates image processing for terminal display with customizable dimensions.

### parserlog (1 file)
ANSI escape sequence parser with detailed logging output. Processes stdin and categorizes ANSI sequences (CSI, ESC, DCS, OSC). Useful for debugging terminal output and understanding ANSI protocol implementation.

### parserlog2 (1 file)
Alternative ANSI sequence decoder using lower-level parsing approach. Provides detailed sequence analysis with parameter extraction and state tracking. Applicable for ANSI protocol analysis and terminal emulator development.

### pen (1 file)
Text wrapping utility using cellbuf PenWriter. Demonstrates automatic text wrapping with ANSI sequence preservation. Relevant for terminal text formatting and content display systems.

### toner (1 file)
Simple stdin-to-stdout processor using the experimental toner package. Minimal implementation for text processing pipelines with toner formatting capabilities.

## Implementation Patterns

### Terminal Management
- Raw terminal mode setup and cleanup
- Alternative screen buffer handling
- Window resize detection and response
- Cross-platform compatibility (Windows exclusions)

### Input/Output Processing
- Event-driven input handling (keyboard, mouse)
- ANSI sequence parsing and generation
- PTY creation and management
- Stream processing with proper error handling

### Styling and Rendering
- Light/dark theme detection and adaptation
- Color palette management and generation
- Layout composition with Lip Gloss
- Image processing for terminal display

### CLI Framework Integration
- Cobra command structure
- Fang enhanced CLI execution
- Flag-based configuration
- Subcommand organization


---

Generated programmatically from 18 INVENTORY.md files - Mon Sep 29 02:49:19 PM UTC 2025
