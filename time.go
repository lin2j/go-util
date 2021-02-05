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

// GetPattern get golang time format pattern
func GetPattern(origin string) string {
	if p, ok := PatternMap[origin]; ok {
		return p
	}
	return NormDateTimePattern
}

func OffsetYears(t *time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}
func OffsetMonth(t *time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}
func OffsetDays(t *time.Time, Days int) time.Time {
	return t.AddDate(0, 0, Days)
}
func OffsetHours(t *time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}
func OffsetMinute(t *time.Time, min int) time.Time {
	return t.Add(time.Duration(min) * time.Minute)
}
func OffsetSecond(t *time.Time, sec int) time.Time {
	return t.Add(time.Duration(sec) * time.Second)
}

// BeginOfToday returns the time like yyyy-MM-dd 00:00:00
// yyyy-MM-dd is the date of today
func BeginOfToday() time.Time {
	beginOfToday, _ := time.ParseInLocation(NormDatePattern, Today(), time.Local)
	return beginOfToday
}

// EndOfToday returns the time like yyyy-MM-dd 23:59:59
// yyyy-MM-dd is the date of today
func EndOfToday() time.Time {
	tomorrow := time.Now().AddDate(0, 0, 1)
	tomorrowStr := tomorrow.Format(NormDatePattern)
	beginOfTomorrow, _ := time.ParseInLocation(NormDatePattern, tomorrowStr, time.Local)
	return beginOfTomorrow.Add(-1 * time.Second)
}

// Now returns the time string that format is yyyy-MM-dd HH:mm:ss
func Now() string {
	return time.Now().Format(NormDateTimePattern)
}

// Today returns the time string that format is yyyy-MM-dd
func Today() string {
	return time.Now().Format(NormDatePattern)
}

// IsPM
func IsPM() bool {
	h := time.Now().In(time.Local).Hour()
	return h >= 12
}

// IsAM
func IsAM() bool {
	return !IsPM()
}

// IsWeekend determines if the day of week is weekend
func IsWeekend() bool {
	dayOfWeek := time.Now().Weekday()
	return dayOfWeek == time.Saturday || dayOfWeek == time.Sunday
}

// IsLeap determines if the given year is a leap year
func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// AgeStr returns the legal age between birthday and the date to compare
func AgeStr(birthdayStr, dateToCompare string) (int, error) {
	birthT, err := time.ParseInLocation(NormDatePattern, birthdayStr, time.Local)
	if err != nil {
		return 0, err
	}
	compareT, err := time.ParseInLocation(NormDatePattern, dateToCompare, time.Local)
	if err != nil {
		return 0, err
	}
	return Age(&birthT, &compareT)
}

// Age returns the legal age between birthday and the date to compare
func Age(birthday, dateToCompare *time.Time) (int, error) {
	if dateToCompare.Before(*birthday) {
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

// AgeOfNowStr returns the legal age from birthday to today
func AgeOfNowStr(birthdayStr string) (int, error) {
	birthT, err := time.ParseInLocation(NormDatePattern, birthdayStr, time.Local)
	if err != nil {
		return 0, err
	}
	return AgeOfNow(&birthT)
}

// AgeOfNow returns the legal age from birthday to today
func AgeOfNow(birthday *time.Time) (int, error) {
	now := time.Now()
	return Age(birthday, &now)
}
