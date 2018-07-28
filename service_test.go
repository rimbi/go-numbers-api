package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestServiceCallShouldReturnEmptyNumbersSliceWhenServiceIsNotAvailable(t *testing.T) {
	// given
	ch := make(chan []int)
	// when
	go collectNumbers("http://127.0.0.1:61173", ch)
	numbers := <-ch
	// then
	if len(numbers) != 0 {
		t.Errorf("Unavailable service call did not return empty numbers!")
	}
}
func TestServiceCallShouldReturnEmptyNumbersWhenResponseIsNotOK(t *testing.T) {
	// given
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"numbers":[1, 2]}`)
	}))
	defer ts.Close()
	ch := make(chan []int)
	// when
	go collectNumbers(ts.URL, ch)
	numbers := <-ch
	// then
	if len(numbers) != 0 {
		t.Errorf("Failing service call did not return empty numbers!")
	}
}

func TestServiceCallShouldReturnTheCorrectNumbers(t *testing.T) {
	// given
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"numbers":[1, 2]}`)
	}))
	defer ts.Close()
	ch := make(chan []int)
	// when
	go collectNumbers(ts.URL, ch)
	numbers := <-ch
	// then
	if reflect.DeepEqual(numbers, []int{1, 2, 3}) {
		t.Errorf("Service call did not return the correct numbers!")
	}
}
