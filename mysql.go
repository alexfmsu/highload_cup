package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	Db, err := sql.Open("mysql", "alexfmsu:321678@/highload2018?charset=utf8")

	checkErr(err)

	return Db
}

func main() {
	Db := ConnectDB()

	id := "11554"

	var cnt int
	_ = Db.QueryRow(`select count(*) from accounts where id=` + id).Scan(&cnt)
	println(cnt)
	// smth, err := Db.Prepare("SELECT * FROM accounts WHERE id=?")
	// checkErr(err)

	// // _ = smth
	// res, err := smth.Exec(id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// println("id: ", id, ", affected: ", affect)

	defer Db.Close()
	// if affect != 1 {

	// 	// return false
	// }

	// Db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
