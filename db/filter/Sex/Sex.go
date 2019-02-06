package Sex

import "highload2018/db/filter/Structs"

func HandleSex(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.SexEq != "" {
		WHERE = append(WHERE, "sex='"+filter.SexEq+"'")
	}

	return WHERE, true
}
