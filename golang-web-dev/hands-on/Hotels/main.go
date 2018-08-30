package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "Hampton",
			Address: "777 e 1way st",
			City:    "Los Angeles",
			Zip:     90059,
			Region:  "Southern",
		},
		hotel{
			Name:    "Hilton",
			Address: "778 e 1way st",
			City:    "Los Angeles",
			Zip:     90059,
			Region:  "Northern",
		},
		hotel{
			Name:    "Westin",
			Address: "779 e 1way st",
			City:    "Los Angeles",
			Zip:     90059,
			Region:  "Central",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
