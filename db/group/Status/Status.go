package Status

import "highload2018/db/group/Structs"

func HandleStatus(WHERE []string, group structs.Group) ([]string, bool) {
	if group.Status != "" {
		WHERE = append(WHERE, "status='"+group.Status+"'")
	}

	return WHERE, true
}
