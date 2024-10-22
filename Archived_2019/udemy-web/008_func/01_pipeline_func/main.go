package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.html").Funcs(funcmap).ParseFiles("tpl.html"))
}

//function convert datatime
func monthDateYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var funcmap = template.FuncMap{
	"fdateMDY": monthDateYear,
}

func main() {

	err := tpl.Execute(os.Stdout, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
