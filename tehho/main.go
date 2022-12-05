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

	day03data := GetDayData("03")
	day03ans := Day03part01(day03data)
	fmt.Printf("Day03 part01: %d\n", day03ans)
	day03ans = Day03part02(day03data)
	fmt.Printf("Day03 part02: %d\n", day03ans)

	day04data := GetDayData("04")
	day04ans := Day04part01(day04data)
	fmt.Printf("Day04 part01: %d\n", day04ans)
	day04ans = Day04part02(day04data)
	fmt.Printf("Day04 part02: %d\n", day04ans)

}
