package Sex

import "highload2018/db/group/Structs"

func HandleSex(WHERE []string, group structs.Group) ([]string, bool) {
	if group.Sex != "" {
		WHERE = append(WHERE, "sex='"+group.Sex+"'")
	}

	return WHERE, true
}
