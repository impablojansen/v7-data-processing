package utils

import (
	"encoding/csv"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"io"
	"log"
	"os"
)

func FileManager(fileName string) dataframe.DataFrame {
	s, err := OpenFile(fileName)
	if err != nil {
		return dataframe.New()
	}
	return CSVToDF(s)
}

func OpenFile(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records [][]string
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}
	return records, nil
}

func CSVToDF(s [][]string) dataframe.DataFrame {
	return dataframe.LoadRecords(s, dataframe.DefaultType(series.String), dataframe.DetectTypes(false))
}
