package Main

import (
	"database/sql"
	"encoding/json"
	filter_birth "highload2018/db/filter/Birth"
	filter_city "highload2018/db/filter/City"
	filter_country "highload2018/db/filter/Country"
	filter_email "highload2018/db/filter/Email"
	filter_fname "highload2018/db/filter/Fname"
	interests "highload2018/db/filter/Interests"
	likes_contains "highload2018/db/filter/LikesContains"
	filter_limit "highload2018/db/filter/Limit"
	filter_phone "highload2018/db/filter/Phone"
	filter_premium "highload2018/db/filter/Premium"
	filter_sex "highload2018/db/filter/Sex"
	filter_sname "highload2018/db/filter/Sname"
	filter_status "highload2018/db/filter/Status"
	"highload2018/structs"
	"strings"

	filter_struct "highload2018/db/filter/Structs"

	_ "github.com/go-sql-driver/mysql"
)

type Map map[string]interface{}

// ----------------------------------------------------------------------------------
// 1 sex

// eq - соответствие конкретному полу - "m" или "f";

func GetSex(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_sex.HandleSex(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 2 email

// domain - выбрать всех, чьи email-ы имеют указанный домен;
// lt - выбрать всех, чьи email-ы лексикографически раньше;
// gt - то же, но лексикографически позже;

func GetEmail(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_email.HandleEmail(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 3 status

// eq - соответствие конкретному статусу;
// neq - выбрать всех, чей статус не равен указанному;

func GetStatus(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_status.HandleStatus(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 4 fname

// eq - соответствие конкретному имени;
// any - соответствие любому имени из перечисленных через запятую;
// null - выбрать всех, у кого указано имя (если 0) или не указано (если 1);

func GetFname(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_fname.HandleFname(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 5 sname

// eq - соответствие конкретной фамилии;
// starts - выбрать всех, чьи фамилии начинаются с переданного префикса;
// null - выбрать всех, у кого указана фамилия (если 0) или не указана (если 1);

func GetSname(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_sname.HandleSname(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 6 phone

// code - выбрать всех, у кого в телефоне конкретный код (три цифры в скобках);
// null - аналогично остальным полям;

func GetPhone(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_phone.HandlePhone(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 7 country

// eq - всех, кто живёт в конкретной стране;
// null - аналогично;

func GetCountry(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_country.HandleCountry(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 8 city

// eq - всех, кто живёт в конкретном городе;
// any - в любом из перечисленных через запятую городов;
// null - аналогично;

func GetCity(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_city.HandleCity(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
// 9 birth

// lt - выбрать всех, кто родился до указанной даты;
// gt - после указанной даты;
// year - кто родился в указанном году;

func GetBirth(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_birth.HandleBirth(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
func GetPremium(WHERE []string, filter filter_struct.Filter) ([]string, bool) {
	WHERE, _ = filter_premium.HandlePremium(WHERE, filter)

	return WHERE, true
}

// ----------------------------------------------------------------------------------
func GetLimit(filter filter_struct.Filter) (string, bool) {
	LIMIT, ok := filter_limit.HandleLimit(filter)

	return LIMIT, ok
}

// ----------------------------------------------------------------------------------

func Select(db *sql.DB, filter filter_struct.Filter) ([]byte, bool) {
	WHERE := []string{}

	SELECT := ""
	FROM := ""
	GROUP_BY := ""

	var ok bool

	// ----------------------------------------------------------------------------------
	// 1 sex
	WHERE, ok = GetSex(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 2 email
	WHERE, ok = GetEmail(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 3 status
	WHERE, ok = GetStatus(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 4 fname
	WHERE, ok = GetFname(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 5 sname
	WHERE, ok = GetSname(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 6 phone
	WHERE, ok = GetPhone(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 7 country
	WHERE, ok = GetCountry(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 8 city
	WHERE, ok = GetCity(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 9 birth
	WHERE, ok = GetBirth(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 11 likes

	// contains - выбрать всех, кто лайкал всех перечисленных пользователей
	// (в значении - перечисленные через запятые id);

	SELECT, FROM, WHERE, GROUP_BY, ok = likes_contains.HandleLikesContains(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 10 interests

	// contains - выбрать всех, у кого есть все перечисленные интересы;
	// any - выбрать всех, у кого есть любой из перечисленных интересов;

	SELECT, FROM, WHERE, GROUP_BY, ok = interests.HandleInterests(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------

	// ----------------------------------------------------------------------------------
	// 12 premium

	// now - все у кого есть премиум на текущую дату;
	// null - аналогично остальным;
	WHERE, ok = GetPremium(WHERE, filter)
	// WHERE, ok = premium.HandlePremium(WHERE, filter)
	// WHERE, ok = premium.HandlePremium(WHERE, filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// * limit
	LIMIT, ok := GetLimit(filter)
	// LIMIT, ok := limit.HandleLimit(filter)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------

	if SELECT == "" {
		SELECT = "*"
	}
	if FROM == "" {
		FROM = "accounts"
	}

	where := ""
	query := ""

	if len(WHERE) > 0 {
		where = strings.Join(WHERE, " AND ")

		if filter.LikesContains != "" {
			query = "SELECT " + "accounts.*"
		} else {
			query = "SELECT " + SELECT
		}
		query += " FROM " + FROM

		if filter.LikesContains != "" {
			query += " JOIN likes ON accounts.id = likes.id1 "
		}

		query += " WHERE " + where
		query += " " + GROUP_BY + " ORDER BY accounts.id DESC " + LIMIT
	} else {
		query = "SELECT " + SELECT + " FROM " + FROM + " " + GROUP_BY + " ORDER BY accounts.id DESC " + LIMIT
	}

	// println("\nQUERY: ", query, "\n")

	stmt, err := db.Query(query)

	checkErr(err)

	var accounts []Map

	for stmt.Next() {
		var acc structs.AccountNull

		Acc := make(Map, 0)

		err := stmt.Scan(&acc.Id, &acc.Email, &acc.Fname, &acc.Sname, &acc.Phone, &acc.Sex, &acc.Birth, &acc.Country, &acc.City, &acc.Joined, &acc.Status, &acc.PremiumStart, &acc.PremiumFinish)

		checkErr(err)

		Acc["id"] = acc.Id
		Acc["email"] = acc.Email.String

		if filter.SexEq != "" {
			Acc["sex"] = acc.Sex.String
		}
		if filter.PhoneNull == "0" || filter.PhoneCode != "" {
			Acc["phone"] = acc.Phone.String
		}
		if filter.CountryNull == "0" || filter.CountryEq != "" {
			Acc["country"] = acc.Country.String
		}
		if filter.CityNull == "0" || filter.CityEq != "" || filter.CityAny != "" {
			Acc["city"] = acc.City.String
		}
		if filter.FnameNull == "0" || filter.FnameAny != "" {
			Acc["fname"] = acc.Fname.String
		}
		if filter.SnameStarts != "" || filter.SnameNull == "0" {
			Acc["sname"] = acc.Sname.String
		}
		if filter.StatusEq != "" || filter.StatusNeq != "" {
			Acc["status"] = acc.Status.String
		}
		if filter.BirthYear != "" || filter.BirthLt != "" || filter.BirthGt != "" {
			Acc["birth"] = acc.Birth.Int64
		}
		if filter.PremiumNow != "" || filter.PremiumNull == "0" {
			Acc["premium"] = map[string]string{
				"start":  acc.PremiumStart.String,
				"finish": acc.PremiumFinish.String,
			}
		}

		accounts = append(accounts, Acc)
	}

	d := make(map[string][]Map, 0)

	if len(accounts) == 0 {
		accounts = make([]Map, 0)
	}

	d["accounts"] = accounts

	js, _ := json.Marshal(d)

	return js, true

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
