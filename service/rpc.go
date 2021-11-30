package service

import (
	"go-deamon-sample/util"
	"net/rpc"
)

type RPCFunc struct {
	client *rpc.Client
}

func NewRPCFunc() (Func, error) {
	client, err := rpc.DialHTTP("unix", util.UnixDomainPath)
	if err != nil {
		return nil, err
	}

	rpcfunc := RPCFunc{
		client: client,
	}
	return Func(&rpcfunc), nil
}

func (f *RPCFunc) Add(args *Args, reply *Reply) error {
	return f.client.Call("RawFunc.Add", args, &reply)
}

func (f *RPCFunc) Sub(args *Args, reply *Reply) error {
	return f.client.Call("RawFunc.Sub", args, &reply)
}
