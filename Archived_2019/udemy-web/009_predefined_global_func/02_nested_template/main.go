package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html",33)
	if err != nil {
		log.Fatalln(err)
	}
}
