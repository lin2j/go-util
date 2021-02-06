package util

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	BaseNumber     = "0123456789"
	BaseChar       = "abcdefghijklmnopqrstuvwxyz"
	BaseCharNumber = BaseChar + BaseNumber
)

// RandomInt 产生一个在 [min, max) 的随机数
func RandomInt(min, max int) int {
	if max < min {
		panic(errors.New(fmt.Sprintf("min must be less than or equal max, min is %d, max is %d", min, max)))
	}
	return min + rand.Intn(max-min)
}

// RandomIntN 产生一个在 [0. limit) 的随机整数
func RandomIntN(limit int) int {
	return RandomInt(0, limit)
}

//  RandomBool 返回随机布尔值
func RandomBool() bool {
	return RandomIntN(2) == 0
}

// RandomStringN 产生长度为 n 的随机字符串，n 需要大于 0
func RandomStringN(n int) string {
	return RandomStringBase(BaseCharNumber, n)
}

// RandomStringBase 以 baseStr 参数的内容为基础，拼接处长度为 length 的字符串
func RandomStringBase(baseStr string, length int) string {
	if length < 0 {
		panic(errors.New("len must not be negative"))
	}
	baseLen := len(baseStr)
	result := ""
	for i := 0; i < length; i++ {
		result += string(baseStr[RandomIntN(baseLen)])
	}
	return result
}

// RandomNumbers 返回一个随机字符串
func RandomNumbers(length int) string {
	return RandomStringBase(BaseNumber, length)
}

// RandomPhone 返回一个随机手机号
func RandomPhone() string {
	return TelHead[RandomIntN(len(TelHead))] + RandomNumbers(8)
}

// RandomEmail 返回一个随机邮箱
func RandomEmail() string {
	return RandomNumbers(8) + EmailSuffix[RandomIntN(len(EmailSuffix))]
}

// RandomChineseName 返回一个随机性别，随机长度的姓名
func RandomChineseName() string {
	// 是否为男性姓名
	male := RandomBool()
	// 姓名长度，一个字或者两个字（复姓算一个长度）
	nameLen := RandomInt(2, 4)
	return RandomChineseNameBool(male, nameLen)
}

// RandomChineseNameBool 返回指定长度、指定性别的名字
func RandomChineseNameBool(male bool, nameLen int) string {
	var nameStr string
	if male {
		nameStr = MaleName
	} else {
		nameStr = FemaleName
	}
	runes := []rune(nameStr)
	n := len(runes)
	firstName := FirstName[RandomIntN(len(FirstName))]
	for i := 0; i < nameLen; i++ {
		firstName += string(runes[RandomIntN(n)])
	}
	return firstName
}

// RandomDateTime 以当天为基准，随机产生一个日期
func RandomDateTime(min, max int) (time.Time, error) {
	return RandomDateTimeBase(time.Now(), DayField, min, max)
}

// RandomDateTimeBase 以 baseTime 为基准，对指定的时间字段进行偏移，产生一个随机时间
func RandomDateTimeBase(baseTime time.Time, field int, min, max int) (time.Time, error) {
	if t, err := OffsetTime(baseTime, field, RandomInt(min, max)); err != nil {
		return time.Time{}, err
	} else {
		return t, nil
	}
}
