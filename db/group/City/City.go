package City

import "highload2018/db/group/Structs"

func HandleCity(WHERE []string, group structs.Group) ([]string, bool) {
	if group.City != "" {
		WHERE = append(WHERE, "city='"+group.City+"'")
	}

	return WHERE, true
}
