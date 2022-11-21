package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("I'm sorry but endpoint %s is not handled", r.RequestURI)))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8002", nil)
}
