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

}
