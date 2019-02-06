package Country

import "highload2018/db/filter/Structs"

func HandleCountry(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.CountryEq != "" {
		WHERE = append(WHERE, "country='"+filter.CountryEq+"'")
	} else if filter.CountryNull != "" {
		if filter.CountryNull == "0" {
			WHERE = append(WHERE, "country IS NOT NULL")
		} else if filter.CountryNull == "1" {
			WHERE = append(WHERE, "country IS NULL")
		}
	}

	return WHERE, true
}
