package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// collectNumbers gets list of numbers from a service located at <url>
// and returns them via channel <ch>. In case error it returns empty list
func collectNumbers(url string, ch chan []int) {
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- []int{}
		log.Println("Could not get response from the service: ", url)
		return
	}
	if resp.StatusCode != 200 {
		ch <- []int{}
		log.Println("Service call failed with status code ", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []int{}
		log.Println("Could not read response body from ", url)
		return
	}
	var result map[string][]int
	json.Unmarshal(body, &result)
	ch <- result["numbers"]
}
