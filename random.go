package util

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	BaseNumber     = "0123456789"
	BaseChar       = "abcdefghijklmnopqrstuvwxyz"
	BaseCharNumber = BaseChar + BaseNumber
)

// RandomInt returns a random integer in [min, max)
func RandomInt(min, max int) int {
	if max < min {
		panic(errors.New(fmt.Sprintf("min must be less than or equal max, min is %d, max is %d", min, max)))
	}
	return min + rand.Intn(max-min)
}

// RandomIntN returns a random integer in [0. limit)
func RandomIntN(limit int) int {
	return RandomInt(0, limit)
}

//  RandomBool returns random bool
func RandomBool() bool {
	return RandomIntN(2) == 0
}

// RandomStringN returns string which length is len
func RandomStringN(len int) string {
	return RandomStringBase(BaseCharNumber, len)
}

// RandomStringBase returns string which content base on baseStr argument
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

// RandomNumbers returns a number string
func RandomNumbers(length int) string {
	return RandomStringBase(BaseNumber, length)
}
