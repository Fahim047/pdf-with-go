package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)


type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{loremipsumsentence:10}"`
	Price       float64 `fake:"{price:1,10}"`
}

func generateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f);

	frt := []string{}
	frt = append(frt, f.Name)
	frt = append(frt, f.Description)
	frt = append(frt, fmt.Sprintf("%.2f", f.Price))

	return frt
}

func FruitList(lenght int) [][]string {
	var fruits [][] string

	for i := 0; i<lenght; i++{
		onefruit := generateFruit()
		fruits = append(fruits, onefruit)
	}

	return fruits
}