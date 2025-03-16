
# QR Code Generator (qrgen)

`qrgen` is a versatile QR code generator built in Go. It offers two modes of operation:

1. **CLI Mode**: Generate QR codes directly in your terminal as ASCII art.
2. **Web Mode**: Run a web server to create and view QR codes through a browser with a user-friendly interface, including dark mode support.

This tool uses the `github.com/skip2/go-qrcode` library for QR code generation and `github.com/spf13/cobra` for the CLI interface.

## Usage

### server

```bash
go run cmd/server/main.go
```

- Go to <http://localhost:8888/>
- Add URL and size, click on "GENERATE QR Code"
- You will see a generated QR code.
- Download QR code by clicking on "View Image"

### cli

```bash
go run cmd/cli/main.go  generate --url="https://viveksahu26.substack.com/"  --size=256
```
