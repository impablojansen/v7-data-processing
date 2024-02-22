package utils

import (
	"errors"
	"github.com/tealeg/xlsx/v3"
	"log"
	"os"
)

func XLSX_converter() {
	wb, err := xlsx.OpenFile("teste.xlsx")
	if err != nil {
		panic(err)
	}

	sheetName := wb.Sheets[0]

	sh, ok := wb.Sheet[sheetName.Name]
	if !ok {
		panic(errors.New("sheet not found"))
	}

	file, err := os.OpenFile("new_test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	colQt, rowQt := sh.MaxCol, sh.MaxRow

	for i := 0; i < rowQt; i++ {
		var rowString string
		for j := 0; j < colQt; j++ {
			c, _ := sh.Cell(i, j)
			if j < colQt-1 {
				str := c.String() + ","
				rowString = rowString + str
			} else {
				str := c.String() + "\n"
				rowString = rowString + str
			}
		}
		_, err = file.Write([]byte(rowString))
		if err != nil {
			file.Close()
			log.Fatal(err)
		}
	}
}
