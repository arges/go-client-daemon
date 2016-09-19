package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

var dataDir = os.Getenv("SNAP_DATA")
var socket = dataDir + "/socket.sock"

func main() {
	// Setup client
	cli := &http.Client{Transport: &http.Transport{Dial: func(_, _ string) (net.Conn, error) {
		return net.Dial("unix", socket)
	},
	}}

	// Format new request
	req, err := http.NewRequest("GET", "http://127.0.0.1/", nil)
	if err != nil {
		log.Fatal("Couldn't format new request: " + err.Error())
	}

	// Issue request
	_, err = cli.Do(req)
	if err != nil {
		log.Fatal("Connection to the daemon failed: " + err.Error())
	}
}

// vim: set tabstop=4:
