package Birth

import "highload2018/db/group/Structs"

func HandleBirth(WHERE []string, group structs.Group) ([]string, bool) {
	if group.Birth != "" {
		WHERE = append(WHERE, "YEAR(convert_tz(FROM_UNIXTIME(birth),@@session.time_zone,'+0:00'))="+group.Birth)
		// WHERE = append(WHERE, "FROM_UNIXTIME(birth, \"%Y\")="+group.Birth)
	}

	return WHERE, true
}
