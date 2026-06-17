# Turbo Garbanzo - Go HTTP Reverse Proxy

A simple HTTP reverse proxy for local development, built with Go's standard library.

## Overview

This proxy forwards incoming HTTP requests to a configured backend target. Perfect for local API development and testing.

## Requirements

- Go 1.21 or later

## Installation

```bash
git clone https://github.com/haru526/turbo-garbanzo.git
cd turbo-garbanzo
```

## Usage

### Basic Usage

```bash
go run proxy.go -listen=:8080 -target=http://localhost:3000
```

### Command-line Flags

- `-listen` (default: `:8080`) — The address and port the proxy listens on
- `-target` (default: `http://localhost:3000`) — The backend URL where requests are forwarded to

### Examples

#### Forward port 8080 to local backend on port 3000
```bash
go run proxy.go -listen=:8080 -target=http://localhost:3000
```

#### Use a different listen port
```bash
go run proxy.go -listen=:9000 -target=http://localhost:5000
```

#### Forward to a remote backend
```bash
go run proxy.go -listen=:8080 -target=http://api.example.com
```

## Testing the Proxy

Once the proxy is running, test it in another terminal:

```bash
# Simple GET request
curl http://localhost:8080/api/users

# POST request with JSON data
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John", "email": "john@example.com"}'

# With custom headers
curl -H "Authorization: Bearer token123" http://localhost:8080/api/protected
```

## Features

- ✅ Simple and lightweight
- ✅ Minimal dependencies (uses Go standard library only)
- ✅ Request logging with client IP detection
- ✅ Error logging and handling
- ✅ Support for all HTTP methods
- ✅ Configurable listen address and target backend

## Logging

The proxy logs:
- Incoming requests (method, path, client IP)
- Proxy errors and failures
- Server startup information

Example log output:
```
2026/06/17 15:30:45 Starting HTTP reverse proxy
2026/06/17 15:30:45   Listen: :8080
2026/06/17 15:30:45   Target: http://localhost:3000
2026/06/17 15:30:45
2026/06/17 15:30:50 Incoming request: GET /api/users (client: 127.0.0.1)
2026/06/17 15:30:51 Incoming request: POST /api/users (client: 127.0.0.1)
```

## Architecture

### Files

- `proxy.go` — Main proxy implementation using Go's `net/http/httputil.NewSingleHostReverseProxy`
- `go.mod` — Go module definition
- `README.md` — This file

### Implementation Details

- Uses `net/http/httputil.NewSingleHostReverseProxy` for HTTP reverse proxying
- Custom error handler returns HTTP 502 Bad Gateway on proxy errors
- Client IP detection supports `X-Forwarded-For` header
- Configurable timeouts (Read: 15s, Write: 15s, Idle: 60s)

## Future Enhancements

Possible additions (not implemented by default):
- HTTPS/TLS termination with certificate support
- HTTP CONNECT tunneling for proxying
- Request/response middleware
- Rate limiting
- Authentication and authorization

## License

MIT

## Author

haru526
