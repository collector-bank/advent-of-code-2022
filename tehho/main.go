package main

import (
	"fmt"
	"os"
)

func GetDayData(day string) string {
	data, err := os.ReadFile("data/Day" + day + ".data")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	day01data := GetDayData("01")
	day01ans := Day01part01(day01data)
	fmt.Printf("Day01 part01: %d\n", day01ans)
	day01ans = Day01part02(day01data)
	fmt.Printf("Day01 part02: %d\n", day01ans)
	day02data := GetDayData("02")
	day02ans := Day02part01(day02data)
	fmt.Printf("Day02 part01: %d\n", day02ans)
	day02ans = Day02part02(day02data)
	fmt.Printf("Day02 part02: %d\n", day02ans)

}
