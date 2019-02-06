package Country

import "highload2018/db/group/Structs"

func HandleCountry(WHERE []string, group structs.Group) ([]string, bool) {
	if group.Country != "" {
		WHERE = append(WHERE, "country='"+group.Country+"'")
	}

	return WHERE, true
}
