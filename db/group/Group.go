package Country

import "highload2018/structs"

func HandleCountry(WHERE []string, group structs.Group) ([]string, bool) {
	if group.CountryEq != "" {
		WHERE = append(WHERE, "country='"+group.CountryEq+"'")
	} else if group.CountryNull != "" {
		if group.CountryNull == "0" {
			WHERE = append(WHERE, "country IS NOT NULL")
		} else if group.CountryNull == "1" {
			WHERE = append(WHERE, "country IS NULL")
		}
	}

	return WHERE, true
}
