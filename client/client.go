package client

import (
	"fmt"
	"go-deamon-sample/deamon"
	"go-deamon-sample/util"
	"log"
	"net/rpc"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Do() {
	client, err := rpc.DialHTTP("unix", util.UnixDomainPath)
	if err != nil {
		log.Fatalln("dialing:", err)
	}

	args := &deamon.Args{
		A: 7,
		B: 8,
	}
	reply := &deamon.Reply{}

	if err = client.Call("Func.Add", args, &reply); err != nil {
		log.Fatalln("func error:", err)
	}
	fmt.Printf("Func.Add(%d, %d) = %d\n", args.A, args.B, reply.C)

	if err = client.Call("Func.Sub", args, &reply); err != nil {
		log.Fatalln("func error:", err)
	}
	fmt.Printf("Func.Sub(%d, %d) = %d\n", args.A, args.B, reply.C)
}
