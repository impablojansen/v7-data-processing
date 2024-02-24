package reg_manager

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strings"
)

func Territorialize(r1, locations, territories dataframe.DataFrame) dataframe.DataFrame {
	r1 = joinAddress(r1)

	locations = locations.LeftJoin(territories, locations.Names()[0])
	r1 = r1.LeftJoin(locations, territories.Names()[0])

	return r1
}

func joinAddress(df dataframe.DataFrame) dataframe.DataFrame {
	addressColumn := make([]string, df.Nrow())

	for _, col := range df.Names() {
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
