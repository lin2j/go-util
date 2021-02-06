package util

import (
	"errors"
	"fmt"
	"time"
)

const (
	Second = "05"
	Minute = "04"
	Hour   = "15"
	Day    = "02"
	Month  = "01"
	Year   = "2006"
	Milli  = "000"
	Micro  = "000000"
	Nano   = "000000000"

	NormDatePattern           = "2006-01-02"              // yyyy-MM-dd
	NormTimePattern           = "15:04:05"                // HH:mm:ss
	NormDateTimeMinutePattern = "2006-01-02 15:04"        // yyyy-MM-dd HH:mm
	NormDateTimePattern       = "2006-01-02 15:04:05"     // yyyy-MM-dd HH:mm:ss
	NormDateTimeMsPattern     = "2006-01-02 15:04:05.000" // yyyy-MM-dd HH.mm.ss.SSS
	PureDatePattern           = "20060102"                // yyyyMMdd
	PureTimePattern           = "150405"                  // HHmmss
	PureHourPattern           = "2006010215"              // yyyyMMddHH
	PureDateTimePattern       = "20060102150405"          // yyyyMMddHHmmss
	PureDateTimeMsPattern     = "20060102150405000"       // yyyyMMddHHmmssSSS
	ChineseDatePattern        = "2006年01月02日"             // yyyy年MM月dd日
	DayHourPattern            = "02-15"                   // dd-HH
	MonthDayPattern           = "01-02"                   // mm-dd

	StrNormDatePattern           = "yyyy-MM-dd"
	StrNormTimePattern           = "HH:mm:ss"
	StrNormDateTimeMinutePattern = "yyyy-MM-dd HH:mm"
	StrNormDateTimePattern       = "yyyy-MM-dd HH:mm:ss"
	StrNormDateTimeMsPattern     = "yyyy-MM-dd HH.mm.ss.SSS"
	StrPureDatePattern           = "yyyyMMdd"
	StrPureTimePattern           = "HHmmss"
	StrPureHourPattern           = "yyyyMMddHH"
	StrPureDateTimePattern       = "yyyyMMddHHmmss"
	StrPureDateTimeMsPattern     = "yyyyMMddHHmmssSSS"
	StrChineseDatePattern        = "yyyy年MM月dd日"
)

const (
	YearField = iota
	MonthField
	DayField
	HourField
	MinuteField
	SecondField
)

var PatternMap = map[string]string{
	StrNormDatePattern:           NormDatePattern,
	StrNormTimePattern:           NormTimePattern,
	StrNormDateTimeMinutePattern: NormDateTimeMinutePattern,
	StrNormDateTimePattern:       NormDateTimePattern,
	StrNormDateTimeMsPattern:     NormDateTimeMsPattern,
	StrPureDatePattern:           PureDatePattern,
	StrPureTimePattern:           PureTimePattern,
	StrPureHourPattern:           PureHourPattern,
	StrPureDateTimePattern:       PureDateTimePattern,
	StrPureDateTimeMsPattern:     PureDateTimeMsPattern,
	StrChineseDatePattern:        ChineseDatePattern,
}

// 偏移时间的函数
type TimeFunc func(t time.Time, offset int) time.Time

// 时间字段和时间偏移函数的对应关系
var TimeFuncMap = map[int]TimeFunc{
	YearField:   OffsetYears,
	MonthField:  OffsetMonths,
	DayField:    OffsetDays,
	HourField:   OffsetHours,
	MinuteField: OffsetMinutes,
	SecondField: OffsetSeconds,
}

// GetPattern get golang time format pattern
func GetPattern(origin string) string {
	if p, ok := PatternMap[origin]; ok {
		return p
	}
	return NormDateTimePattern
}

func OffsetTime(t time.Time, field, offset int) (time.Time, error) {
	if f, ok := TimeFuncMap[field]; !ok {
		return time.Time{}, errors.New("unknown time field")
	} else {
		return f(t, offset), nil
	}
}

func OffsetYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}
func OffsetMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}
func OffsetDays(t time.Time, Days int) time.Time {
	return t.AddDate(0, 0, Days)
}
func OffsetHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}
func OffsetMinutes(t time.Time, min int) time.Time {
	return t.Add(time.Duration(min) * time.Minute)
}
func OffsetSeconds(t time.Time, sec int) time.Time {
	return t.Add(time.Duration(sec) * time.Second)
}

// BeginOfToday returns the time like yyyy-MM-dd 00:00:00
// yyyy-MM-dd is the date of today
func BeginOfToday() time.Time {
	beginOfToday, _ := time.ParseInLocation(NormDatePattern, Today(), time.Local)
	return beginOfToday
}

// EndOfToday 返回一个形如 yyyy-MM-dd 23:59:59 的字符串
// yyyy-MM-dd 是今天的日期
func EndOfToday() time.Time {
	tomorrow := time.Now().AddDate(0, 0, 1)
	tomorrowStr := tomorrow.Format(NormDatePattern)
	beginOfTomorrow, _ := time.ParseInLocation(NormDatePattern, tomorrowStr, time.Local)
	return beginOfTomorrow.Add(-1 * time.Second)
}

// Now 返回当前时间，格式为 yyyy-MM-dd HH:mm:ss
func Now() string {
	return time.Now().Format(NormDateTimePattern)
}

// Today 返回当前日期，格式为 yyyy-MM-dd
func Today() string {
	return time.Now().Format(NormDatePattern)
}

// 判断是否是下午
func IsPM() bool {
	h := time.Now().In(time.Local).Hour()
	return h >= 12
}

// 判断是否是上午
func IsAM() bool {
	return !IsPM()
}

// IsWeekend 判断今天是否是周末，周末指周六和周日
func IsWeekend() bool {
	dayOfWeek := time.Now().Weekday()
	return dayOfWeek == time.Saturday || dayOfWeek == time.Sunday
}

// IsLeap 判断给定的年份是否是闰年
func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// AgeStr 返回生日到某一个日期的法定年龄
func AgeStr(birthdayStr, dateToCompare string) (int, error) {
	birthT, err := time.ParseInLocation(NormDatePattern, birthdayStr, time.Local)
	if err != nil {
		return 0, err
	}
	compareT, err := time.ParseInLocation(NormDatePattern, dateToCompare, time.Local)
	if err != nil {
		return 0, err
	}
	return Age(birthT, compareT)
}

// Age 返回生日到某一个日期的法定年龄
func Age(birthday, dateToCompare time.Time) (int, error) {
	if dateToCompare.Before(birthday) {
		date := dateToCompare.Format(NormDatePattern)
		return 0, errors.New(fmt.Sprintf("birthday is after the date %s", date))
	}
	age := dateToCompare.Year() - birthday.Year()
	if dateToCompare.Month() == birthday.Month() {
		if birthday.Day() < dateToCompare.Day() {
			age--
		}
	} else if dateToCompare.Month() < birthday.Month() {
		age--
	}
	return age, nil
}

// AgeOfNowStr 返回生日距今的法定年龄
func AgeOfNowStr(birthdayStr string) (int, error) {
	birthT, err := time.ParseInLocation(NormDatePattern, birthdayStr, time.Local)
	if err != nil {
		return 0, err
	}
	return AgeOfNow(birthT)
}

// AgeOfNow 返回生日距今的法定年龄
func AgeOfNow(birthday time.Time) (int, error) {
	now := time.Now()
	return Age(birthday, now)
}
