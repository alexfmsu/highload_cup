package Status

import "highload2018/db/filter/Structs"

func HandleStatus(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.StatusEq != "" {
		WHERE = append(WHERE, "status='"+filter.StatusEq+"'")
	} else if filter.StatusNeq != "" {
		WHERE = append(WHERE, "status<>'"+filter.StatusNeq+"'")
	}

	return WHERE, true
}
