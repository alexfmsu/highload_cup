package Birth

import "highload2018/db/filter/Structs"

func HandleBirth(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.BirthYear != "" {
		WHERE = append(WHERE, "YEAR(convert_tz(FROM_UNIXTIME(birth),@@session.time_zone,'+0:00'))="+filter.BirthYear)
		// WHERE = append(WHERE, "FROM_UNIXTIME(birth, \"%Y\")="+filter.BirthYear)
	} else if filter.BirthLt != "" {
		WHERE = append(WHERE, "birth < "+filter.BirthLt)
	} else if filter.BirthGt != "" {
		WHERE = append(WHERE, "birth > "+filter.BirthGt)
	}

	return WHERE, true
}
