package client

import (
	"fmt"
	"go-daemon-sample/service"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func newFunc(standalone bool) service.Func {
	if standalone {
		return service.NewRawFunc()
	} else {
		f, err := service.NewRPCFunc()
		if err != nil {
			log.Fatalln(err)
		}
		return f
	}
}

func Add(a, b int, standalone bool) {
	f := newFunc(standalone)

	args := &service.Args{
		A: a,
		B: b,
	}
	reply := &service.Reply{}

	if err := f.Add(args, reply); err != nil {
		log.Fatalln("func error:", err)
	}
	fmt.Printf("Add(%d, %d) = %d\n", args.A, args.B, reply.C)
}

func Sub(a, b int, standalone bool) {
	f := newFunc(standalone)

	args := &service.Args{
		A: a,
		B: b,
	}
	reply := &service.Reply{}

	if err := f.Sub(args, reply); err != nil {
		log.Fatalln("func error:", err)
	}
	fmt.Printf("Sub(%d, %d) = %d\n", args.A, args.B, reply.C)
}
