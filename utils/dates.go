package utils

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strings"
)

func TransformDates(df dataframe.DataFrame) dataframe.DataFrame {
	dfCols := df.Names()
	for _, c := range dfCols {
		if strings.HasPrefix(c, "DAT") || strings.HasPrefix(c, "DTA") {
			s := SeriesToDate(df.Col(c), c)
			df = df.Mutate(s)
		}
	}
	return df
}

func SeriesToDate(rawDates series.Series, seriesName string) series.Series {
	var dates []string

	for i := 0; i < rawDates.Len(); i++ {
		y := rawDates.Elem(i).String()[4:8]
		m := rawDates.Elem(i).String()[2:4]
		d := rawDates.Elem(i).String()[0:2]

		s := []string{y, m, d}
		fmtDate := strings.Join(s, "-")

		dates = append(dates, fmtDate)
	}

	return series.New(dates, series.String, seriesName)
}
