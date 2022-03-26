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

func ConvertInputToDate(givenDate string) Date {
	splitDate := strings.Split(givenDate, "/")
	var day int
	var month int
	var year int
	var err error
	if day, err = strconv.Atoi(splitDate[0]); err != nil {
		fmt.Printf("Day= %d, type: %T\n Should Be a valid Integer", day, day)
	}
	if month, err = strconv.Atoi(splitDate[1]); err != nil {
		fmt.Printf("Month= %d , type: %T \n Should Be a valid Integer", month, month)
	}
	if year, err = strconv.Atoi(splitDate[2]); err != nil {
		fmt.Printf("Year= %d, type: %T\n Should Be a valid Integer", year, year)
	}
	return Date{day: day, month: month, year: year}
}

// func CheckDateValidity(date date) bool {

// }
func main() {
	fmt.Println("This Program will help you find the number of days between two dates")
	fmt.Println("Enter First Date in DD/MM/YYYY format")
	var dateA string
	fmt.Scanln(&dateA)
	fmt.Println("Enter Second Date in DD/MM/YYYY format")
	var dateB string
	fmt.Scanln(&dateB)
	//print(ConvertInputToDate(dateA))
}
