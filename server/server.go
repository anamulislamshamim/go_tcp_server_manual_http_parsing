package server

import (
	// "bufio"
	"bufio"
	"crud_2/handlers"
	"fmt"
	"io"
	"strconv"

	// "io"
	"net"
	"strings"
)

// StartServer starts a raw tcp server
func StartServer() {
	// Listen on port 3000
	ln, err := net.Listen("tcp", ":3000")

	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on http://localhost:3000")

	for {
		conn, _ := ln.Accept()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Read request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		conn.Write([]byte(handlers.HttpResponse(400, `{"error":"bad request line"}`)))
		return
	}
	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		conn.Write([]byte(handlers.HttpResponse(400, `{"error":"bad request line"}`)))
		return
	}
	method, path := parts[0], parts[1]

	// Read headers
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			conn.Write([]byte(handlers.HttpResponse(400, `{"error":"bad headers"}`)))
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // end of headers
		}
		kv := strings.SplitN(line, ":", 2)
		if len(kv) == 2 {
			headers[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	// Debug: print Content-Length if present
	if cl, ok := headers["Content-Length"]; ok {
		fmt.Printf("Content-Length header: %s\n", cl)
	}

	// Read body ONLY for methods that typically have a body
	body := ""
	if method == "POST" || method == "PUT" || method == "PATCH" {
		if val, ok := headers["Content-Length"]; ok {
			length, err := strconv.Atoi(val)
			if err != nil {
				conn.Write([]byte(handlers.HttpResponse(400, `{"error":"invalid Content-Length"}`)))
				return
			}
			// fmt.Printf("Expected body length: %d bytes\n", length)

			// Read exactly 'length' bytes for the body
			buf := make([]byte, length)
			n, err := io.ReadFull(reader, buf)
			if err != nil {
				fmt.Printf("Read error: %v, bytes read: %d/%d\n", err, n, length)
				conn.Write([]byte(handlers.HttpResponse(400, `{"error":"could not read body"}`)))
				return
			}
			body = string(buf)
			body = strings.TrimSpace(body)
			body = strings.Trim(body, "\r")
			// fmt.Printf("Actual body length: %d bytes\n", len(body))
		}
	}

	// fmt.Printf("Method: %s, Path: %s, Body: %q\n", method, path, body)

	// Dispatch to handler
	resp := handlers.HandleProducts(method, path, body)
	conn.Write([]byte(resp))
}
