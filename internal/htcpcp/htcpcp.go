package htcpcp

type HTCPCP interface {
	Get() ([]byte, int)
	Brew() ([]byte, int)
	Propfind() ([]byte, int)
	SayWhen() ([]byte, int)
}
