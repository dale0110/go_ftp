package main

import (
  "io"
	"log"
	"net"
	"time"
	"fmt"
)

func main() {
/*
	cert, err := tls.LoadX509KeyPair("rui.crt", "rui.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader
*/
	service := "127.0.0.1:8000"

	listener, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}

	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Printf("server: conn: read: %s %d", err,n)
			}
			break
		}
		
		t := time.Now
		fmt.Println(t)
		//fmt.Println("YEAR:%d month:%d day:%d \n",t.Year(),t.Mon(),t.Day())
		//buf += string(rand.Reader)
		log.Printf("server: conn: echo %q time:%d\n", string(buf[:n]),t.String())
		n, err = conn.Write(buf[:n])
		log.Printf("server: conn: wrote %d bytes", n)

		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")
}
