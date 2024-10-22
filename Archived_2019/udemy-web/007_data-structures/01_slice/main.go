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


func main() {
	savages := []string{"one", "two", "three", "four"}

	err := tpl.Execute(os.Stdout, savages)

	if err != nil {
		log.Fatalln(err)
	}
}


