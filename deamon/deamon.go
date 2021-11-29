package deamon

import (
	"bufio"
	"go-deamon-sample/util"
	"io"
	"log"
	"net"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Do() {
	os.Remove(util.UnixDomainPath)
	listener, err := net.Listen("unix", util.UnixDomainPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()

			r := bufio.NewReader(conn)
			line, err := r.ReadBytes(byte('\n'))
			if err != nil && err != io.EOF {
				log.Println("[ERROR]", err)
				return
			}
			log.Printf("REQ:%s\n", line)

			if _, err := conn.Write(line); err != nil {
				log.Println("[ERROR]", err)
				return
			}
			log.Printf("SEND:%s\n", line)
		}()
	}
}
