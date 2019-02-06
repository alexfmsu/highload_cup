package LikesContains

import (
	"fmt"
	"highload2018/db/filter/Structs"
	"strconv"
	"strings"
)

func HandleLikesContains(WHERE []string, filter structs.Filter) (string, string, []string, string, bool) {
	SELECT := ""
	FROM := ""
	GROUP_BY := ""

	if filter.LikesContains != "" {
		s := strings.Split(filter.LikesContains, ",")

		len_ := len(s)

		ss := ""
		for k := range s {
			ss += "'" + s[k] + "',"
		}

		ss = ss[:len(ss)-1]
		fmt.Printf("ss=%+v", ss)
		SELECT = "accounts.*"
		GROUP_BY = "GROUP BY accounts.ID HAVING COUNT(DISTINCT likes.id1)=" + strconv.Itoa(len_)
		// GROUP_BY = "GROUP BY accounts.ID")

		WHERE = append(WHERE, "likes.id2 IN ("+ss+")")
		WHERE = append(WHERE, "accounts.id=likes.id1")

		// FROM = "accounts, likes"
		FROM = "accounts"
	}

	return SELECT, FROM, WHERE, GROUP_BY, true
}
