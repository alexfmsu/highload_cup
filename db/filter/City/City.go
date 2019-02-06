package City

import (
	"fmt"
	"highload2018/db/filter/Structs"
	"strings"
)

func HandleCity(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.CityEq != "" {
		WHERE = append(WHERE, "city='"+filter.CityEq+"'")
	} else if filter.CityNull != "" {
		if filter.CityNull == "0" {
			WHERE = append(WHERE, "city IS NOT NULL")
		} else if filter.CityNull == "1" {
			WHERE = append(WHERE, "city IS NULL")
		}
	} else if filter.CityAny != "" {
		s := strings.Split(filter.CityAny, ",")

		ss := ""
		for k := range s {
			ss += "'" + s[k] + "',"
			// 	// element is the element from someSlice for where we are
		}
		ss = ss[:len(ss)-1]
		// ss := strings.Join(s, ",")
		fmt.Println("%+v", s)
		WHERE = append(WHERE, "city IN ("+ss+")")
	}

	return WHERE, true
}
