package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	for i := 0; i < 1786; i++ {
		url := "http://www.wikiobd.co.uk/vehicle.php?vh=" + strconv.Itoa(i)
		resp, err := soup.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		doc := soup.HTMLParse(resp)

		model := doc.Find("ul", "class", "ui-corner-top ui-shadow").Find("li", "class", "ui-li-divider").Find("h3").Text()
		year := doc.Find("ul", "class", "ui-corner-top ui-shadow").Find("li", "class", "ui-li-divider").Find("p").Text()
		desc := doc.Find("div", "id", "description").Find("p").Text()
		photos := doc.Find("div", "class", "content jqm-content").FindAll("img", "class", "locatorimage")

		modelslice := strings.Split(model, " ")
		fmt.Println("Brand : " + modelslice[0])
		fmt.Printf("Model : ")
		for i := 1; i < len(modelslice); i++ {
			fmt.Printf(modelslice[i] + " ")
		}
		fmt.Println("")
		fmt.Println("Year : " + year)
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		fmt.Println("Description : " + desc)
		for _, i := range photos {
			fmt.Println("Photo : http://www.wikiobd.co.uk/" + i.Attrs()["src"])
		}
		fmt.Println("--------")
	}
}
