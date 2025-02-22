# shout-server

This is a simple UDP echo server written in Go using low-level system calls via the [syscall](https://pkg.go.dev/syscall) package. 

It listens on port `7777`, receives messages, converts them to uppercase, and sends them back to the sender.

## Test

Start the server:

```shell
go run .
```

Send a message via UDP to server with netcat:

```shell
nc -u 127.0.0.1 7777
hey
HEY
hello
HELLO
```