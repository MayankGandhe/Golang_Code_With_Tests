package main

import "testing"

func TestDateConversion(t *testing.T) {
	testDate := Date{day: 12, month: 12, year: 2000}
	if ConvertInputToDate("12/12/2000") != testDate {
		t.Error("Date convert Function Not working")
	}
}
func TestDateConversionInvalidDate(t *testing.T) {
	testDate := Date{}
	if ConvertInputToDate("1a/12/2000") != testDate {
		t.Error("Date convert Function Not working")
	}
}
func TestIsLeapYear(t *testing.T) {
	if !IsLeapYear(2000) {
		t.Error("Leap year validator not working")
	}
	if IsLeapYear(1900) {
		t.Error("Leap year validator not working")
	}
	if !IsLeapYear(1920) {
		t.Error("Leap year validator not working")
	}
	if IsLeapYear(1921) {
		t.Error("Leap year validator not working")
	}
}
func TestNumberOfDaysInMonth(t *testing.T) {
	if NumberOfDaysInMonth(1, 2000) != 31 {
		t.Error("Not able to calculate correct number of days in a year")
	}
	if NumberOfDaysInMonth(2, 2000) != 29 {
		t.Error("Not able to calculate correct number of days in a year")
	}
	if NumberOfDaysInMonth(2, 1900) != 28 {
		t.Error("Not able to calculate correct number of days in a year")
	}
	if NumberOfDaysInMonth(4, 1900) != 30 {
		t.Error("Not able to calculate correct number of days in a year")
	}
}

func TestDateValidation(t *testing.T) {
	if ValidateDate(Date{12, 12, 12}) {
		t.Error("Date validate Function Not working")
	}
	if !ValidateDate(Date{12, 12, 2012}) {
		t.Error("Date validate Function Not working")
	}
	if !ValidateDate(Date{29, 2, 2012}) {
		t.Error("Date validate Function Not working")
	}
	if !ValidateDate(Date{29, 2, 2000}) {
		t.Error("Date validate Function Not working")
	}
	if ValidateDate(Date{29, 2, 2001}) {
		t.Error("Date validate Function Not working")
	}

}
func TestFunctionToGetDaysFromStart(t *testing.T) {
	if CalculateNumberOfDaysFromStart(Date{1, 1, 1}) != 1 {
		t.Error("Function Not Working")
	}
	numberOfDays := CalculateNumberOfDaysFromStart(Date{1, 1, 2001})
	if numberOfDays != 1 {
		t.Error("Function Not Working {}", numberOfDays)
	}
}
