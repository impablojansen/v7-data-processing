package main

import (
	"fmt"
	"tratar_base/utils"
)

func main() {
	df := utils.FileManager("csv_principal.csv")
	df2 := utils.FileManager("csv_suporte.csv")

	df = df.LeftJoin(df2, "NOM-LOCAL-FAM")
	df = utils.TransformDates(df)

	t := utils.TransformNumbers(df)
	withAddress := utils.JoinAddressParts(t)
	fmt.Println(withAddress)
}
