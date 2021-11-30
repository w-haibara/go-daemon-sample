package deamon

import (
	"go-deamon-sample/util"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Func struct {
}

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func (f *Func) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (f *Func) Sub(args *Args, reply *Reply) error {
	reply.C = args.A - args.B
	return nil
}

func Do() {
	f := new(Func)
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
