package gjpd

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	d, _ := time.Parse(time.DateOnly, "2025-01-01")
	g, err := GetGengoByTime(d)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}

	actual := g.Format("ge.m.d")
	expected := "R7.1.1"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggee年m月d日")
	expected = "令07年1月1日"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggee年mm月dd日")
	expected = "令07年01月01日"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggge年m月d日")
	expected = "令和7年1月1日"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggge年m月d日(ddd)")
	expected = "令和7年1月1日(Wed)"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggge年m月d日(dddd)")
	expected = "令和7年1月1日(Wednesday)"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggge年m月d日(aaa)")
	expected = "令和7年1月1日(水)"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = g.Format("ggge年m月d日(aaaa)")
	expected = "令和7年1月1日(水曜日)"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}
