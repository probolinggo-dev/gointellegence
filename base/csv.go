package base

import (
	"github.com/go-gota/gota/dataframe"
	"strings"
)

func readCsv(csvStr string) dataframe.DataFrame {
	df := dataframe.ReadCSV(strings.NewReader(csvStr))
	return df
}

func exportToCsv()  {
}
