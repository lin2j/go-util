package util

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(fmt.Sprintf("%-30s", NormDatePattern), now.Format(NormDatePattern))
	fmt.Println(fmt.Sprintf("%-30s", NormTimePattern), now.Format(NormTimePattern))
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimeMinutePattern), now.Format(NormDateTimeMinutePattern))
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), now.Format(NormDateTimePattern))
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimeMsPattern), now.Format(NormDateTimeMsPattern))
	fmt.Println(fmt.Sprintf("%-30s", PureDatePattern), now.Format(PureDatePattern))
	fmt.Println(fmt.Sprintf("%-30s", PureTimePattern), now.Format(PureTimePattern))
	fmt.Println(fmt.Sprintf("%-30s", PureDateTimePattern), now.Format(PureDateTimePattern))
	fmt.Println(fmt.Sprintf("%-30s", PureDateTimeMsPattern), now.Format(PureDateTimeMsPattern))
	fmt.Println(fmt.Sprintf("%-30s", ChineseDatePattern), now.Format(ChineseDatePattern))
	fmt.Println(fmt.Sprintf("%-30s", DayHourPattern), now.Format(DayHourPattern))
	fmt.Println(fmt.Sprintf("%-30s", MonthDayPattern), now.Format(MonthDayPattern))
}

func TestTimeOp(t *testing.T) {
	now := time.Now()
	offset := OffsetDays(now, -3)
	hour := OffsetHours(now, 96)

	fmt.Println(fmt.Sprintf("%-30s", PureHourPattern), now.Format(PureHourPattern))
	fmt.Println(fmt.Sprintf("%-30s", PureHourPattern), offset.Format(PureHourPattern))
	fmt.Println(fmt.Sprintf("%-30s", PureHourPattern), hour.Format(PureHourPattern))

	newT, _ := OffsetTime(now, YearField, 1)
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), newT.Format(NormDateTimePattern))
	newT, _ = OffsetTime(now, MonthField, 1)
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), newT.Format(NormDateTimePattern))
	newT, _ = OffsetTime(now, DayField, 1)
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), newT.Format(NormDateTimePattern))
	newT, _ = OffsetTime(now, HourField, 1)
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), newT.Format(NormDateTimePattern))
	newT, _ = OffsetTime(now, MinuteField, 1)
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), newT.Format(NormDateTimePattern))
	newT, _ = OffsetTime(now, SecondField, 1)
	fmt.Println(fmt.Sprintf("%-30s", NormDateTimePattern), newT.Format(NormDateTimePattern))
}

func TestToday(t *testing.T) {
	fmt.Println(Now())
	fmt.Println(Today())
	fmt.Println(EndOfToday().Format(NormDateTimePattern))
	fmt.Println(BeginOfToday().Format(NormDateTimePattern))
	fmt.Println("the hour of today", time.Now().Hour())
	fmt.Println("pm", IsPM())
	fmt.Println("am", IsAM())
	fmt.Println("is today weekend?", IsWeekend())
	fmt.Println("is 2008 leap year?", IsLeap(2008))
	fmt.Println("is 2009 leap year?", IsLeap(2009))
	age, err := AgeStr("1997-04-07", "2021-02-04")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("birthday: 1997-04-07, date to compare: 2020-02-04", age)
	}
	age, err = AgeOfNowStr("1997-04-07")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("birthday: 1997-04-07", age)
	}
}
