package create_dataframe

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/northbright/xls2csv-go/xls2csv"
	"log"
	"strings"
)

type df struct {
	data     *dataframe.DataFrame
	fileType string
	options  *dataframe.LoadOption
}

var (
	dfInstance *df
)

//func NewDF(data) (*df, error) {
//	strings.Contains()
//	df, err :=
//	dfInstance = &df{df}
//	df{
//		*
//	}
//}}

func DetectType(fileName string) string {
	_, fileType, _ := strings.Cut(fileName, ".")
	return fileType
}

func xlsxToCsv(filePath string) [][]string {
	records, err := xls2csv.XLS2CSV(filePath, 0)
	if err != nil {
		log.Fatal(err)
	}
	return records
}
