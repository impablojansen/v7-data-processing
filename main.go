package main

import (
	"fmt"
	"tratar_base/utils"
)

func main() {
	//fmt.Println("a")
	//file1, _ := os.Open("csv_principal.csv")
	//fmt.Println("a")
	//file2, err := os.Open("csv_suporte.csv")
	//defer file1.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file2.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//df := dataframe.ReadCSV(file1, dataframe.DefaultType(series.String), dataframe.DetectTypes(false))
	//df2 := dataframe.ReadCSV(file2)

	df := utils.FileManager("csv_principal.csv")

	//df2 := utils.FileManager("csv_suporte.csv")

	//joinVariable := df.LeftJoin(df2, "NOM-LOCAL-FAM")

	fmt.Println(df)

	col1 := df.Col("DATA")

	fmt.Println(col1)
	//var dates []string
	//
	//for i := 0; i < col1.Len(); i++ {
	//	var year string = col1.Elem(i).String()[4:8]
	//	var month string = col1.Elem(i).String()[2:4]
	//	var day string = col1.Elem(i).String()[0:2]
	//
	//	s := []string{year, month, day}
	//	dt := strings.Join(s, "-")
	//
	//	dates = append(dates, dt)
	//}
	seriesTest := utils.TransformDate(col1, "Datas")

	fmt.Println(seriesTest)

	//df = df.Capply(dateTransform)
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
