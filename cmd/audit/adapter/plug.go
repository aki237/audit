package adapter

type Plug uint

const (
	OS Plug = iota + 1
	NETWORK
	LOG
)
