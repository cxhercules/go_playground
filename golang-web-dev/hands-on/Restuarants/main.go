package main

import (
	"log"
	"os"
	"text/template"
)

type fooditem struct {
	Name, Description string
	Price             float32
}

type mealtime struct {
	MealTime string
	Menu     []fooditem
}

type mealtimes []mealtime

type restuarant struct {
	Name, Address, City, Zip string
	Services                 mealtimes
}

type restuarants []restuarant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := restuarants{
		restuarant{
			Name:    "Best Meals",
			Address: "111 e 707th st",
			City:    "San Francisco",
			Zip:     "92678",
			Services: mealtimes{
				mealtime{
					MealTime: "Breakfast",
					Menu: []fooditem{
						fooditem{
							Name:        "Pancakes",
							Price:       1.89,
							Description: "Blah blah real good",
						},
						fooditem{
							Name:        "Toast",
							Price:       2.89,
							Description: "Blah blah nice and good",
						},
					},
				},
				mealtime{
					MealTime: "Lunch",
					Menu: []fooditem{
						fooditem{
							Name:        "Chicken and Waffles",
							Price:       5.89,
							Description: "Blah blah hmm good",
						},
						fooditem{
							Name:        "Vegan Crab Cakes",
							Price:       11.89,
							Description: "Blah blah what the",
						},
					},
				},
				mealtime{
					MealTime: "Dinner",
					Menu: []fooditem{
						fooditem{
							Name:        "Lobster",
							Price:       21.89,
							Description: "Blah blah market what",
						},
						fooditem{
							Name:        "Pizza",
							Price:       10.89,
							Description: "Blah blah real deal",
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
