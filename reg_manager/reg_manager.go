package reg_manager

import (
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
	"strings"
	"tratar_base/create_dataframe"
	"tratar_base/utils"
)

func RegManager(path string) {
	reg := create_dataframe.DataframeGenerator(path)
	reg = utils.TransformNumbers(reg)
	reg = utils.TransformDates(reg)
	reg = utils.DropColumns(reg)
	if strings.Contains(path, "reg_01") {
		loc := create_dataframe.DataframeGenerator("./localidades corrigidas.csv")
		ter := create_dataframe.DataframeGenerator("./territorializacao - cras.csv")
		reg = Territorialize(reg, loc, ter)
	}

	writer(reg, "reg_01.csv")
}

func writer(df dataframe.DataFrame, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_ = df.WriteCSV(f)
}
