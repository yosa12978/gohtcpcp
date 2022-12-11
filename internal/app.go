package app

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yosa12978/htcpcp/internal/htcpcp"
	"github.com/yosa12978/htcpcp/internal/server"
)

func Run(protocol htcpcp.HTCPCP) {
	port := flag.Int("port", 8080, "Server port.")
	ip := flag.String("ip", "0.0.0.0", "IP address (without port).")
	flag.Parse()

	go server.NewServer(protocol, *port, *ip).ListenAndServe()

	log.Printf("Server is running on port %d", *port)

	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	log.Printf("%s", (<-out).String())
}
