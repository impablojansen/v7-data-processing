package utils

import (
	"errors"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"os"
)

func FileManager(fileName string) dataframe.DataFrame {
	file, err := OpenFile(fileName)
	if err != nil {
		return dataframe.New()
	}
	return CSVToDF(file)
}

func OpenFile(fileName string) (os.File, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	fmt.Println("openfile:", f)
	if err != nil {
		return nil, errors.New("Teste")
	}
	return f, err
}

func CSVToDF(file *os.File) dataframe.DataFrame {
	return dataframe.ReadCSV(file, dataframe.DefaultType(series.String), dataframe.DetectTypes(false))
}
