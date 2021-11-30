package daemon

import (
	"go-daemon-sample/service"
	"go-daemon-sample/util"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Do() {
	f := new(service.RawFunc)
	if err := rpc.Register(f); err != nil {
		log.Fatalln(err)
	}

	os.Remove(util.UnixDomainPath)
	l, err := net.Listen("unix", util.UnixDomainPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	rpc.HandleHTTP()
	if err := http.Serve(l, nil); err != nil {
		log.Fatalln(err)
	}
}
