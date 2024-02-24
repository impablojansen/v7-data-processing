package create_dataframe

import (
	"errors"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"log"
	"os"
	"strings"
	"tratar_base/utils"
)

var (
	allowedExtensions = [...]string{"csv", "xlsx"}
)

func DataframeGenerator(path string) dataframe.DataFrame {
	fileExt, err := DetectType(path)
	if err != nil {
		log.Fatal(err)
	}

	if fileExt != "csv" {
		path = utils.XlsxConverter(fileExt, path)
	}

	return transformDf(path)
}

func DetectType(path string) (string, error) {
	names := strings.Split(path, ".")
	fileExt := names[len(names)-1]

	for _, val := range allowedExtensions {
		if fileExt == val {
			return fileExt, nil
		}
	}
	return "", errors.New("extension not allowed")
}

func transformDf(path string) dataframe.DataFrame {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(f, dataframe.DefaultType(series.String), dataframe.DetectTypes(false))
	return df
}
