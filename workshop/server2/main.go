package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	counterGet    int = 1
	counterCreate int = 1
)

func main() {
	// api
	http.HandleFunc("/create", create)
	http.HandleFunc("/get", get)

	// start server
	fmt.Println("Server two Run")
	http.ListenAndServe(":9992", nil)
	return
}

func create(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	longURL := queryValues.Get("code")

	var shortURL string
	var err error

	fmt.Println("Counter: ", counterCreate)
	counterCreate++

	switch longURL {
	case "wwww.facebook.com":
		shortURL = "FB12"
	case "www.instagram.com":
		shortURL = "IG12"
	case "www.google.com":
		shortURL = "GGLX"
	default:
		err = errors.New("belum ada di list")
	}

	if err != nil {
		w.Write([]byte("link kamu belum bisa untuk di buat pendek, silahkan tambahkan code sendiri :)"))
		return
	}

	w.Write([]byte(shortURL))
	return
}

func get(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	shortURL := queryValues.Get("code")

	var realLink string
	var err error

	fmt.Println("Counter: ", counterGet)
	counterGet++

	switch shortURL {
	case "FB12":
		realLink = "www.facebook.com"
	case "IG12":
		realLink = "www.instagram.com"
	case "GGLX":
		realLink = "www.google.com"
	default:
		err = errors.New("not found sob")
	}

	if err != nil {
		w.Write([]byte("Cannot get the link, you can use /create for create the link"))
		return
	}

	w.Write([]byte(realLink))
	return
}
