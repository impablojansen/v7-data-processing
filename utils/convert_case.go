package utils

import (
	"github.com/go-gota/gota/dataframe"
	"strings"
)

func ConvertCase(df dataframe.DataFrame) dataframe.DataFrame {
	for _, col := range df.Names() {
		df = df.Rename(strings.ToLower(col), col)
	}
	return df
}
