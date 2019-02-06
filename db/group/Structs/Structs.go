package structs

import "database/sql"

type AccountNull struct {
	Id            uint32
	Email         sql.NullString
	Fname         sql.NullString
	Sname         sql.NullString
	Phone         sql.NullString
	Sex           sql.NullString
	Birth         sql.NullInt64
	Country       sql.NullString
	City          sql.NullString
	Joined        sql.NullString
	Status        sql.NullString
	PremiumStart  sql.NullString
	PremiumFinish sql.NullString
	Interests     sql.NullString
}

type Group struct {
	// 1
	Keys      string
	Interests string
	Joined    string
	Order     string
	Birth     string
	Country   string
	Sex       string
	Status    string
	City      string
	Sname     string
	Fname     string
	Likes     string

	Limit string
}
