package gjpd

import (
	"testing"
	"time"
)

func TestStringNextWeekday(t *testing.T) {

	date := "2024-02-10"

	actual, err := StringNextWeekday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ := time.Parse(time.DateOnly, "2024-02-13")
	if actual.Compare(expected) != 0 {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2024-02-23"

	actual, err = StringNextWeekday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2024-02-26")
	if actual.Compare(expected) != 0 {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2024-02-09"

	actual, err = StringNextWeekday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2024-02-09")
	if actual.Compare(expected) != 0 {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}

func TestIsStringHolyday(t *testing.T) {
	date := "2024-02-11"

	actual, err := IsStringHolyday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected := true
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2024-02-13"

	actual, err = IsStringHolyday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = false
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2024-02-17"

	actual, err = IsStringHolyday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = true
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2024-02-18"

	actual, err = IsStringHolyday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = true
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}

func TestIsStringNationalHolyday(t *testing.T) {
	date := "2024-02-11"

	actual, err := IsStringNationalHolyday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected := true
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2024-01-02"

	actual, err = IsStringNationalHolyday(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = false
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
