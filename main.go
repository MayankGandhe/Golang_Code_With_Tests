package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateNumberOfLeapYears(year int) int {
	leapYears := (year/4 + year/400) - year/100
	return leapYears
}

type Date struct {
	day   int
	month int
	year  int
}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
func ConvertInputToDate(givenDate string) Date {
	splitDate := strings.Split(givenDate, "/")
	var day int
	var month int
	var year int
	var err error
	if day, err = strconv.Atoi(splitDate[0]); err != nil {
		fmt.Printf("Day= %d, type: %T\nShould Be a valid Integer\n", day, day)
		return Date{}
	}
	if month, err = strconv.Atoi(splitDate[1]); err != nil {
		fmt.Printf("Month= %d , type: %T \nShould Be a valid Integer\n", month, month)
		return Date{}

	}
	if year, err = strconv.Atoi(splitDate[2]); err != nil {
		fmt.Printf("Year= %d, type: %T\nShould Be a valid Integer\n", year, year)
		return Date{}

	}
	return Date{day: day, month: month, year: year}
}

func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}
func CalculateDifference(date1, date2 Date) int {
	return 1
}
func main() {
	fmt.Println("This Program will help you find the number of days between two dates")
	fmt.Println("Enter First Date in DD/MM/YYYY format")
	var dateA string
	fmt.Scanln(&dateA)
	date1 := ConvertInputToDate(dateA)
	fmt.Println("Enter Second Date in DD/MM/YYYY format")
	var dateB string
	fmt.Scanln(&dateB)
	date2 := ConvertInputToDate(dateA)
	CalculateDifference(date1, date2)

}
