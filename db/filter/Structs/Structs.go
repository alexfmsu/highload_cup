package structs

import "database/sql"

type Account2 struct {
	Id      uint32            `json:id`
	Email   string            `json:email`
	Fname   string            `json:fname`
	Sname   string            `json:sname`
	Phone   string            `json:phone`
	Sex     string            `json:sex`
	Birth   uint64            `json:birth`
	Country string            `json:country`
	City    string            `json:city`
	Joined  string            `json:joined`
	Status  string            `json:status`
	Premium map[string]string `json:premium`
}

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
}

type Filter struct {
	// 1
	SexEq string

	// 2
	EmailDomain string
	EmailLt     string
	EmailGt     string

	// 3
	StatusEq  string
	StatusNeq string

	// 4
	FnameEq   string
	FnameAny  string
	FnameNull string

	// 5
	SnameEq     string
	SnameStarts string
	SnameNull   string

	// 6
	PhoneCode string
	PhoneNull string

	// 7
	CountryEq   string
	CountryNull string

	// 8
	CityEq   string
	CityAny  string
	CityNull string

	// 9
	BirthYear string
	BirthLt   string
	BirthGt   string

	// 10
	LikesContains string

	InterestsContains string
	InterestsAny      string

	PremiumNow  string
	PremiumNull string

	Limit string
}
