package gjpd

import (
	"log"
	"testing"
	"time"
)

func TestGetGengo(t *testing.T) {
	g := GetGengo("heisei")
	expected := "平成"
	actual := g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	g = GetGengo("reiwa")
	expected = "令和"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}

func TestGetGengoByTime(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2019-05-01")
	g, err := GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected := "令和"
	actual := g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "2019-04-30")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "平成"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1989-01-08")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "平成"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1989-01-07")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "昭和"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1926-12-25")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "昭和"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1926-12-24")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "大正"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1912-07-30")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "大正"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1912-07-29")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "明治"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	date, _ = time.Parse(time.DateOnly, "1868-10-23")
	g, err = GetGengoByTime(date)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected = "明治"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}
func TestGetGengoByKey(t *testing.T) {

	g := GetGengoByKey("heisei")
	expected := "平成"
	actual := g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	g = GetGengoByKey("showa")
	expected = "昭和"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}
func TestGetGengoByAbbr(t *testing.T) {

	g := GetGengoByAbbr("H")
	expected := "平成"
	actual := g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	g = GetGengoByAbbr("s")
	expected = "昭和"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}

func TestGetGengoByName(t *testing.T) {

	g := GetGengoByName("平成")
	expected := "平成"
	actual := g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	g = GetGengoByName("")
	expected = "令和"
	actual = g.Name
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}

func TestSetDate(t *testing.T) {
	g := GetGengoByName("令和")
	t1, err := g.SetDate("7", "1", "1")
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	t2, _ := time.Parse(time.DateOnly, "2025-01-01")
	if t1 != t2 {
		t.Errorf("got: actual: %v, expected: %v", t1, t2)
		log.Printf("g: %+v", g)
	}

	t1, err = g.SetDate("1", "1", "1")
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	t2, _ = time.Parse(time.DateOnly, "2019-01-01")
	if t1 != t2 {
		t.Errorf("got: actual: %v, expected: %v", t1, t2)
	}
}

func TestYears(t *testing.T) {
	g := GetGengoByName("昭和")
	actual := g.Years()
	expected := 64
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	g = GetGengoByName("平成")
	actual = g.Years()
	expected = 31
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
