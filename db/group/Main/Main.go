package Main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	group_birth "highload2018/db/group/Birth"
	group_city "highload2018/db/group/City"
	group_country "highload2018/db/group/Country"
	group_joined "highload2018/db/group/Joined"
	group_limit "highload2018/db/group/Limit"
	group_sex "highload2018/db/group/Sex"
	group_sname "highload2018/db/group/Sname"
	group_status "highload2018/db/group/Status"
	group_structs "highload2018/db/group/Structs"
	"highload2018/structs"
	"strings"
)

type Map map[string]interface{}

func GetSex(WHERE []string, group group_structs.Group) ([]string, bool) {
	WHERE, _ = group_sex.HandleSex(WHERE, group)

	return WHERE, true
}

func GetSname(WHERE []string, group group_structs.Group) ([]string, bool) {
	WHERE, _ = group_sname.HandleSname(WHERE, group)

	return WHERE, true
}

func Select(Db *sql.DB, group group_structs.Group) ([]byte, bool) {
	WHERE := []string{}

	SELECT := ""
	FROM := ""
	GROUP_BY := ""
	// ----------------------------------------------------------------------------------
	// 1 sex

	// eq - соответствие конкретному полу - "m" или "f";

	// WHERE, ok := sex.HandleSex(WHERE, group)
	WHERE, ok := GetSex(WHERE, group)

	if ok == false {
		return []byte(""), false
	}

	WHERE, ok = GetSname(WHERE, group)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 7 country

	// eq - всех, кто живёт в конкретной стране;
	// null - аналогично;
	WHERE, ok = group_country.HandleCountry(WHERE, group)

	if ok == false {
		return []byte(""), false
	}

	WHERE, ok = group_joined.HandleJoined(WHERE, group)

	if ok == false {
		return []byte(""), false
	}

	WHERE, ok = group_status.HandleStatus(WHERE, group)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 8 city

	// eq - всех, кто живёт в конкретном городе;
	// any - в любом из перечисленных через запятую городов;
	// null - аналогично;

	WHERE, ok = group_city.HandleCity(WHERE, group)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 9 birth

	// lt - выбрать всех, кто родился до указанной даты;
	// gt - после указанной даты;
	// year - кто родился в указанном году;

	WHERE, ok = group_birth.HandleBirth(WHERE, group)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// * limit
	LIMIT, ok := group_limit.HandleLimit(group)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	keys := strings.Split(group.Keys, ",")

	select_ := ""

	if len(keys) != 0 {
		for _, v := range keys {
			if v == "interests" {
				select_ += "interests.interest" + ","
			} else {
				select_ += v + ","
			}
		}
	}

	SELECT = " * "

	query := ""

	if FROM == "" {
		FROM = "accounts"
	}

	where := ""
	where = strings.Join(WHERE, " AND ")

	if len(WHERE) > 0 {
		query = "SELECT " + SELECT + " "

		if where != "" {
			query += " WHERE " + where
		}

		query += " FROM " + FROM + " "
		query += GROUP_BY + " ORDER BY ID DESC " + LIMIT

	} else {
		query = "SELECT " + SELECT + " "
		query += " FROM " + FROM + " "
		query += GROUP_BY + " ORDER BY ID DESC " + LIMIT
	}

	joined_interests := false

	if group.Interests != "" {
		// if v == "interests" {
		joined_interests = true

		where += " AND " + "interests.interest='" + group.Interests + "' "
		// WHERE = append(WHERE, "interests.interest="+group.Interests)
		// break
		// }
	}
	key_interests := false
	for _, v := range keys {
		if v == "interests" {
			joined_interests = true
			key_interests = true

			// where += " AND " + "interests.interest='" + group.Interests + "' "
			break
		}
	}

	if select_ == "interests," {
		// query = "SELECT country FROM accounts WHERE country='Румания' AND FROM_UNIXTIME(birth, '%Y')=2018  ORDER BY ID DESC limit 45;"
	} else {
		println("dsssssssssssssssssss\n\n\n\n\n\n", joined_interests)

		SELECT = "accounts.*" + ",count(DISTINCT accounts.id)"
		order := "count(*)"

		if group.Order == "-1" {
			order += " DESC,"

			for _, v := range keys {
				if v == "interests" {
					order += "interests.interest" + " DESC,"
				} else {
					order += v + " DESC,"
				}
			}
		} else {
			order += " ASC,"

			for _, v := range keys {
				if v == "interests" {
					order += "interests.interest" + " ASC,"
				} else {
					order += v + " ASC,"
				}
			}
		}

		join := ""
		if joined_interests == true {
			join += " JOIN interests ON accounts.id=interests.id1 "

		}
		if group.Likes != "" {
			join += " JOIN likes ON accounts.id=likes.id1 WHERE likes.id2=" + group.Likes
		}

		order = order[0 : len(order)-1]

		GROUP_BY = "GROUP BY " + select_[0:len(select_)-1]

		query = "SELECT " + SELECT + " "
		query += "FROM " + FROM + " "

		if join != "" {
			query += join + " "

			if where != "" {
				query += " AND " + where + " "
			}
		} else {
			if where != "" {
				query += "WHERE " + where + " "
			}
		}

		if group.Order == "-1" {
			query += " " + GROUP_BY + " ORDER BY " + order + " " + LIMIT
		} else {
			query += " " + GROUP_BY + " ORDER BY " + order + " " + LIMIT
		}
	}

	// if select_ == "interests," {
	// 	println("ssssssssssssssssssss\n\n\n\n\n\n")
	// 	SELECT = "accounts.*" + ",count(DISTINCT accounts.id)"
	// 	order := "count(*)"

	// 	if group.Order == "-1" {
	// 		order += " DESC,"

	// 		for _, v := range keys {
	// 			order += v + " DESC,"
	// 		}
	// 	} else {
	// 		order += " ASC,"

	// 		for _, v := range keys {
	// 			if v == "interests" {
	// 				order += "interests.interest" + " ASC,"

	// 			} else {
	// 				order += v + " ASC,"
	// 			}
	// 		}
	// 	}

	// 	join := " JOIN interests ON accounts.id=interests.id "

	// 	order = order[0 : len(order)-1]

	// 	// GROUP_BY = "GROUP BY " + select_[0:len(select_)-1]
	// 	GROUP_BY = "GROUP BY interests.interest"

	// 	query = "SELECT " + SELECT + " "
	// 	query += "FROM " + FROM + " "

	// 	if join != "" {
	// 		query += join + " "

	// 		if where != "" {
	// 			query += " AND " + where + " "
	// 		}
	// 	} else {
	// 		if where != "" {
	// 			query += "WHERE " + where + " "
	// 		}
	// 	}

	// 	if group.Order == "-1" {
	// 		query += " " + GROUP_BY + " ORDER BY " + order + " " + LIMIT
	// 	} else {
	// 		query += " " + GROUP_BY + " ORDER BY " + order + " " + LIMIT
	// 	}
	// }

	println("--------------------------------------------------------------------------")
	println("QUERY: ", query, "\n")
	fmt.Printf("Group: %+v", group, "\n")
	println("--------------------------------------------------------------------------")
	println("")

	stmt, err := Db.Query(query)

	checkErr(err)

	var accounts []Map

	count := ""

	for stmt.Next() {
		var acc structs.AccountNull

		Acc := make(Map, 0)

		if key_interests == false {
			err = stmt.Scan(&acc.Id, &acc.Email, &acc.Fname, &acc.Sname, &acc.Phone, &acc.Sex, &acc.Birth, &acc.Country, &acc.City, &acc.Joined, &acc.Status, &acc.PremiumStart, &acc.PremiumFinish, &count)
		} else {
			err = stmt.Scan(&acc.Interests, &acc.Id, &acc.Email, &acc.Fname, &acc.Sname, &acc.Phone, &acc.Sex, &acc.Birth, &acc.Country, &acc.City, &acc.Joined, &acc.Status, &acc.PremiumStart, &acc.PremiumFinish, &count)

		}
		checkErr(err)

		for _, v := range keys {
			if v == "country" && acc.Country.String != "" {
				Acc["country"] = acc.Country.String
			} else if v == "status" && acc.Status.String != "" {
				Acc["status"] = acc.Status.String
			} else if v == "sex" {
				Acc["sex"] = acc.Sex.String
			} else if v == "city" && acc.City.String != "" {
				Acc["city"] = acc.City.String
			} else if v == "interests" && acc.Interests.String != "" {
				Acc["interests"] = acc.Interests.String
			}
		}

		Acc["count"] = count

		accounts = append(accounts, Acc)
	}

	d := make(map[string][]Map, 0)

	if len(accounts) == 0 {
		accounts = make([]Map, 0)
	}

	d["groups"] = accounts

	js, _ := json.Marshal(d)

	return js, true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
