package main

import (
	"fmt"
	"tratar_base/utils"
)

func main() {
	df := utils.FileManager("csv_principal.csv")
	df2 := utils.FileManager("csv_suporte.csv")

	joinVariable := df.LeftJoin(df2, "NOM-LOCAL-FAM")

	//col1 := joinVariable.Col("DATA")

	//seriesTest := utils.SeriesToDate(col1, "DATA")
	//joinVariable = joinVariable.Mutate(seriesTest)

	//t := utils.TransformDates(joinVariable)
	//fmt.Println(t)
	t := utils.TransformNumbers(joinVariable)
	fmt.Println(t)
}
