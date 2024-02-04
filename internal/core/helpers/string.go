package helpers

import (
	"strconv"
	"strings"
)

func EscapeColumns(columns []string) []string {
	for i := range columns {
		if strings.HasPrefix(columns[i], `"`) {
			continue
		}

		columns[i] = strconv.Quote(columns[i])
	}

	return columns
}
