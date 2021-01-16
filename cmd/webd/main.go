package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:5000"

	log.Printf("Start server on address %s", addr)

	lisner, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer lisner.Close()

	for {
		log.Print("waiting for a connection")
		conn, err := lisner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Print("accepted connection")
		err = handelConnection(conn)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handelConnection(conn net.Conn) error {
	defer conn.Close()
	defer log.Print("done handling connection")

	log.Print("writing to connecction")
	_, err := conn.Write([]byte("hello world"))
	if err != nil {
		return fmt.Errorf("couldn't write to connection %v", err)
	}

	log.Print("reading form connection..")
	resp := make([]byte, 32)

	n, err := conn.Read(resp)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return fmt.Errorf("couldn't read to connection %v", err)
	}
	resp = resp[:n]

	log.Printf("recived message: %q", string(resp))

	return nil
}
