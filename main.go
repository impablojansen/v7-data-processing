package main

import (
	"fmt"
	"tratar_base/create_dataframe"
)

func main() {
	test := create_dataframe.DetectType("reg_01.json")
	fmt.Println(test)

	//df := utils.FileManager("reg_01.csv")
	//df2 := utils.FileManager("localidades corrigidas.csv")
	//df3 := utils.FileManager("territorializacao - cras.csv")
	//
	//df = utils.TransformDates(df)
	//
	//df = utils.TransformNumbers(df)
	//df = reg_manager.CorrectReg01(df, df2, df3)

	//df = reg_manager.JoinAddressParts(df)
}
