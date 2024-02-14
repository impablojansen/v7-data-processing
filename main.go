package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"log"
	"os"
	"strings"
)

func main() {
	file1, err := os.Open("csv_principal.csv")
	file2, err := os.Open("csv_suporte.csv")
	defer file1.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file1, dataframe.DefaultType(series.String), dataframe.DetectTypes(false))
	//df2 := dataframe.ReadCSV(file2)

	//joinVariable := df.LeftJoin(df2, "LOCALIDADE")

	dateTransform := func(s series.Series) series.Series {
		strings := s.String()
		//fmt.Println(s.Elem(0))
		//fmt.Println(s.Elem(1))
		return series.Strings(strings)
	}

	col1 := df.Col("DATA")

	var dates []string

	for i := 0; i < col1.Len(); i++ {
		var year string = col1.Elem(i).String()[4:8]
		var month string = col1.Elem(i).String()[2:4]
		var day string = col1.Elem(i).String()[0:2]

		s := []string{year, month, day}
		dt := strings.Join(s, "-")

		dates = append(dates, dt)
	}
	seriesTest := series.New(dates, series.String, "DATA")

	fmt.Println(seriesTest)

	df = df.Capply(dateTransform)
	//fmt.Println(joinVariable)

	//var s []string

	//const colName string = "col1"
	//const colIdx int = 0
	//for i := 0; i < df.Nrow(); i++ {
	//	val := df.Elem(i, colIdx)
	//	//fmt.Println(val)
	//	filter1 := df2.Filter(
	//		dataframe.F{
	//			Colname:    colName,
	//			Comparator: series.Eq,
	//			Comparando: val,
	//		})
	//	anotherVal := filter1.Elem(0, 1)
	//	//s = append(s, anotherVal)
	//	fmt.Println(anotherVal)
	//}
}
