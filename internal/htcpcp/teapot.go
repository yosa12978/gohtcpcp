package htcpcp

type Teapot struct {
}

func NewTeapot() HTCPCP {
	return &Teapot{}
}

func (tp *Teapot) Get() ([]byte, int) {
	return []byte("I'm a teapot"), 418
}

func (tp *Teapot) Brew() ([]byte, int) {
	return []byte("I'm a teapot"), 418
}

func (tp *Teapot) Propfind() ([]byte, int) {
	return []byte("I'm a teapot"), 418
}

func (tp *Teapot) SayWhen() ([]byte, int) {
	return []byte("I'm a teapot"), 418
}
