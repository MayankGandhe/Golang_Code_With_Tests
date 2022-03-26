package main

import (
	"fmt"
	"math"
	"os"
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
func NumberOfDaysInMonth(month, year int) int {
	switch month {
	case 1,
		3,
		5,
		7,
		8,
		10,
		12:
		return 31
	case 4,
		6,
		9,
		11:
		return 30
	}
	if month == 2 {
		if IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	}
	return 0
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
func ValidateDate(date Date) bool {
	if InBetween(date.year, 1900, 2999) {
		if InBetween(date.month, 1, 12) {
			if InBetween(date.day, 1, NumberOfDaysInMonth(date.month, date.year)) {
				return true
			}
		}
	}
	return false
}
func CalculateNumberOfDaysFromStart(date Date) int {
	NumberOfdaysFromYear := (date.year - 1) * 365
	var NumberOfdaysFromMonth int
	for month := 1; month < date.month; month++ {
		NumberOfdaysFromMonth += NumberOfDaysInMonth(month, date.year)
	}
	return NumberOfdaysFromMonth + NumberOfdaysFromYear + date.day + CalculateNumberOfLeapYears(date.year-1)
}
func CalculateDifference(date1, date2 Date) int {
	numberOfDays := int(math.Abs(float64(CalculateNumberOfDaysFromStart(date1)) - float64(CalculateNumberOfDaysFromStart(date2))))
	if numberOfDays == 1 {
		return 0
	}
	return numberOfDays - 1
}
func main() {
	arguments := os.Args[1:]
	if arguments[0] == "help" {

		fmt.Println("This Program will help you find the number of days between two dates")
		fmt.Println("Enter both dates in DD/MM/YYYY format and as an argument to this binary")
		fmt.Println("For example: ./calculateDays 22/12/1999 4/4/2004")
		fmt.Println("./main 22/12/1999 4/4/2004")
		fmt.Println("#########IMPORTANT#######")
		fmt.Println("This programm will only Support dates from 1/1/1900 to 31/12/2999")

		os.Exit(0)
	}
	dateA := arguments[0]
	date1 := ConvertInputToDate(dateA)
	if !ValidateDate(date1) {
		fmt.Println("First date entered is Not Valid")
		os.Exit(1)
	}
	dateB := arguments[1]

	date2 := ConvertInputToDate(dateB)
	if !ValidateDate(date2) {
		fmt.Println("Second date entered is Not Valid")
		os.Exit(1)
	}
	fmt.Println(CalculateDifference(date1, date2))

}
