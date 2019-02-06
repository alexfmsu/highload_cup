package Phone

import "highload2018/db/filter/Structs"

func HandlePhone(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.PhoneNull != "" {
		if filter.PhoneNull == "0" {
			WHERE = append(WHERE, "phone IS NOT NULL")
		} else if filter.PhoneNull == "1" {
			WHERE = append(WHERE, "phone IS NULL")
		}
	} else if filter.PhoneCode != "" {
		WHERE = append(WHERE, "phone LIKE '%("+filter.PhoneCode+")%'")
	}

	return WHERE, true
}
