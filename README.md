# go_sse

## Description
This is a simple example of how to use Server-Sent Events (SSE) in Go.
Has a simple server that sends a message every second to the client.

### Utilization
- Go
- Docker

Has two endpoints:
- /events - Is the endpoint that the clients will connect to receive the messages.
- / - Is the endpoint that the clients will subscribe to read changes in the messages.
- /webhook - Is the endpoint that the clients will send a message to the server.
- /webhook/:id - Is the endpoint that the clients will send a message to the server with a specific id.

#### JSON Params for /webhook and /webhook/:id
- id as string : used to identify the message by client, no rules for the value
- data as interface{} : used to send the message to the server, no rules for the value
- action as string : used to identify the action of the message(create, update, delete), no rules for the value


## How to run
Create a .env file with the following content:
```dotenv
APP_PORT=8060
```
Run the following commands:
```bash
go run main.go
```

## Docker
```bash
docker build -t go_sse .
docker run -p 8060:8060 go_sse
```

## Docker Image
```bash
docker pull ghcr.io/az-gt/go_az_sse:latest
```

