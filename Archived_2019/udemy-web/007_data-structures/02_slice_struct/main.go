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

type info struct {
	Name    string
	Version string
}

func main() {

	classic := info{
		Name:    "Classic Wow",
		Version: "1",
	}

	tbc := info{
		Name:    "The Burning Crusade Wow",
		Version: "2",
	}

	wotlk := info{
		Name:    "Wrath of The Lich King",
		Version: "3",
	}

	cata := info{
		Name:    "Cataclysm",
		Version: "4",
	}

	wows := []info{classic,tbc,wotlk,cata}

	err := tpl.Execute(os.Stdout, wows)

	if err != nil {
		log.Fatalln(err)
	}
}
