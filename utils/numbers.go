package utils

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strconv"
	"strings"
)

func TransformNumbers(df dataframe.DataFrame) dataframe.DataFrame {
	dfCols := df.Names()
	for _, col := range dfCols {
		if strings.HasPrefix(col, "VLR-") || strings.HasPrefix(col, "VAL-") || strings.HasPrefix(col, "QTD-") || strings.HasPrefix(col, "NUM-LOG") {
			s := df.Col(col)
			strToInt, _ := s.Int()

			if strings.HasPrefix(col, "VLR-RENDA") {
				strToInt = TransformMoney(strToInt)
				df = df.Mutate(series.New(strToInt, series.Int, col))
			} else if strings.HasPrefix(col, "NUM-LOG") {
				addressList := TransformAddressNumber(strToInt)
				df = df.Mutate(series.New(addressList, series.String, col))
			} else {
				df = df.Mutate(series.New(strToInt, series.Int, col))
			}
		}
	}
	return df
}

func TransformMoney(n []int) []int {
	for i, val := range n {
		n[i] = val / 100
	}
	return n
}

func TransformAddressNumber(n []int) []string {
	var addressList []string
	for i, val := range n {
		if n[i] > 0 {
			num := val / 1
			addressList = append(addressList, strconv.Itoa(num))
		} else {
			addressList = append(addressList, "")
		}
	}
	return addressList
}
