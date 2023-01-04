package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term   string
	Course []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	//tpl = template.Must(template.ParseFiles("src/template/tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Course: []course{
				{Name: "CSCI-40", Number: "4", Units: "Introduction to Programming in Go"},
				{Name: "CSCI-130", Number: "4", Units: "Introduction to Programming with Go"},
				{Name: "CSCI-40", Number: "4", Units: "Mobile Go"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Course: []course{
				{Name: "CSCI-50", Number: "5", Units: "Introduction to Programming in Go55"},
				{Name: "CSCI-190", Number: "5", Units: "Introduction to Programming with Go55"},
				{Name: "CSCI-191", Number: "5", Units: "Mobile Go55"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
