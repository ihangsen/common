package time

import (
	"github.com/ihangsen/common/src/entity/enums/amuser"
	"time"
)

// StartDay 当天起始时间
func StartDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// EndDay 当天结束时间
func EndDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
}

// LastDay 上一天起始时间
func LastDay(now time.Time) time.Time {
	return now.AddDate(0, 0, -1)
}

// NextDay 第二天起始时间
func NextDay(now time.Time) time.Time {
	return now.AddDate(0, 0, 1)
}

// StartWeek 本周起始时间
func StartWeek(now time.Time) time.Time {
	weekday := int(now.Weekday())
	daysToMonday := weekday - 1
	monday := now.AddDate(0, 0, -daysToMonday)
	return time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
}

// EndWeek 本周结束时间
func EndWeek(now time.Time) time.Time {
	weekday := int(now.Weekday())
	daysToSunday := 7 - weekday
	sunday := now.AddDate(0, 0, daysToSunday)
	return time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 0, 0, 0, 0, sunday.Location())
}

// LastWeek  上周起始时间
func LastWeek(now time.Time) time.Time {
	weekday := int(now.Weekday())
	daysToLastMonday := weekday + 6
	lastMonday := now.AddDate(0, 0, -daysToLastMonday)
	return time.Date(lastMonday.Year(), lastMonday.Month(), lastMonday.Day(), 0, 0, 0, 0, lastMonday.Location())
}

// LastTwoWeek 上上周起始时间
func LastTwoWeek(now time.Time) time.Time {
	weekday := int(now.Weekday())
	daysToTwoWeeksAgoMonday := weekday + 6 + 7
	twoWeeksAgoMonday := now.AddDate(0, 0, -daysToTwoWeeksAgoMonday)
	return time.Date(twoWeeksAgoMonday.Year(), twoWeeksAgoMonday.Month(), twoWeeksAgoMonday.Day(), 0, 0, 0, 0, twoWeeksAgoMonday.Location())
}

// NextWeek 下一周起始时间
func NextWeek(now time.Time) time.Time {
	weekday := int(now.Weekday())
	daysToNextMonday := 8 - weekday
	nextMonday := now.AddDate(0, 0, daysToNextMonday)
	return time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 0, 0, 0, 0, nextMonday.Location())
}

// CalculateConstellation 星座计算
func CalculateConstellation(birthdayTimestamp int64) uint8 {
	birthday := time.Unix(birthdayTimestamp/1000, 0)

	month := birthday.Month()
	day := birthday.Day()

	switch {
	case (month == time.March && day >= 21) || (month == time.April && day <= 19):
		return amuser.Aries //"白羊座"
	case (month == time.April) || (month == time.May && day <= 20):
		return amuser.Taurus //"金牛座"
	case (month == time.May) || (month == time.June && day <= 21):
		return amuser.Gemini //"双子座"
	case (month == time.June && day >= 22) || (month == time.July && day <= 22):
		return amuser.Cancer //"巨蟹座"
	case (month == time.July && day >= 23) || (month == time.August && day <= 22):
		return amuser.Leo //"狮子座"
	case (month == time.August && day >= 23) || (month == time.September && day <= 22):
		return amuser.Virgo //"处女座"
	case (month == time.September && day >= 23) || (month == time.October && day <= 23):
		return amuser.Libra //"天秤座"
	case (month == time.October && day >= 24) || (month == time.November && day <= 22):
		return amuser.Scorpio //"天蝎座"
	case (month == time.November && day >= 23) || (month == time.December && day <= 21):
		return amuser.Sagittarius //"射手座"
	case (month == time.December && day >= 22) || (month == time.January && day <= 19):
		return amuser.Capricorn //"摩羯座"
	case (month == time.January && day >= 20) || (month == time.February && day <= 18):
		return amuser.Aquarius //"水瓶座"
	default:
		return amuser.Pisces //"双鱼座"
	}
}
