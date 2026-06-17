# turbo-garbanzo

A simple Go HTTP reverse proxy for local development.

## Overview

This is a lightweight reverse proxy that forwards incoming HTTP requests to a configured backend target. It's designed for local development scenarios where you need to proxy requests to another service.

## Features

- **Simple HTTP Reverse Proxy**: Uses Go's standard library for minimal dependencies
- **Command-line Configuration**: Configure listen address and target via flags
- **Request Logging**: Logs all incoming requests with method, path, and client IP
- **Error Handling**: Graceful error handling with appropriate HTTP status codes

## Building

```bash
go build -o proxy proxy.go
```

## Usage

### Basic Usage

```bash
go run proxy.go -listen=:8080 -target=http://localhost:3000
```

### Command-line Flags

- `-listen` (default: `:8080`): Address for the proxy to listen on
- `-target` (default: `http://localhost:3000`): Backend URL to forward requests to

### Examples

**Example 1: Forward requests from port 8080 to port 3000**
```bash
go run proxy.go -listen=:8080 -target=http://localhost:3000
```

**Example 2: Forward requests from port 9000 to a remote server**
```bash
go run proxy.go -listen=:9000 -target=http://api.example.com
```

**Example 3: Forward to HTTPS backend**
```bash
go run proxy.go -listen=:8080 -target=https://api.example.com
```

## Testing

1. Start the target backend (e.g., a Node.js API on port 3000):
   ```bash
   npm start
   ```

2. Start the proxy:
   ```bash
   go run proxy.go -listen=:8080 -target=http://localhost:3000
   ```

3. Make requests through the proxy:
   ```bash
   curl http://localhost:8080/api/users
   curl -X POST http://localhost:8080/api/data -d '{"key":"value"}'
   ```

4. Check the proxy logs to see request details and any errors.

## How It Works

- The proxy listens on the specified address
- All incoming requests are logged with method, path, and client IP
- Requests are forwarded to the target URL, preserving path and query parameters
- Responses from the target are returned to the client
- Any proxy errors are logged and a 502 Bad Gateway response is returned

## Limitations & Future Considerations

This implementation focuses on basic HTTP reverse proxying. It does not currently include:
- HTTPS/TLS certificate management
- CONNECT method for proxying CONNECT requests
- Proxy authentication
- Advanced load balancing or failover

These features can be added if needed for your use case.
