# Source Viewer (sv)

A command-line tool that captures code files as screenshots and combines them into a single PDF document. The tool uses the Carbonara API to generate beautiful syntax-highlighted screenshots of your code.

## Installation

### Using Go

```bash
go install github.com/exgene/syntax-view@latest
```

### Build from source

```bash
git clone https://github.com/exgene/syntax-view.git
cd syntax-view
go build
```

## Usage

### Basic Usage

```bash
sv capture -d /path/to/directory -o output.pdf
```

```bash
sv capture [flags]
Flags:
-d, --dir string Directory to capture (required)
-m, --markdown Produces output in a markdown file
-o, --output string Output file path (required)
-h, --help Help for capture command
```

## Supported File Extensions (Which files are parsed)

The tool automatically processes files with the following extensions:

- `.go` (Go)
- `.js` (JavaScript)
- `.py` (Python)
- `.java` (Java)
- `.cpp` (C++)
- `.c` (C)
- `.h` (Header files)
- `.rs` (Rust)
- `.rb` (Ruby)
- `.php` (PHP)
- `.ts` (TypeScript)

It will only read these files for now

TODO: Add flags to support files and ignore files
