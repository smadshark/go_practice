package main

import (
	"fmt"
	"net/http"
)

type hotDog int

func (m hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
	fmt.Fprintln(w, r.Header["Connection"])
}

type noDap int

func (m noDap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No dap page")
}
func main() {
	var d hotDog
	var no noDap
	http.ListenAndServe(":8080", d)
	http.ListenAndServe(":8080/nodap", no)

	fmt.Println("end line")
}
