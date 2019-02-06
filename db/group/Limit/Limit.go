package Limit

import (
	"highload2018/db/group/Structs"
	"strconv"
)

func HandleLimit(group structs.Group) (string, bool) {
	LIMIT := ""

	if group.Limit != "" {
		val, _ := strconv.Atoi(group.Limit)

		if val > 0 {
			LIMIT = "limit " + group.Limit
		} else {
			return "", false
		}
	}

	return LIMIT, true
}
