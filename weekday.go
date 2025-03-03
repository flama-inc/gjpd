package gjpd

import "time"

const (
	LayoutWeekdayLong    = "aaaa"
	LayoutWeekdayShort   = "aaa"
	LayoutWeekdayEnLong  = "dddd"
	LayoutWeekdayEnShort = "ddd"
)

var (
	WeekdaysLong = []string{
		"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日",
	}
	WeekdaysShort = []string{
		"日", "月", "火", "水", "木", "金", "土",
	}
)

func GetShortWeekdays() []string {
	return WeekdaysShort
}

func GetLongWeekdays() []string {
	return WeekdaysLong
}

func TimeToWeekday(t time.Time, layout string) (s string) {
	w := t.Weekday()
	switch layout {
	case LayoutWeekdayEnLong:
		s = w.String()
	case LayoutWeekdayEnShort:
		s = w.String()[0:3]
	case LayoutWeekdayShort:
		s = WeekdaysShort[w]
	default:
		s = WeekdaysLong[w]
	}
	return
}
