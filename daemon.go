package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	dataDir = os.Getenv("SNAP_DATA")
	socket  = dataDir + "/socket.sock"
)

func main() {

	// Check if socket already exists. In cases of a crash the daemon may not
	// have been cleanly shutdown.
	if _, err := os.Stat(socket); err == nil {
		fmt.Println("Socket already exists; daemon did not exit cleanly! Removing old socket.")
		os.Remove(socket)
	}

	// Setup main listener socket.
	listener, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Fixup permissions for anyone to access
	err = os.Chmod(socket, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Serve up endpoint
	mux := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.Serve(listener, mux)
}

// vim: set tabstop=4:
