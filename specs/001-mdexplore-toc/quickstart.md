# Quickstart: mdexplore

Display table of contents from markdown files.

## Installation

```bash
go build -o mdexplore ./cmd/mdexplore
go install ./cmd/mdexplore
```

## Usage

```bash
mdexplore README.md --toc
```

## Development

```bash
go test ./...
go test -cover ./...
```

## Troubleshooting

- File not found: Check path exists
- Permission denied: Check file permissions
- Display issues: Ensure UTF-8 and color support
