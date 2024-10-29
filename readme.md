# TLS Client-Server Application

This is a simple TLS client-server application written in Go. The application allows for secure communication between a client and a server using TLS (Transport Layer Security). The server listens for incoming connections, while the client connects to the server and exchanges messages.

## Features

- **TLS Communication**: Establishes a secure connection using TLS.
- **Client and Server Modes**: Can run in either client or server mode.
- **Cipher Suite Information**: Displays the cipher suite used for the TLS connection.
- **Certificate Details**: Prints details of the server's certificate upon connection.

## Requirements

- Go (1.15 or higher)
- TLS certificate files (`server.crt` and `server.key`)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Serhatcck/ssl-search
cd ssl-search
```

2.	Ensure you have the required TLS certificate and key files (server.crt and server.key) in the same directory as the source code.


# Usage

## Running the Server

To start the server, use the following command: 

```bash
go run main.go -typeOf server -port <port>      
```
   

•	Replace <port> with the desired port number (default is 8686).

## Running the Client
To start the client, use the following command:

```bash
go run main.go -typeOf client -ip <server-ip> -port 
```  

•	Replace <server-ip> with the server’s IP address (default is localhost).
•	Replace <port> with the port number on which the server is listening (default is 8686).

## Walkthrough 
https://medium.com/@serhatcck/how-do-ssl-scanners-work-b4977308e981


