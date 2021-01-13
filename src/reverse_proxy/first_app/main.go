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
	fmt.Println("Welcome To Apps ReverseProxy :)")
	http.ListenAndServe(":1111", nil)
	return
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, selamat datang di second app :) ke-", counter)
	counter++

	w.Write([]byte("return from apps"))
	return
}
