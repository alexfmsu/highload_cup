package Limit

import (
	"highload2018/db/filter/Structs"
	"strconv"
)

func HandleLimit(filter structs.Filter) (string, bool) {
	LIMIT := ""

	if filter.Limit != "" {
		val, _ := strconv.Atoi(filter.Limit)

		// println("\n", filter.Limit, val, "\n")

		if val > 0 {
			// println("\n", "aaa", filter.Limit, val, "\n")
			LIMIT = "limit " + filter.Limit
		} else {
			return "", false
		}
	}

	return LIMIT, true
}
