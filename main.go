package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/tealeg/xlsx"
)

func hodo(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var carcounter = 0
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("obdVampire")
	hodo(err)
	row = sheet.AddRow()
	for i := 0; i < 1786; i++ {
		url := "http://www.wikiobd.co.uk/vehicle.php?vh=" + strconv.Itoa(i)
		resp, err := soup.Get(url)
		doc := soup.HTMLParse(resp)
		model := doc.Find("ul", "class", "ui-corner-top ui-shadow").Find("li", "class", "ui-li-divider").Find("h3").Text()
		year := doc.Find("ul", "class", "ui-corner-top ui-shadow").Find("li", "class", "ui-li-divider").Find("p").Text()
		desc := doc.Find("div", "id", "description").Find("p").Text()
		photos := doc.Find("div", "class", "content jqm-content").FindAll("img", "class", "locatorimage")
		modelslice := strings.Split(model, " ")

		row = sheet.AddRow()
		//fmt.Println("Brand : " + modelslice[0])
		cell = row.AddCell()
		cell.Value = modelslice[0]
		// fmt.Printf("Model : ")
		// for i := 1; i < len(modelslice); i++ {
		// 	fmt.Printf(modelslice[i] + " ")
		// }
		// fmt.Println("")
		cell = row.AddCell()
		var tmpModelslice = ""
		for i := 1; i < len(modelslice); i++ {
			tmpModelslice = tmpModelslice + " " + modelslice[i]
		}
		cell.Value = tmpModelslice
		// fmt.Println("Year : " + year)
		cell = row.AddCell()
		cell.Value = year
		cell = row.AddCell()
		cell.Value = ""
		cell = row.AddCell()
		cell.Value = ""
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		// fmt.Println("Description : " + desc)
		cell = row.AddCell()
		cell.Value = desc
		// for _, i := range photos {
		// 	fmt.Println("Photo : http://www.wikiobd.co.uk/" + i.Attrs()["src"])
		// }
		cell = row.AddCell()
		var tmpPhotos = ""
		for _, i := range photos {
			tmpPhotos = tmpPhotos + ",http://www.wikiobd.co.uk/" + i.Attrs()["src"]
		}
		cell.Value = tmpPhotos
		cell = row.AddCell()
		cell.Value = ""

		carcounter++
		fmt.Println("Car count: ", carcounter)

		err = file.Save("obdVampire2.xlsx")
		hodo(err)
	}
}
