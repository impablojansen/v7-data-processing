package utils

import (
	"github.com/go-gota/gota/dataframe"
	"strings"
)

var (
	dropColsNames = [...]string{"VAZIO", "CHV-NATURAL-PREFEITURA-FAM"}
)

func DropColumns(df dataframe.DataFrame) dataframe.DataFrame {
	dfCols := df.Names()

	for _, dropName := range dropColsNames {
		for i, colName := range dfCols {
			if strings.Contains(colName, dropName) {
				df = df.Drop(i)
			}
		}
	}
	return df
}
