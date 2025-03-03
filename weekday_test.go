package gjpd

import (
	"testing"
	"time"
)

func TestTimeToWeekday(t *testing.T) {

	d, _ := time.Parse(time.DateOnly, "2025-03-03")
	actual := TimeToWeekday(d, "aaaa")
	expected := "月曜日"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = TimeToWeekday(d, "aaa")
	expected = "月"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = TimeToWeekday(d, "dddd")
	expected = "Monday"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

	actual = TimeToWeekday(d, "ddd")
	expected = "Mon"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}
}
