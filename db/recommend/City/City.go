package City

import (
	"highload2018/structs"
)

func HandleCity(WHERE []string, recommend structs.Recommend) ([]string, bool) {
	if recommend.City != "" {
		WHERE = append(WHERE, "city='"+recommend.City+"'")
	}
	// else if recommend.CityNull != "" {
	// 	if recommend.CityNull == "0" {
	// 		WHERE = append(WHERE, "city IS NOT NULL")
	// 	} else if recommend.CityNull == "1" {
	// 		WHERE = append(WHERE, "city IS NULL")
	// 	}
	// } else if recommend.CityAny != "" {
	// 	s := strings.Split(recommend.CityAny, ",")

	// 	ss := ""
	// 	for k := range s {
	// 		ss += "'" + s[k] + "',"
	// 		// 	// element is the element from someSlice for where we are
	// 	}
	// 	ss = ss[:len(ss)-1]
	// 	// ss := strings.Join(s, ",")
	// 	fmt.Println("%+v", s)
	// 	WHERE = append(WHERE, "city IN ("+ss+")")
	// }

	return WHERE, true
}
