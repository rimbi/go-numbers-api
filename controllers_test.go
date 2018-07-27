package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

type FakeResponseWriter struct {
	content string
	header  http.Header
}

func (w *FakeResponseWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *FakeResponseWriter) Write(content []byte) (int, error) {
	w.content = string(content)
	return 0, nil
}

func (w *FakeResponseWriter) WriteHeader(statusCode int) {

}

func TestWriteNumbers(t *testing.T) {
	w := new(FakeResponseWriter)
	writeNumbers(w, []int{1, 2, 3})
	if w.content != "{\"numbers\":[1,2,3]}\n" {
		t.Errorf("Write numbers does not write as expected!")
	}
}

func TestNumbersHandlerParsesUrlsProperly(t *testing.T) {
	// given
	w := new(FakeResponseWriter)
	var urls = [2]string{"http://localhost/fibo", "http://localhost/primes"}
	fakeCollectIntegers := func(u string, ch chan []int) {
		for i := 0; i < len(urls); i++ {
			if u == urls[i] {
				urls[i] = ""
			}
		}
		ch <- []int{}
	}
	controller := numbersHandler(fakeCollectIntegers)
	request := new(http.Request)
	request.URL, _ = url.Parse(fmt.Sprintf("http://localhost/?u=%s&u=%s", urls[0], urls[1]))
	// when
	controller(w, request)
	// then
	for _, u := range urls {
		if u != "" {
			t.Errorf("NumbersHandler can not parse urls properly!")
		}
	}
}
