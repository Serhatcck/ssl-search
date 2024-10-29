package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	var (
		typeOf string
		ip     string
		port   string
	)
	flag.StringVar(&typeOf, "typeOf", "server", "Server or client")
	flag.StringVar(&ip, "ip", "localhost", "Server or client connection ip")
	flag.StringVar(&port, "port", "8686", "Server or client connection port")
	flag.Parse()

	if typeOf == "server" {
		server(port)
	} else if typeOf == "client" {
		client(ip, port)
	}
}

func client(serverIP string, serverPort string) {

	// Create TLS configuration
	config := &tls.Config{
		InsecureSkipVerify: true, // Ignore certificate verification for testing
	}

	// Establish TLS connection
	address := fmt.Sprintf("%s:%s", serverIP, serverPort)
	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Print SSL/TLS details
	state := conn.ConnectionState()
	for _, cert := range state.PeerCertificates {
		fmt.Printf("Server Certificate:\n")
		fmt.Printf("\tCommon Name: %s\n", cert.Subject.CommonName)
		fmt.Printf("\tIssuer: %s\n", cert.Issuer)
		fmt.Printf("\tValidity:\n\t\tFrom: %v\n\t\tTo: %v\n", cert.NotBefore, cert.NotAfter)
	}

	// Send message to server
	message := "Hello from the client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Create a buffered reader to read the response in chunks
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1024) // Adjust buffer size as needed

	fmt.Println("Response from server:")
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Printf("Failed to read response: %v", err)
			return
		}
		// Print the received chunk
		fmt.Print(string(buffer[:n]))
	}
	fmt.Println("\nEnd of server response.")
}

func server(port string) {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}

	// Set up TLS configuration
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	// Start TLS listener
	listener, err := tls.Listen("tcp", ":"+port, config)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()
	fmt.Printf("Server listening on port %s over TLS...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())
	tlsConn, ok := conn.(*tls.Conn)
	if ok {
		state := tlsConn.ConnectionState()
		fmt.Printf("Cipher Suite: %s\n", tls.CipherSuiteName(state.CipherSuite))
	} else {
		log.Println("Connection is not a TLS connection")
	}
	// Create a buffered reader to read data in chunks
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1024) // Adjust buffer size as needed

	for {
		// Read into the buffer
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected:", conn.RemoteAddr())
				break
			}
			log.Println("Failed to read data from client:", err)
			return
		}

		// Process the received data
		message := string(buffer[:n])
		fmt.Printf("Message from client: %s\n", message)

		// Send response
		response := "Hello from the server!"
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println("Failed to send response:", err)
			return
		}
	}
}
