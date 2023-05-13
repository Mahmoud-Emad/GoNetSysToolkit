# UDP Echo Server

This is a simple UDP echo server implemented in Go. It listens on UDP port 8080 and echoes back any message it receives.

## Usage

To start the server, run:

```sh
go build
```

This command will build a binary file for the server
You can test the server using a tool like nc (netcat):

```sh
echo "Hello, world!" | nc -u localhost 8080
```

The server should echo back the message:

```text
Hello, world!
```
