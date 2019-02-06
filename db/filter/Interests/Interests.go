package Interests

import (
	"highload2018/db/filter/Structs"
	"strconv"
	"strings"
)

func HandleInterests(WHERE []string, filter structs.Filter) (string, string, []string, string, bool) {
	SELECT := ""
	FROM := ""
	GROUP_BY := ""

	if filter.InterestsContains != "" {
		// println("\nbbbbbbbb\n")
		s := strings.Split(filter.InterestsContains, ",")

		len_ := len(s)

		ss := ""
		for k := range s {
			ss += "'" + s[k] + "',"
		}

		ss = ss[:len(ss)-1]

		SELECT = "accounts.*"
		GROUP_BY = "GROUP BY accounts.ID HAVING COUNT(DISTINCT interests.interest)=" + strconv.Itoa(len_)

		WHERE = append(WHERE, "interests.interest IN ("+ss+")")
		WHERE = append(WHERE, "accounts.id=interests.id1")

		FROM = "accounts, interests"
	} else if filter.InterestsAny != "" {
		// println("\nbbbbbbbb\n")
		s := strings.Split(filter.InterestsAny, ",")

		ss := ""
		for k := range s {
			ss += "'" + s[k] + "',"
		}

		ss = ss[:len(ss)-1]

		SELECT = "accounts.*"
		GROUP_BY = "GROUP BY accounts.ID"

		WHERE = append(WHERE, "interests.interest IN ("+ss+")")
		WHERE = append(WHERE, "accounts.id=interests.id1")

		FROM = "accounts, interests"
	}

	return SELECT, FROM, WHERE, GROUP_BY, true
}
