# Redstone

Small reverse proxy, designed to forward TCP connections to an internal backend (1:1 map). Operates at the transport layer.

## Environment Variables

- `REDSTONE_PORT` - Port to listen on. Default: `7000`
- `REDSTONE_DESTINATION_ADDR` - Destination address to forward TCP connections to. (Required)
  - Format: `<host>:<port>`
  - Examples:
    - 127.0.0.1:7001
    - 192.168.100.5:80
    - myawesomeapi.com:8080 
    - myawesomeapi.com:80

## Usage Example

1. spin up an HTTP server at localhost:7001 with route - `GET /hello`
2. run 'redstone' after setting $REDSTONE_DESTINATION_ADDR to localhost:7001
   1. `export REDSTONE_DESTINATION_ADDR=localhost:7001`
   2. `go run main.go`
3. `curl localhost:7000/hello`