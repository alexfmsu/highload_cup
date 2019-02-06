package Sname

import "highload2018/db/group/Structs"

func HandleSname(WHERE []string, group structs.Group) ([]string, bool) {
	if group.Sname != "" {
		WHERE = append(WHERE, "sname='"+group.Sname+"'")
	}

	return WHERE, true
}
