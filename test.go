package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// "strconv"
	// "github.com/xuri/excelize/v2"
)

type Response_JSON struct {
	Name string `json:"name"`
}

func searchtxt(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println("body: ", err)

	searchtxt := Response_JSON{
		Name: body.Name,
	}

	fmt.Println("todo: ", searchtxt)
	qqexport(searchtxt.Name)
	return c.SendString(searchtxt.Name)

}

func qqexport(inputName string) {
	txt, err := os.Open(`./` + inputName + `.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = txt.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(txt)

	str := string(b)
	messages := strings.SplitAfter(str, "\n")
	categories := map[string]string{
		"A1": "Name", "B1": "Product", "C1": "Quantity", "D1": "Price", "E1": "Date",
	}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	a := []Record{}
	// values := map[string]string{}
	for _, message := range messages {
		// fmt.Println(len(between(message, "{", "}")))
		if len(between(message, "{", "}")) == 0 {
			fmt.Println("Loading...")
		}
		if len(between(message, "{", "}")) != 0 {
			// fmt.Println("Product:" + between(strings.SplitAfter(between(message, "{", "}"), ";")[0], ":", ";") + ";Quantity:" + between(strings.SplitAfter(between(message, "{", "}"), ";")[1], ":", ";") + ";Price:" + strings.Split(strings.SplitAfter(between(message, "{", "}"), ";")[0], ":")[1])
			validRecord := map[string]string{
				"Name":     strings.Split(strings.Split(message, ":")[2], "]")[1],
				"Product":  between(strings.SplitAfter(between(message, "{", "}"), ";")[0], ":", ";"),
				"Quantity": between(strings.SplitAfter(between(message, "{", "}"), ";")[1], ":", ";"),
				"Price":    strings.Split(strings.SplitAfter(between(message, "{", "}"), ";")[2], ":")[1],
				"Date":     between(message, "[", "]"),
			}
			a = append(a, Record{validRecord["Name"], validRecord["Product"], validRecord["Quantity"], validRecord["Price"], validRecord["Date"]})

		}

	}
	// fmt.Println(a)
	// fmt.Println(len(a))
	values := map[string]string{}
	// count := len(a) - 1
	for x, record := range a {

		values = map[string]string{"A" + strconv.Itoa(x+2): record.Name, "B" + strconv.Itoa(x+2): record.Product, "C" + strconv.Itoa(x+2): record.Quantity, "D" + strconv.Itoa(x+2): record.Price, "E" + strconv.Itoa(x+2): record.Date}
		for k, v := range values {
			f.SetCellValue("Sheet1", k, v)
		}

	}

	if err := f.SaveAs(inputName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

type Record struct {
	Name     string
	Product  string
	Quantity string
	Price    string
	Date     string
}

// func main() {
// 	app := fiber.New()
// 	app.Static("/", ".")
// 	app.Post("/newFileInput", searchtxt)
// 	app.Listen(":3000")
// }
