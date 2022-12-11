package main

import (
	app "github.com/yosa12978/htcpcp/internal"
	"github.com/yosa12978/htcpcp/internal/htcpcp"
)

func main() {
	app.Run(htcpcp.NewTeapot())
}
