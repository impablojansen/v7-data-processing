package utils

import (
	"github.com/go-gota/gota/series"
	"strings"
)

func TransformDate(rawDates series.Series, seriesName string) series.Series {
	var dates []string

	for i := 0; i < rawDates.Len(); i++ {
		var year string = rawDates.Elem(i).String()[4:8]
		var month string = rawDates.Elem(i).String()[2:4]
		var day string = rawDates.Elem(i).String()[0:2]

		s := []string{year, month, day}
		formattedDate := strings.Join(s, "-")

		dates = append(dates, formattedDate)
	}

	return series.New(dates, series.String, seriesName)
}
