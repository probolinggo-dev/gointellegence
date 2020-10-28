package base

import (
	"github.com/go-gota/gota/dataframe"
	"strings"
)

func readJSON(jsonStr string) dataframe.DataFrame {
	df := dataframe.ReadJSON(strings.NewReader(jsonStr))
	return df
}

func exportToJSON() {

}
