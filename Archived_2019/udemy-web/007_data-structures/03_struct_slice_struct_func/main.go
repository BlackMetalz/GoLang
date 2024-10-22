package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var funcx = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

// Get only 3 first string
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func init() {
	tpl = template.Must(template.New("tpl.html").Funcs(funcx).ParseFiles("tpl.html"))
}

type game_info struct {
	Name    string
	Version string
}

type core_info struct {
	Name    string
	Version string
}

type data struct {
	Core       []core_info
	GameClient []game_info
}

func main() {

	classic := game_info{
		Name:    "classic Wow",
		Version: "1",
	}

	tbc := game_info{
		Name:    "the Burning Crusade Wow",
		Version: "2",
	}

	wotlk := game_info{
		Name:    "wrath of The Lich King",
		Version: "3",
	}

	cata := game_info{
		Name:    "Cataclysm",
		Version: "4",
	}

	wow_versions := []game_info{classic, tbc, wotlk, cata}

	mangos_one := core_info{
		Name:    "classic_mangos",
		Version: "1.2.1",
	}

	mangos_two := core_info{
		Name:    "tbc_mangos",
		Version: "2.4.3",
	}

	trinitycore_335a := core_info{
		Name:    "trinity 335a",
		Version: "3.3.5a",
	}

	wow_cores := []core_info{mangos_one, mangos_two, trinitycore_335a}

	data := data{
		Core:       wow_cores,
		GameClient: wow_versions,
	}
	//data := struct {
	//	Core	[]core_info
	//	GameClient []game_info
	//}{
	//	wow_cores,
	//	wow_versions,
	//}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
