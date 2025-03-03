package gjpd

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {

	date := "2025-01-01"
	actual, err := Parse(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ := time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2025/01/01"
	actual, err = Parse(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "2025/1/1"
	actual, err = Parse(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "R7.1.1"
	actual, err = Parse(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "令和7年1月1日"
	actual, err = Parse(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
func TestParseAsAbbr(t *testing.T) {

	date := "R7.1.1"
	actual, err := ParseAsAbbr(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ := time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
func TestParseAsName(t *testing.T) {

	date := "令和7年1月1日"
	actual, err := ParseAsName(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ := time.Parse(time.DateOnly, "2025-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "平成元年1月1日"
	actual, err = ParseAsName(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "1989-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "昭和64年1月1日"
	actual, err = ParseAsName(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "1989-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date = "令和元年1月1日"
	actual, err = ParseAsName(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected, _ = time.Parse(time.DateOnly, "2019-01-01")
	if actual.Year() != expected.Year() {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
