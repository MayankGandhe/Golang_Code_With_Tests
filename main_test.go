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
