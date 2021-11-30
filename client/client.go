package client

import (
	"fmt"
	"go-deamon-sample/service"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Do() {
	args := &service.Args{
		A: 7,
		B: 8,
	}
	reply := &service.Reply{}

	//f := service.NewRawFunc()
	f, err := service.NewRPCFunc()
	if err != nil {
		log.Fatalln(err)
	}

	if err := f.Add(args, reply); err != nil {
		log.Fatalln("func error:", err)
	}
	fmt.Printf("Add(%d, %d) = %d\n", args.A, args.B, reply.C)

	if err := f.Sub(args, reply); err != nil {
		log.Fatalln("func error:", err)
	}
	fmt.Printf("Sub(%d, %d) = %d\n", args.A, args.B, reply.C)
}
