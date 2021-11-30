package service

type RawFunc struct {
}

func NewRawFunc() Func {
	rawfunc := RawFunc{}
	return Func(&rawfunc)
}

func (f *RawFunc) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (f *RawFunc) Sub(args *Args, reply *Reply) error {
	reply.C = args.A - args.B
	return nil
}
