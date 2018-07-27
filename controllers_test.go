package main

import (
	"net/http"
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

func TestNumbersHandlerCallsServicesPassedAsParameter(t *testing.T) {
	w := new(FakeResponseWriter)
	writeNumbers(w, []int{1, 2, 3})
	if w.content != "{\"numbers\":[1,2,3]}\n" {
		t.Errorf("Write numbers does not write as expected!")
	}
}
