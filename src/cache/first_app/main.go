package main

import (
	"fmt"
	"net/http"
)

var counter int

func main() {
	// api
	http.HandleFunc("/", hiHandler)
	http.HandleFunc("/dynamic", hitHandlerDynamic)

	// start server
	fmt.Println("welcome to Cache Apps")
	http.ListenAndServe(":4444", nil)
	return
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, selamat datang di first app :) ke-", counter)
	counter++

	w.Write([]byte("return from apps"))
	return
}

func hitHandlerDynamic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, selamat datang di dynamic app :) ke-", counter)
	counter++

	var value []string
	for _, v := range r.URL.Query() {
		value = v
	}

	w.Write([]byte(fmt.Sprintf("return from apps dynamic-%s", value)))
	return
}
