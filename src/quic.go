package main

import (
	"flag"
	"isl/samplecodes/go-quic/src/quicclient"
	"isl/samplecodes/go-quic/src/quicserver"
)

func main() {
	var server = flag.Bool("server", false, "")
	var ip = flag.String("ip", "127.0.0.1", "")
	var port = flag.Int("port", 4242, "")

	flag.Parse()

	if *server {
		quicserver.Serve(*ip, *port)
	} else {
		quicclient.Serve(*ip, *port)
	}
}
