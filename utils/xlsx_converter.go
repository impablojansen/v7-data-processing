package utils

import (
	"errors"
	"github.com/tealeg/xlsx/v3"
	"log"
	"os"
	"strings"
)

func XlsxConverter(fileExt, path string) string {
	oldS, newS := "."+fileExt, ".csv"
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		panic(err)
	}

	newPath := strings.Replace(path, oldS, newS, -1)
	sheetName := wb.Sheets[0].Name
	sheet, ok := wb.Sheet[sheetName]

	if !ok {
		panic(errors.New("sheet not found"))
	}

	file, err := os.OpenFile(newPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	xlsxToCsv(sheet, file, sheet.MaxCol, sheet.MaxRow)

	return newPath
}

func xlsxToCsv(sheet *xlsx.Sheet, file *os.File, cols, rows int) {
	for i := 0; i < rows; i++ {
		var rowString string
		for j := 0; j < cols; j++ {
			cellValue, err := sheet.Cell(i, j)
			if err != nil {
				file.Close()
				log.Fatal(err)
				return
			} else if j < cols-1 {
				str := cellValue.String() + ","
				rowString = rowString + str
			} else {
				str := cellValue.String() + "\n"
				rowString = rowString + str
			}
		}
		_, err := file.Write([]byte(rowString))
		if err != nil {
			file.Close()
			log.Fatal(err)
		}
	}
}
