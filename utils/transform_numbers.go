package utils

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"regexp"
	"strconv"
	"strings"
)

func TransformNumbers(df dataframe.DataFrame) dataframe.DataFrame {
	dfCols := df.Names()
	for _, col := range dfCols {
		if strings.HasPrefix(col, "VLR-") || strings.HasPrefix(col, "VAL-") || strings.HasPrefix(col, "QTD-") || strings.HasPrefix(col, "NUM-LOG") {
			s := df.Col(col)

			if strings.HasPrefix(col, "VLR-RENDA") {
				strToInt, _ := s.Int()
				newDf := transformMoney(strToInt)
				df = df.Mutate(series.New(newDf, series.Int, col))
			} else if strings.HasPrefix(col, "NUM-LOG") {
				strToInt := s.Records()
				addressList := transformAddressNumber(strToInt)
				df = df.Mutate(series.New(addressList, series.String, col))
			} else {
				strToInt, _ := s.Int()
				df = df.Mutate(series.New(strToInt, series.Int, col))
			}
		} else if strings.HasPrefix(col, "NUM-TEL-CONTATO-1-FAM") {
			df = transformPhoneNumber(df)
		}
	}
	return df
}

func transformMoney(n []int) []int {
	for i, val := range n {
		n[i] = val / 100
	}
	return n
}

func transformAddressNumber(n []string) []string {
	var addressList []string
	for i, val := range n {
		if n[i] != "" {
			num, _ := strconv.Atoi(val)
			num = num / 1
			addressList = append(addressList, strconv.Itoa(num))
		} else {
			addressList = append(addressList, "")
		}
	}
	return addressList
}

func transformPhoneNumber(df dataframe.DataFrame) dataframe.DataFrame {
	telString, tel1, tel2 := make([]string, df.Nrow()), make([]string, df.Nrow()), make([]string, df.Nrow())
	colPrefix := "NUM-TEL-CONTATO-"

	re := regexp.MustCompile(`[^0-9]+`)

	for _, col := range df.Names() {
		if strings.Contains(col, colPrefix) {
			s := df.Col(col)
			strVal := s.Records()
			for i, val := range strVal {
				val = re.ReplaceAllString(val, "")
				telString[i] = strings.TrimSpace(telString[i] + val)
			}
		}
	}

	for i, tel := range telString {
		if len(tel) >= 12 && tel[0:12] != "000000000000" {
			tel1[i] = tel[0:2] + tel[3:12]
		}
		if len(tel) >= 24 && tel[12:24] != "000000000000" {
			tel2[i] = tel[12:14] + tel[15:24]
		}

	}

	dfTel1 := dataframe.New(series.New(tel1, series.String, "TEL-1"))
	dfTel2 := dataframe.New(series.New(tel2, series.String, "TEL-2"))
	df = df.CBind(dfTel1)
	df = df.CBind(dfTel2)

	return df
}
