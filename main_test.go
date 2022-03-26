package main

import "testing"

func TestDateConversion(t *testing.T) {
	testDate := Date{day: 12, month: 12, year: 2000}
	if ConvertInputToDate("12/12/2000") != testDate {
		t.Error("Date convert Function Not working")
	}
}
