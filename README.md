# filecat

A powerful Go command-line tool designed to simplify the extraction and concatenation of relevant information from code repositories and other structured text sources. This consolidated information can then be used as context for Large Language Models (LLMs) or other text processing tasks.

## Motivation

The challenge of managing and providing context to LLMs when dealing with extensive codebases or document collections is a common hurdle. `filecat` addresses this by allowing you to selectively extract, filter, and combine content from multiple files, creating a focused context that enhances the effectiveness of your LLM applications.

## Features

- **Recursive File Concatenation:** Combine files within a directory and its subdirectories.
- **Flexible Filtering:** 
    - Exclude specific directories and file extensions.
    - Include specific directories and file extensions.
    - Ignore or include files without extensions.
- **Customizable Output:**  Output formatted with file paths and code blocks, ready for LLM consumption.
- **Easy-to-use Command Line Interface:** Built with Cobra for a user-friendly experience.


## Installation

```bash
go get github.com/rodneyxr/filecat
go install github.com/rodneyxr/filecat
```

Add the Go bin directory to your system's PATH environment variable to access the `filecat` command from anywhere.

## Usage

```bash
$ filecat run --help
Concatenate files in a directory with filtering

Usage:
  filecat run [directory] [flags]

Flags:
  -d, --exclude-dirs strings    Directories to exclude
  -e, --exclude-exts strings    File extensions to exclude (without the dot)
  -h, --help                    help for run
  -i, --ignore-extensionless    Ignore files without extensions
  -D, --include-dirs strings    Directories to include
  -I, --include-extensionless   Include files without extensions
  -E, --include-exts strings    File extensions to include (without the dot)
```

- `directory`: The path to the directory (e.g., your code repository) containing the files you want to process. Defaults to the current directory if not provided.

### Flags
| Flag | Shorthand | Description
| ---- | --------- | -----------
| `--exclude-dirs`     | `-d`     | A comma-separated list or multiple flags specifying directories to exclude.
| `--exclude-exts`     | `-e`     | A comma-separated list or multiple flags specifying file extensions to exclude (without the leading dot).
| `--include-dirs`     | `-D`     | A comma-separated list or multiple flags specifying directories to include. If empty, all directories are included.
| `--include-exts`     | `-E`     | A comma-separated list or multiple flags specifying file extensions to include (without the leading dot). If empty, all extensions are included.
| `--ignore-extensionless` | `-i`     | Ignores files without any file extension.
| `--include-extensionless` | `-I`     | Includes files without any file extension. This overrides the `--ignore-extensionless` flag if both are used.

### Examples

```bash
# Concatenate all files in the current directory, excluding "node_modules" and ".txt" files
filecat run -d "node_modules" -e txt

# Concatenate only Go and Markdown files in a specific directory
filecat run /path/to/project -E o,md

# Include only files in "src" and "docs" directories
filecat run -D src,docs

# Ignore files without extensions
filecat run -i

# Include files without extensions
filecat run -I

# Extract documentation and configuration from a code repository, excluding tests
filecat run /path/to/code/repo -d "tests" -E "md,yml,yaml" -i 

# Combine relevant Python scripts for LLM analysis
filecat run /path/to/scripts -E "py"
```

## Output Format

The concatenated output is written to standard output (your terminal) in the following markdown-friendly format:

~~~markdown
## FILE: "path/to/hello.py"
```py
print("Hello, World!")
```

## FILE: "path/to/README.md"
```md
# Markdown Content
This is some markdown content.
```
~~~

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
