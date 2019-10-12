package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

type Person struct {
	Name string
	Age int
	Male bool
}

func main() {

	data := Person{
		Name: "Kienlt",
		Age:  27,
		Male: true,
	}
	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
