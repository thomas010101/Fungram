package util

import (
	"strconv"
	"strings"
	"time"
)

const (
	DateLayout             = "2006-01-02 15:04:05"
	NotSeparatorDataLayout = "20060102150405"
)

func FormatDate(t time.Time, layout string) string {
	return t.Format(layout)
}

// StrToIntMonth 字符串月份转整数月份
func StrToIntMonth(month string) int {
	var data = map[string]int{
		"January":   0,
		"February":  1,
		"March":     2,
		"April":     3,
		"May":       4,
		"June":      5,
		"July":      6,
		"August":    7,
		"September": 8,
		"October":   9,
		"November":  10,
		"December":  11,
	}
	return data[month]
}

// GetTodayYMD 得到以sep为分隔符的年、月、日字符串(今天)
func GetTodayYMD(sep string) string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	year := now.Year()
	month := StrToIntMonth(now.Month().String())
	date := now.Day()

	var monthStr string
	var dateStr string
	if month < 9 {
		monthStr = "0" + strconv.Itoa(month+1)
	} else {
		monthStr = strconv.Itoa(month + 1)
	}

	if date < 10 {
		dateStr = "0" + strconv.Itoa(date)
	} else {
		dateStr = strconv.Itoa(date)
	}
	return strconv.Itoa(year) + sep + monthStr + sep + dateStr
}

// GetTodayYM 得到以sep为分隔符的年、月字符串(今天所属于的月份)
func GetTodayYM(sep string) string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	year := now.Year()
	month := StrToIntMonth(now.Month().String())

	var monthStr string
	if month < 9 {
		monthStr = "0" + strconv.Itoa(month+1)
	} else {
		monthStr = strconv.Itoa(month + 1)
	}
	return strconv.Itoa(year) + sep + monthStr
}

// GetYesterdayYMD 得到以sep为分隔符的年、月、日字符串(昨天)
func GetYesterdayYMD(sep string) string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	todaySec := today.Unix()            //秒
	yesterdaySec := todaySec - 24*60*60 //秒
	yesterdayTime := time.Unix(yesterdaySec, 0)
	yesterdayYMD := yesterdayTime.Format("2006-01-02")
	return strings.Replace(yesterdayYMD, "-", sep, -1)
}

// GetTomorrowYMD 得到以sep为分隔符的年、月、日字符串(明天)
func GetTomorrowYMD(sep string) string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	todaySec := today.Unix()           //秒
	tomorrowSec := todaySec + 24*60*60 //秒
	tomorrowTime := time.Unix(tomorrowSec, 0)
	tomorrowYMD := tomorrowTime.Format("2006-01-02")
	return strings.Replace(tomorrowYMD, "-", sep, -1)
}

// GetTodayTime 返回今天零点的time
func GetTodayTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	// now.Year(), now.Month(), now.Day() 是以本地时区为参照的年、月、日
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	return today
}

// GetYesterdayTime 返回昨天零点的time
func GetYesterdayTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	// now.Year(), now.Month(), now.Day() 是以本地时区为参照的年、月、日
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	yesterdaySec := today.Unix() - 24*60*60
	return time.Unix(yesterdaySec, 0)
}

// GetYesterdayTime 返回昨天零点的时间戳
func GetYesterdayStartUnix() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)

	// now.Year(), now.Month(), now.Day() 是以本地时区为参照的年、月、日
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	yesterdayUnix := today.Unix() - 24*60*60
	return yesterdayUnix
}

// GetYesterdayTime 返回昨天零点的时间戳
func GetYesterdayEndUnix() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)
	// now.Year(), now.Month(), now.Day() 是以本地时区为参照的年、月、日
	today := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, location)
	yesterdayUnix := today.Unix() - 24*60*60
	return yesterdayUnix
}

// 返回指定星期的第一天和最后一天
func GetFirstDayAndLastDayOfWeek(year, week int) (firstDay, lastDay time.Time) {
	location, _ := time.LoadLocation("Asia/Shanghai")

	firstDay = time.Date(year, 0, 0, 0, 0, 0, 0, location)
	isoYear, isoWeek := firstDay.ISOWeek()

	// iterate back to Monday
	for firstDay.Weekday() != time.Monday {
		firstDay = firstDay.AddDate(0, 0, -1)
		isoYear, isoWeek = firstDay.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		firstDay = firstDay.AddDate(0, 0, 7)
		isoYear, isoWeek = firstDay.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		firstDay = firstDay.AddDate(0, 0, 7)
		isoYear, isoWeek = firstDay.ISOWeek()
	}

	lastDay = time.Date(firstDay.Year(), firstDay.Month(), firstDay.Day()+6, 23, 59, 59, 0, location)

	return
}

func ToTimeString(i int64) string {
	if i > 1e12 {
		i = i / 1000
	}
	return time.Unix(i, 0).Format(DateLayout)
}
