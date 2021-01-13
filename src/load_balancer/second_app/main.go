package main

import (
	"fmt"
	"net/http"
)

var counter int

func main() {
	// api
	http.HandleFunc("/", hiHandler)

	// start server
	fmt.Println("welcome to LoadBalancer #2")
	http.ListenAndServe(":3333", nil)
	return
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, selamat datang di first app :) ke-", counter)
	counter++

	w.Write([]byte("return from apps"))
	return
}
