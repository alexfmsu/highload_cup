package Sname

import "highload2018/db/filter/Structs"

func HandleSname(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.SnameEq != "" {
		WHERE = append(WHERE, "sname='"+filter.SnameEq+"'")
	} else if filter.SnameNull != "" {
		if filter.SnameNull == "0" {
			WHERE = append(WHERE, "sname IS NOT NULL")
		} else if filter.SnameNull == "1" {
			WHERE = append(WHERE, "sname IS NULL")
		}
	} else if filter.SnameStarts != "" {
		WHERE = append(WHERE, "sname LIKE '"+filter.SnameStarts+"%'")
	}

	return WHERE, true
}
