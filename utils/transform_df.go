package utils

import "github.com/go-gota/gota/dataframe"

func CorrectReg01(r01, l, t dataframe.DataFrame) {

}

func JoinDF(df1, df2 dataframe.DataFrame, colName string) dataframe.DataFrame {
	return df1.LeftJoin(df2, colName)
}
