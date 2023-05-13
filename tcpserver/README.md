# TCP Server

A simple TCP server that listens for incoming connections and sends a "Hello World" message to clients.

## Getting Started

1. Clone this repository to your local machine.

2. Navigate to the `tcpserver` directory:

```sh
cd tcpserver 
```

3. Build the server binary:

```sh
go build

```

4. Start the server:

```sh
./tcpserver
```

5. The server will start listening for incoming connections on port `8080`. You can connect to the server using a tool like `telnet`:

```sh
telnet localhost 8080

```

6. Once you connect to the server, you should see a "Hello World" message sent back to you:

```text
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 12

Hello World!
Connection closed by foreign host.
```