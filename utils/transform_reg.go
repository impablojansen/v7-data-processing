package utils

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strings"
)

func CorrectReg01(df, loc, ter dataframe.DataFrame) dataframe.DataFrame {
	df = JoinAddressParts(df)
	df = JoinDF(df, loc, "NOM-LOCALIDADE-FAM")
	df = JoinDF(df, ter, "LOCAL-CORRETA")

	return df
}

func JoinDF(df1, df2 dataframe.DataFrame, colName string) dataframe.DataFrame {
	return df1.LeftJoin(df2, colName)
}

func JoinAddressParts(df dataframe.DataFrame) dataframe.DataFrame {
	dfCols := df.Names()
	dfRows := df.Nrow()
	addressColumn := make([]string, dfRows)

	for _, col := range dfCols {
		if (strings.Contains(col, "LOGRADOURO") && !strings.Contains(col, "CEP")) || strings.Contains(col, "COMPLEMENTO") {
			s := df.Col(col)
			strVal := s.Records()
			for i, val := range strVal {
				addressColumn[i] = strings.TrimSpace(addressColumn[i] + " " + val)
			}
		}
	}

	newDf := dataframe.New(series.New(addressColumn, series.String, "ENDERECO"))
	return df.CBind(newDf)
}
