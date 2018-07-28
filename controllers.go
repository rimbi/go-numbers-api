package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeNumbers(w http.ResponseWriter, numbers []int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{"numbers": numbers})
}

// numbersHandler handles GET Requests
func numbersHandler(collectNumbers func(string, chan []int)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		urls, ok := r.URL.Query()["u"]

		if !ok || len(urls[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			writeNumbers(w, []int{})
			return
		}
		channels := make([]chan []int, len(urls))
		for i, url := range urls {
			ch := make(chan []int)
			go collectNumbers(url, ch)
			channels[i] = ch
		}
		var set = NewIntSet()
		for _, ch := range channels {
			set.Add(<-ch)
		}
		writeNumbers(w, set.AsSortedSlice())
	}
}
