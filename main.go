package main

import (
	"flag"
	"net/http"
)

// Prepares http server and starts its event loop
func runEventLoop(addr string) {
	http.HandleFunc("/numbers", numbersHandler(collectNumbers))
	http.ListenAndServe(addr, nil)
}

func main() {
	listenAddr := flag.String("http.addr", ":8080", "http listen address")
	flag.Parse()

	runEventLoop(*listenAddr)
}
