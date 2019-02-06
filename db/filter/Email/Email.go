package Email

import "highload2018/db/filter/Structs"

func HandleEmail(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.EmailDomain != "" {
		WHERE = append(WHERE, "email LIKE '%"+filter.EmailDomain+"'")
	} else if filter.EmailGt != "" {
		WHERE = append(WHERE, "email>'"+filter.EmailGt+"'")
	} else if filter.EmailLt != "" {
		WHERE = append(WHERE, "email<'"+filter.EmailLt+"'")
	}

	return WHERE, true
}
