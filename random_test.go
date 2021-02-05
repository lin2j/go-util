package util

import (
	"fmt"
	"testing"
)

func TestRandomInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := RandomInt(0, 10)
		if n < 0 || n > 10 {
			fmt.Println(n)
			break
		}
	}
	fmt.Println("end")
}

func TestRandomIntn(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := RandomIntN(10)
		if n < 0 || n > 10 {
			fmt.Println(n)
			break
		}
	}
	fmt.Println("end")
}

func TestRandomBool(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := RandomBool()
		fmt.Println(n)
	}
	fmt.Println("end")
}

func TestRandomStringN(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomStringN(10))
	}
}

func TestRandomNumbers(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomNumbers(5))
	}
}

func TestRandomPhone(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomPhone())
	}
}

func TestRandomEmail(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomEmail())
	}
}

func TestRandomChineseName(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(RandomChineseName())
	}
}
