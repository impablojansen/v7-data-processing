package reg_manager

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strings"
)

func CorrectReg01(df, loc, ter dataframe.DataFrame) dataframe.DataFrame {
	df = DropColumns(df)
	df = JoinAddressParts(df)
	locName := loc.Names()
	terName := ter.Names()
	loc = JoinDF(loc, ter, locName[0])
	df = JoinDF(df, loc, terName[0])

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

func DropColumns(df dataframe.DataFrame) dataframe.DataFrame {
	//df = df.Drop(0)
	dfCols := df.Names()
	for i, col := range dfCols {
		if col == "VAZIO" || col == "\ufeffCHV-NATURAL-PREFEITURA-FAM" {
			df = df.Drop(i)
		}
	}
	return df
}
