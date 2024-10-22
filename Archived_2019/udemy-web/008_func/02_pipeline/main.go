package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.html").Funcs(funcmap).ParseFiles("tpl.html"))
}

var funcmap = template.FuncMap{
	"double":  double,
	"triple":  triple,
	"divive2": divive2,
}

func double(x int) int {
	return x + x
}

func triple(x int) int {
	return x * 3
}

func divive2(x int) int {
	return x / 2
}

func main() {

	err := tpl.Execute(os.Stdout, 2)

	if err != nil {
		log.Fatalln(err)
	}
}
