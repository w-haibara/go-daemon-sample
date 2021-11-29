package client

import (
	"bufio"
	"go-deamon-sample/util"
	"io"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Do() {
	conn, err := net.Dial("unix", util.UnixDomainPath)
	if err != nil {
		panic(err)
	}

	msg := []byte("hello\n")
	log.Printf("SEND: %s\n", msg)
	if _, err := conn.Write(msg); err != nil {
		panic(err)
	}

	r := bufio.NewReader(conn)
	line, err := r.ReadBytes(byte('\n'))
	if err != nil && err != io.EOF {
		log.Fatalln("[ERROR]", err)
	}

	log.Printf("RESP: %s\n", line)
}
