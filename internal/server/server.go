package server

import (
	"net/http"
	"strconv"

	"github.com/yosa12978/htcpcp/internal/htcpcp"
)

func NewServer(protocol htcpcp.HTCPCP, port int, addr string) *http.Server {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    addr + ":" + strconv.Itoa(port),
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var status int
		var message []byte
		if r.Method == "GET" {
			message, status = protocol.Get()
		} else if r.Method == "BREW" || r.Method == "POST" {
			message, status = protocol.Brew()
		} else if r.Method == "PROPFIND" {
			message, status = protocol.Propfind()
		} else if r.Method == "WHEN" {
			message, status = protocol.SayWhen()
		} else {
			status = 405
			message = []byte("Method not allowed")
		}
		w.WriteHeader(status)
		w.Write(message)
	})

	return &server
}
