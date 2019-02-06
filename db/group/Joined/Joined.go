package Joined

import "highload2018/db/group/Structs"

func HandleJoined(WHERE []string, group structs.Group) ([]string, bool) {
	if group.Joined != "" {
		WHERE = append(WHERE, "FROM_UNIXTIME(joined, \"%Y\")="+group.Joined)
	}

	return WHERE, true
}
