package main

import (
	"fmt"
	"time"
	"tratar_base/reg_manager"
	"tratar_base/utils"
)

func main() {
	start := time.Now()
	df := utils.FileManager("reg_01.csv")
	df2 := utils.FileManager("localidades corrigidas.csv")
	df3 := utils.FileManager("territorializacao - cras.csv")

	df = utils.TransformDates(df)

	df = utils.TransformNumbers(df)
	df = reg_manager.CorrectReg01(df, df2, df3)

	t := time.Now()

	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	//df = reg_manager.JoinAddressParts(df)
}
