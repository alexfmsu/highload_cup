package main

import (
	"encoding/json"
	"fmt"
	"highload2018/db"
	filter_main "highload2018/db/filter/Main"
	filter_structs "highload2018/db/filter/Structs"
	group_main "highload2018/db/group/Main"
	group_structs "highload2018/db/group/Structs"
	"highload2018/structs"

	"strings"

	"log"
	"regexp"
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var (
	group_allowed_keys = []string{
		"birth",
		"order",
		"limit",
		"keys",
		"sex",
		"country",
		"interests",
		"joined",
		"query_id",
		"status",
		"city",
		"likes",
	}

	allowed_keys = []string{
		"sex_eq",

		"email_domain",
		"email_lt",
		"email_gt",

		"status_eq",
		"status_neq",

		"fname_eq",
		"fname_any",
		"fname_null",

		"sname_eq",
		"sname_starts",
		"sname_null",

		"phone_null",
		"phone_code",
		"limit",
		"query_id",

		"birth_year",
		"birth_lt",
		"birth_gt",

		"city_eq",
		"city_null",
		"city_any",

		"country_eq",
		"country_null",

		"likes_contains",
		"likes_any",

		"interests_contains",
		"interests_any",

		"premium_now",
		"premium_null",
	}

	re_id_recommend = regexp.MustCompile(`[^/]+/recommend`)
	re_id_suggest   = regexp.MustCompile(`[^/]+/suggest`)
)

func accounts_get(ctx *fasthttp.RequestCtx) {
	path := ctx.UserValue("path").(string)
	println(path)

	if path == "/filter/" {
		fmt.Fprint(ctx, "GET: accounts/filter\n")

		args := ctx.QueryArgs()
		Args := strings.Split(args.String(), "&")

		params := map[string]string{}

		for _, f := range Args {
			v := strings.Split(f, "=")

			if len(v) != 2 {
				return
			}

			params[v[0]] = v[1]

			found := false
			for i := range allowed_keys {
				if ok := allowed_keys[i] == v[0]; ok {
					found = true
					break
				}
			}

			if found == false {
				ctx.Response.SetStatusCode(400)
				ctx.Write([]byte(""))
				print(v[0])
				return
			}
		}

		fmt.Printf("args:%+v\n", args)
		// -----------------------------------------------------
		var filter filter_structs.Filter

		// 1
		filter.SexEq = params["sex_eq"]

		// 2
		filter.EmailDomain = string(args.Peek("email_domain"))
		filter.EmailLt = string(args.Peek("email_lt"))
		filter.EmailGt = string(args.Peek("email_gt"))

		// 3
		filter.StatusEq = string(args.Peek("status_eq"))
		filter.StatusNeq = string(args.Peek("status_neq"))

		// 4
		filter.FnameEq = string(args.Peek("fname_eq"))
		filter.FnameAny = string(args.Peek("fname_any"))
		filter.FnameNull = string(args.Peek("fname_null"))

		// 5
		filter.SnameEq = string(args.Peek("sname_eq"))
		filter.SnameStarts = string(args.Peek("sname_starts"))
		filter.SnameNull = string(args.Peek("sname_null"))

		// 6
		filter.PhoneCode = string(args.Peek("phone_code"))
		filter.PhoneNull = string(args.Peek("phone_null"))

		// 7
		filter.CountryEq = string(args.Peek("country_eq"))
		filter.CountryNull = string(args.Peek("country_null"))

		// 8
		filter.CityEq = string(args.Peek("city_eq"))
		filter.CityAny = string(args.Peek("city_any"))
		filter.CityNull = string(args.Peek("city_null"))

		// 9
		filter.BirthYear = string(args.Peek("birth_year"))
		filter.BirthLt = string(args.Peek("birth_lt"))
		filter.BirthGt = string(args.Peek("birth_gt"))

		// 10
		filter.LikesContains = string(args.Peek("likes_contains"))
		// filter.LikesContains = params["likes_contains"]

		filter.InterestsContains = string(args.Peek("interests_contains"))
		filter.InterestsAny = string(args.Peek("interests_any"))

		filter.PremiumNow = string(args.Peek("premium_now"))
		filter.PremiumNull = string(args.Peek("premium_null"))

		filter.Limit = string(args.Peek("limit"))
		// -----------------------------------------------------

		fmt.Printf("%+v\n", filter)

		json, err := filter_main.Select(DB.Db, filter)

		if err != false {
			ctx.Response.SetStatusCode(200)

			ctx.Write([]byte(json))
		} else {
			ctx.Response.SetStatusCode(400)
		}

		ctx.SetContentType("application/json; charset=utf-8")
	} else if path == "/group/" {
		// println("\n\n\nssssssssssss\n\n\n")
		fmt.Fprint(ctx, "GET: accounts/group\n")

		args := ctx.QueryArgs()
		Args := strings.Split(args.String(), "&")

		params := map[string]string{}

		for _, f := range Args {
			v := strings.Split(f, "=")

			if len(v) != 2 {
				return
			}

			params[v[0]] = v[1]

			found := false
			for i := range group_allowed_keys {
				if ok := group_allowed_keys[i] == v[0]; ok {
					found = true
					break
				}
			}

			if found == false {
				ctx.Response.SetStatusCode(400)
				ctx.Write([]byte(""))
				print(v[0])
				return
			}
		}

		fmt.Printf("args:%+v\n", args)

		var group group_structs.Group

		group.Interests = string(args.Peek("interests"))
		group.Country = string(args.Peek("country"))
		group.Sex = string(args.Peek("sex"))
		group.Keys = string(args.Peek("keys"))
		group.Birth = string(args.Peek("birth"))
		group.Joined = string(args.Peek("joined"))
		group.Limit = string(args.Peek("limit"))
		group.Status = string(args.Peek("status"))
		group.Order = string(args.Peek("order"))
		group.City = string(args.Peek("city"))
		group.Sname = string(args.Peek("sname"))
		group.Fname = string(args.Peek("fname"))
		group.Likes = string(args.Peek("likes"))

		keys := strings.Split(group.Keys, ",")

		for _, v1 := range keys {
			key_found := false

			for _, v2 := range group_allowed_keys {
				if ok := v1 == v2; ok {
					println(v1, v2)
					key_found = true
					break
				}
			}

			if key_found == false {
				ctx.Response.SetStatusCode(400)

				return
			}
		}

		fmt.Printf("%+v\n", group)

		json, err := group_main.Select(DB.Db, group)

		if err != false {
			ctx.Response.SetStatusCode(200)

			ctx.Write([]byte(json))
		} else {
			ctx.Response.SetStatusCode(400)
		}

		ctx.SetContentType("application/json; charset=utf-8")
	} else if re_id_recommend.MatchString(path) == true {
		fmt.Fprint(ctx, "GET: accounts/filter\n")

		args := ctx.QueryArgs()

		var recommend structs.Recommend

		s := strings.Split(path, "/")

		recommend.Id = s[1]

		// println(path, "\n s=", s[1], "\n")

		// fmt.Printf("%+v\n", ctx)
		// fmt.Printf("%+v\n", args)
		// -----------------------------------------------------

		// 1
		// filter.SexEq = string(args.Peek("sex_eq"))

		// if filter.SexEq == "m" {
		// 	filter.SexEq = "f"
		// } else if filter.SexEq == "f" {
		// 	filter.SexEq = "m"
		// }

		// 2
		// filter.EmailDomain = string(args.Peek("email_domain"))
		// filter.EmailLt = string(args.Peek("email_lt"))
		// filter.EmailGt = string(args.Peek("email_gt"))

		// // 3
		// filter.StatusEq = string(args.Peek("status_eq"))
		// filter.StatusNeq = string(args.Peek("status_neq"))

		// // 4
		// filter.FnameEq = string(args.Peek("fname_eq"))
		// filter.FnameAny = string(args.Peek("fname_any"))
		// filter.FnameNull = string(args.Peek("fname_null"))

		// // 5
		// filter.SnameEq = string(args.Peek("sname_eq"))
		// filter.SnameStarts = string(args.Peek("sname_starts"))
		// filter.SnameNull = string(args.Peek("sname_null"))

		// // 6
		// filter.PhoneCode = string(args.Peek("phone_code"))
		// filter.PhoneNull = string(args.Peek("phone_null"))

		// // 7
		// filter.CountryEq = string(args.Peek("country_eq"))
		// filter.CountryNull = string(args.Peek("country_null"))

		// 8
		recommend.City = string(args.Peek("city"))
		// filter.CityAny = string(args.Peek("city_any"))
		// filter.CityNull = string(args.Peek("city_null"))

		// // 9
		// filter.BirthYear = string(args.Peek("birth_year"))
		// filter.BirthLt = string(args.Peek("birth_lt"))
		// filter.BirthGt = string(args.Peek("birth_gt"))

		// // 10
		// filter.LikeContains = string(args.Peek("likes_contains"))

		// filter.InterestsContains = string(args.Peek("interests_contains"))
		// filter.InterestsAny = string(args.Peek("interests_any"))

		recommend.Limit = string(args.Peek("limit"))
		// -----------------------------------------------------
		fmt.Printf("%+v\n", recommend)

		json, err := DB.SelectRecommend(recommend)

		if err != false {
			ctx.Response.SetStatusCode(200)

			ctx.Write([]byte(json))
		} else {
			ctx.Response.SetStatusCode(404)
		}

		ctx.SetContentType("application/json; charset=utf-8")

		// fmt.Fprint(ctx, "GET: accounts/id/recommend\n")
		// ctx.Response.SetStatusCode(502)
	} else if re_id_suggest.MatchString(path) == true {
		fmt.Fprint(ctx, "GET: accounts/id/suggest\n")
	} else {

		ctx.Response.SetStatusCode(404)

	}
}

func MARSHALL() {
	type Artist struct {
		Id int
	}

	temp := []byte(`{"Id":12}`)

	var artist Artist
	err := json.Unmarshal(temp, &artist)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Println(artist.Id)
	fmt.Printf("%+v", artist)
}

type Account3 struct {
	Id      uint32            `json:id`
	Email   string            `json:email`
	Fname   string            `json:fname`
	Sname   string            `json:sname`
	Phone   string            `json:phone`
	Sex     string            `json:sex`
	Birth   uint64            `json:birth`
	Country string            `json:country`
	City    string            `json:city`
	Joined  uint32            `json:joined`
	Status  string            `json:status`
	Premium map[string]uint32 `json:premium`
}

func ValidatePostUpdate(account structs.Account3) bool {
	if account.Status != "" {
		if account.Status != "свободны" && account.Status != "заняты" && account.Status != "всё сложно" {
			return false
		}
	}

	if account.Email != "" {
		err := checkmail.ValidateFormat(account.Email)
		if err != nil {
			return false
		}
	}

	return true
}

func accounts_post(ctx *fasthttp.RequestCtx) {
	path := ctx.UserValue("path")

	var account structs.Account3
	// var account Account3
	// var account filter_structs.Account2

	if path == "/new" {
		body := ctx.PostBody()

		if err := json.Unmarshal(body, &account); err != nil {
			print("errorr\n")
		} else {
			fmt.Printf("NEW %+v\n", account)

			if DB.Insert2(account) {
				// if correct, code := DB.Insert2(account); correct == true {
				ctx.Response.SetStatusCode(200)
			} else {
				ctx.Response.SetStatusCode(400)
			}

			ctx.SetContentType("application/json; charset=utf-8")
			ctx.Write([]byte("{}"))
		}
	} else if path == "/likes" {
		fmt.Fprint(ctx, "POST: accounts/likes\n")
	} else if path != "" {
		path := ctx.UserValue("path").(string)

		path_ := strings.Split(path, "/")

		val_, _ := strconv.ParseInt(path_[1], 10, 32)

		if val_ == 0 {
			ctx.Response.SetStatusCode(400)
			return
		}

		body := ctx.PostBody()

		if err := json.Unmarshal(body, &account); err != nil {
			println("Error: update, marshal\n")
			ctx.Response.SetStatusCode(400)
			return
		} else {
			fmt.Printf("UPDATE %+v\n", account)

			if ValidatePostUpdate(account) == false {
				ctx.Response.SetStatusCode(400)
				return
			}

			val, _ := strconv.ParseInt(path[1:], 10, 32)

			account.Id = uint32(val)

			// if account.Joined == 0 {
			// 	ctx.Response.SetStatusCode(400) //202)
			// 	return
			// }

			if correct, code := DB.Update2(path[1:], account); correct == true {
				ctx.Response.SetStatusCode(code) //202)
			} else {
				ctx.Response.SetStatusCode(code) // 404
				println("Error: update, can't update\n")
			}

			ctx.SetContentType("application/json; charset=utf-8")
			ctx.Write([]byte("{}"))
		}
	}
}

func main() {
	DB.ConnectDB()

	router := fasthttprouter.New()

	router.GET("/accounts/*path", accounts_get)
	router.POST("/accounts/*path", accounts_post)

	println("Listening...\n")

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
