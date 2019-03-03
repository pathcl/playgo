// create a struct and then uses a pointer
// create a method
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

// create a method for a given func

func (w *webPage) get() {
	resp, _ := http.Get(w.url)

	// copy not assign
	w.body, _ = ioutil.ReadAll(resp.Body)
}

func main() {
	w := &webPage{url: "http://httpbin.org/ip"}
	w.get()
	fmt.Printf("%s", w.body)
}
