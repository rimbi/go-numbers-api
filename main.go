package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

// Gets list of ints from a service located at <url>
// and returns them via channel <ch>
func collectIntegers(url string, ch chan []int) {
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Could not get response from the service: ", url)
		return
	}
	if resp.StatusCode != 200 {
		ch <- []int{}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Could not read response body from ", url)
	}
	var result map[string][]int
	json.Unmarshal(body, &result)
	ch <- result["numbers"]
}

// Prepares http server and starts its event loop
func runEventLoop(addr string) {
	http.HandleFunc("/numbers", numbersHandler(collectIntegers))
	http.ListenAndServe(addr, nil)
}

func main() {
	listenAddr := flag.String("http.addr", ":8080", "http listen address")
	flag.Parse()

	runEventLoop(*listenAddr)
}
