package sender

type Sender interface {
	Send(Instruction) error
	IsConnected() bool
	Disconnect() error
}

type Instruction interface {
	Package() ([]byte, error)
}
