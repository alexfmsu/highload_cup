package DB

import (
	"database/sql"
	"encoding/json"
	"fmt"
	recommend_city "highload2018/db/recommend/City"
	recommend_limit "highload2018/db/recommend/Limit"
	"highload2018/structs"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() {
	Db, _ = sql.Open("mysql", "alexfmsu:321678@/highload2018?charset=utf8")

	// checkErr(err)

	_ = Db
}

func get_interests_by_id(id string) []string {
	query := "SELECT DISTINCT interest FROM interests WHERE ID1=" + id

	smth, err := Db.Query(query)

	checkErr(err)

	interests_ := []string{}

	interest_ := ""

	for smth.Next() {
		err := smth.Scan(&interest_)

		checkErr(err)

		interests_ = append(interests_, interest_)
	}

	return interests_
}

type Map map[string]interface{}

// func get_interests_by_id(id string) []string {
// 	query := "SELECT DISTINCT interest FROM interests WHERE ID1=" + id

// 	smth, err := db.Query(query)

// 	checkErr(err)

// 	interests_ := []string{}

// 	interest_ := ""

// 	for smth.Next() {
// 		err := smth.Scan(&interest_)

// 		checkErr(err)

// 		interests_ = append(interests_, interest_)
// 	}

// 	return interests_
// }

func SelectRecommend(recommend structs.Recommend) ([]byte, bool) {
	// ----------------------------------------------------------
	qquery := "SELECT sex FROM accounts WHERE ID=" + recommend.Id

	println("\n", qquery, "\n")

	smth2, err2 := Db.Query(qquery)

	checkErr(err2)

	sex_ := ""

	if smth2.Next() == false {
		return []byte(""), false
	} else {
		err := smth2.Scan(&sex_)

		checkErr(err)

		// fmt.Println(recommend.Id)
	}

	if sex_ == "m" {
		sex_ = "f"
	} else if sex_ == "f" {
		sex_ = "m"
	}
	// ----------------------------------------------------------
	interests_ := get_interests_by_id(recommend.Id)
	fmt.Printf("%+v", interests_)

	// if Next() {

	WHERE := []string{}
	WHERE = append(WHERE, "sex='"+sex_+"'")

	SELECT := ""
	FROM := ""
	GROUP_BY := ""

	// ----------------------------------------------------------------------------------
	// 1 sex

	// eq - соответствие конкретному полу - "m" или "f";

	// WHERE, ok := sex.HandleSex(WHERE, filter)
	// fmt.Printf("%+v\n", WHERE)
	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 2 email

	// domain - выбрать всех, чьи email-ы имеют указанный домен;
	// lt - выбрать всех, чьи email-ы лексикографически раньше;
	// gt - то же, но лексикографически позже;

	// WHERE, ok = email.HandleEmail(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 3 status

	// eq - соответствие конкретному статусу;
	// neq - выбрать всех, чей статус не равен указанному;

	// WHERE, ok = status.HandleStatus(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 4 fname

	// eq - соответствие конкретному имени;
	// any - соответствие любому имени из перечисленных через запятую;
	// null - выбрать всех, у кого указано имя (если 0) или не указано (если 1);

	// WHERE, ok = fname.HandleFname(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 5 sname

	// eq - соответствие конкретной фамилии;
	// starts - выбрать всех, чьи фамилии начинаются с переданного префикса;
	// null - выбрать всех, у кого указана фамилия (если 0) или не указана (если 1);

	// WHERE, ok = sname.HandleSname(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 6 phone

	// code - выбрать всех, у кого в телефоне конкретный код (три цифры в скобках);
	// null - аналогично остальным полям;

	// WHERE, ok = phone.HandlePhone(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 7 country

	// eq - всех, кто живёт в конкретной стране;
	// null - аналогично;

	// WHERE, ok = country.HandleCountry(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 8 city

	// eq - всех, кто живёт в конкретном городе;
	// any - в любом из перечисленных через запятую городов;
	// null - аналогично;

	WHERE, ok := recommend_city.HandleCity(WHERE, recommend)

	if ok == false {
		return []byte(""), false
	}
	// ----------------------------------------------------------------------------------
	// 9 birth

	// lt - выбрать всех, кто родился до указанной даты;
	// gt - после указанной даты;
	// year - кто родился в указанном году;

	// WHERE, ok = birth.HandleBirth(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 10 interests

	// contains - выбрать всех, у кого есть все перечисленные интересы;
	// any - выбрать всех, у кого есть любой из перечисленных интересов;

	// SELECT, FROM, WHERE, GROUP_BY, ok = interests.HandleInterests(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 11 likes

	// contains - выбрать всех, кто лайкал всех перечисленных пользователей
	// (в значении - перечисленные через запятые id);

	// SELECT, FROM, WHERE, GROUP_BY, ok = likes_contains.HandleLikesContains(WHERE, filter)

	// if ok == false {
	// 	return []byte(""), false
	// }
	// ----------------------------------------------------------------------------------
	// 12 premium

	// now - все у кого есть премиум на текущую дату;
	// null - аналогично остальным;

	// ----------------------------------------------------------------------------------
	// * limit
	LIMIT, ok := recommend_limit.HandleLimit(recommend)

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

	where := strings.Join(WHERE, " AND ")

	query := "SELECT " + SELECT + " FROM " + FROM + " WHERE " + where + " " + GROUP_BY + " ORDER BY ID DESC " + LIMIT

	println("\n", query, "\n")

	stmt, err := Db.Query(query)

	checkErr(err)

	var accounts []Map

	for stmt.Next() {
		var acc structs.AccountNull

		Acc := make(Map, 0)

		err := stmt.Scan(&acc.Id, &acc.Email, &acc.Fname, &acc.Sname, &acc.Phone, &acc.Sex, &acc.Birth, &acc.Country, &acc.City, &acc.Joined, &acc.Status, &acc.PremiumStart, &acc.PremiumFinish)

		checkErr(err)

		Acc["id"] = acc.Id
		Acc["email"] = acc.Email.String

		// if filter.SexEq != "" {
		// 	Acc["sex"] = acc.Sex.String
		// }
		// if filter.PhoneNull == "0" {
		// 	Acc["phone"] = acc.Phone.String
		// }
		// if filter.CountryNull == "0" || filter.CountryEq != "" {
		// 	Acc["country"] = acc.Country.String
		// }
		if recommend.City != "" {
			Acc["city"] = acc.City.String
		}
		// if filter.FnameNull == "0" {
		// 	Acc["fname"] = acc.Fname.String
		// }
		// if filter.SnameStarts != "" || filter.SnameNull == "0" {
		// 	Acc["sname"] = acc.Sname.String
		// }
		// if filter.StatusEq != "" || filter.StatusNeq != "" {
		// 	Acc["status"] = acc.Status.String
		// }
		// if filter.BirthYear != "" || filter.BirthLt != "" || filter.BirthGt != "" {
		// 	Acc["birth"] = acc.Birth.Int64
		// }

		accounts = append(accounts, Acc)
	}

	// for k, v := accounts{

	// }

	d := make(map[string][]Map, 0)

	if len(accounts) == 0 {
		return []byte(""), false
		// accounts = make([]Map, 0)
	}

	// fmt.Printf("%+v", accounts)

	d["accounts"] = accounts

	js, _ := json.Marshal(d)

	fmt.Printf("%s", js)

	return js, true

}

func Insert(acc structs.Account2) bool {
	if acc.Sex != "m" && acc.Sex != "f" {
		return false
	}

	// if acc.Phone == "" {
	// 	return false
	// }
	// fmt.Print("%+v", acc)
	// fmt.Println("Email:", acc.Email)
	// fmt.Println("Fname:", acc.Fname)
	// fmt.Println("Sname:", acc.Sname)
	// fmt.Println("Phone:", acc.Phone)
	// fmt.Println("Sex:", acc.Sex)
	// fmt.Println("Birth:", acc.Birth)
	// fmt.Println("Country:", acc.Country)

	// ConnectDB
	// db_, err_ := sql.Open("mysql", "alexfmsu:321678@/highload2018?charset=utf8")

	// _, err := db.Prepare("INSERT accounts SET email=?")
	stmt, err := Db.Prepare("INSERT accounts SET id=?, email=?, fname=?,sname=?, phone=?, sex=?, birth=?, country=?, city=?")
	checkErr(err)
	// _ = smth
	res, err := stmt.Exec(acc.Id, acc.Email, acc.Fname, acc.Sname, acc.Phone, acc.Sex, acc.Birth, acc.Country, acc.City)
	checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	if res != nil {
		return true
	} else {
		return false
	}
}

func Insert2(acc structs.Account3) bool {
	if acc.Sex != "m" && acc.Sex != "f" {
		return false
	}

	stmt, err := Db.Prepare("INSERT accounts SET id=?, email=?, fname=?,sname=?, phone=?, sex=?, birth=?, country=?, city=?")
	checkErr(err)
	// _ = smth
	res, err := stmt.Exec(acc.Id, acc.Email, acc.Fname, acc.Sname, acc.Phone, acc.Sex, acc.Birth, acc.Country, acc.City)
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	if res != nil {
		return true
	} else {
		return false
	}
}

func Update(acc structs.Account2) bool {
	stmt, err := Db.Prepare("UPDATE accounts SET email=?, fname=?,sname=?, phone=?, sex=?, birth=?, country=?, city=? where id=?")
	checkErr(err)
	// _ = smth
	res, err := stmt.Exec(acc.Email, acc.Fname, acc.Sname, acc.Phone, acc.Sex, acc.Birth, acc.Country, acc.City, acc.Id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	println("id", acc.Id)
	println("affect", affect)
	// id, err := res.LastInsertId()
	// checkErr(err)

	if affect != 0 {
		return true
	} else {
		return false
	}
}

func id_exists(DB *sql.DB, id string) int {
	var cnt int

	_ = Db.QueryRow(`select count(*) from accounts where id=` + id).Scan(&cnt)
	println("cnt=", cnt)

	return cnt
}

func Update2(id string, acc structs.Account3) (bool, int) {
	if id_exists(Db, id) == 0 {
		println("NOT EXIST ", id)
		return false, 404
	}

	fields := "UPDATE accounts SET "

	if acc.Email != "" {
		fields += "email='" + acc.Email + "',"
	}

	if acc.Status != "" {
		fields += "status='" + acc.Status + "',"
	}

	if acc.Sname != "" {
		fields += "sname='" + acc.Sname + "',"
	}

	if acc.Fname != "" {
		fields += "fname='" + acc.Fname + "',"
	}

	if acc.Country != "" {
		fields += "country='" + acc.Country + "',"
	}

	if acc.City != "" {
		fields += "city='" + acc.City + "',"
	}

	if acc.Sex != "" {
		fields += "sex='" + acc.Sex + "',"
	}

	if acc.Phone != "" {
		fields += "phone='" + acc.Phone + "',"
	}

	if acc.Premium["start"] != 0 {
		fields += "premium_start='" + fmt.Sprint(acc.Premium["start"]) + "',"
	}
	if acc.Premium["finish"] != 0 {
		fields += "premium_finish='" + fmt.Sprint(acc.Premium["finish"]) + "',"
	}

	// if acc.Joined != "" {
	// 	fields += "joined='" + acc.Joined + "',"
	// }

	println("----------------------------------------------------")
	// println("QUERY: ", fields[0:len(fields)-1]+" "+"WHERE id="+id)

	smth, err := Db.Query(fields[0:len(fields)-1] + " " + "WHERE id=" + id)

	if err != nil {
		return false, 400
	}
	// checkErr(err)
	// println("ok")

	for smth.Next() {
		return true, 202
	}

	// println("ok")

	return true, 202
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
