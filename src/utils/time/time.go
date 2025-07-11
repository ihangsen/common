package time

import (
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

// EndWeek 本周结束时间(最后一天的起始时间)
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

// StartMonth 本月起始时间
func StartMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// EndMonth 本月结束时间(最后一天的起始时间)
func EndMonth(now time.Time) time.Time {
	// 获取下个月的第一天，然后减去一纳秒，得到本月最后一天
	nextMonth := now.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, nextMonth.Location())
	return firstDayOfNextMonth.AddDate(0, 0, -1)
}

// StartYear 本年起始时间
func StartYear(now time.Time) time.Time {
	return time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
}

// EndYear 本年结束时间(最后一天的起始时间)
func EndYear(now time.Time) time.Time {
	// 获取下一年的第一天，然后减去一天，得到本年最后一天
	nextYear := now.AddDate(1, 0, 0)
	firstDayOfNextYear := time.Date(nextYear.Year(), time.January, 1, 0, 0, 0, 0, nextYear.Location())
	return firstDayOfNextYear.AddDate(0, 0, -1)
}
