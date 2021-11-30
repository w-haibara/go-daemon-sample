package service

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Func interface {
	Add(*Args, *Reply) error
	Sub(*Args, *Reply) error
}
