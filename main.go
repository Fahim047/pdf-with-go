package main

import (
	"fmt"
	"os"
	"pdf-with-Go/data"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildFruitList(m)
	
	err := m.OutputFileAndClose("pdf/go.pdf")
	if err != nil {
		fmt.Println("Oppsssssss!!!! Something wrong, couldn't save PDF: ", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully!!!")

}

func buildFruitList(m pdf.Maroto) {
	tableHeadings := []string{"Fruit", "Description", "Price"}
	contents := data.FruitList(20)
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Product List and Price", props.Text{
				Top: 2,
				Size: 13,
				Color: color.NewWhite(),
				Family: consts.Courier,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size: 9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size: 9,
			GridSizes: []uint{3, 7, 2},
		},
		Align: consts.Left,
		HeaderContentSpace: 1,
		Line: false,
		AlternatedBackground: &lightPurpleColor,
	})
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/fruits.jpg", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("image could not be loaded!!!!!")
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("We are learning how to make pdf using golang", props.Text{
				Top: 3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}