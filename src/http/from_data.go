package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type hotDog1 int

func (m hotDog1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(os.Getwd())
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "src/http/index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("src/http/index.gohtml"))
}

func main() {
	var d hotDog
	http.ListenAndServe(":8080", d)
}
