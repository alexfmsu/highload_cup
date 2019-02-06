package Limit

import (
	"highload2018/structs"
	"strconv"
)

func HandleLimit(recommend structs.Recommend) (string, bool) {
	LIMIT := ""

	if recommend.Limit != "" {
		val, _ := strconv.Atoi(recommend.Limit)

		// println("\n", recommend.Limit, val, "\n")

		if val > 0 {
			// println("\n", "aaa", recommend.Limit, val, "\n")
			LIMIT = "limit " + recommend.Limit
		} else {
			return "", false
		}
	}

	return LIMIT, true
}
