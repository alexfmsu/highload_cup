package Premium

import "highload2018/db/filter/Structs"

func HandlePremium(WHERE []string, filter structs.Filter) ([]string, bool) {
	if filter.PremiumNow != "" {
		WHERE = append(WHERE, "premium_start < 1548341893 AND premium_finish > 1548341893")
	} else if filter.PremiumNull != "" {
		if filter.PremiumNull == "0" {
			WHERE = append(WHERE, "(premium_start IS NOT NULL OR premium_finish IS NOT NULL)")
		} else if filter.PremiumNull == "1" {
			WHERE = append(WHERE, "(premium_start IS NULL OR premium_finish IS NULL)")
		}
	}

	return WHERE, true
}
