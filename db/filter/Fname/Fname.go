package Fname

import (
	"fmt"
	"highload2018/db/filter/Structs"
	"strings"
)

func HandleFname(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.FnameEq != "" {
		WHERE = append(WHERE, "fname='"+filter.FnameEq+"'")
	} else if filter.FnameNull != "" {
		if filter.FnameNull == "0" {
			WHERE = append(WHERE, "fname IS NOT NULL")
		} else if filter.FnameNull == "1" {
			WHERE = append(WHERE, "fname IS NULL")
		}
	} else if filter.FnameAny != "" {
		s := strings.Split(filter.FnameAny, ",")

		ss := ""
		for k := range s {
			ss += "'" + s[k] + "',"
			// 	// element is the element from someSlice for where we are
		}
		ss = ss[:len(ss)-1]
		// ss := strings.Join(s, ",")
		fmt.Println("%+v", s)
		WHERE = append(WHERE, "fname IN ("+ss+")")
	}

	return WHERE, true
}
