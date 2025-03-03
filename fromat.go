package gjpd

import (
	"fmt"
	"strings"
	"time"
)

const (
	PrefixAbbr       = "g"
	PrefixShort      = "gg"
	PrefixLong       = "ggg"
	LayoutYearLong   = "ee"
	LayoutYearShort  = "e"
	LayoutMonthLong  = "mm"
	LayoutMonthShort = "m"
	LayoutDayLong    = "dd"
	LayoutDayShort   = "d"
)

func (g Gengo) Format(layout string) (s string) {
	if g.time == nil {
		return
	}
	s = strings.ToLower(layout)
	startedAt, _ := time.Parse(time.DateOnly, g.StartedAt)

	isLongPrefix := strings.HasPrefix(s, PrefixLong)
	isShortPrefix := strings.HasPrefix(s, PrefixShort) && !isLongPrefix
	isAbbrPrefix := strings.HasPrefix(s, PrefixAbbr) && !isShortPrefix && !isLongPrefix

	if isShortPrefix {
		s = strings.Replace(s, PrefixShort, g.Short, 1)
	} else if isAbbrPrefix {
		s = strings.Replace(s, PrefixAbbr, g.Abbr, 1)
	} else {
		s = strings.Replace(s, PrefixLong, g.Name, 1)
	}

	isLongWeekday := strings.Contains(s, LayoutWeekdayLong)
	isShortWeekday := strings.Contains(s, LayoutWeekdayShort) && !isLongWeekday
	isLongEnWeekday := strings.Contains(s, LayoutWeekdayEnLong)
	isShortEnWeekday := strings.Contains(s, LayoutWeekdayEnShort) && !isLongEnWeekday
	var sw string
	if isShortWeekday {
		sw = TimeToWeekday(*g.time, LayoutWeekdayShort)
		s = strings.Replace(s, LayoutWeekdayShort, sw, 1)
	} else if isLongEnWeekday {
		sw = TimeToWeekday(*g.time, LayoutWeekdayEnLong)
		s = strings.Replace(s, LayoutWeekdayEnLong, sw, 1)
	} else if isShortEnWeekday {
		sw = TimeToWeekday(*g.time, LayoutWeekdayEnShort)
		s = strings.Replace(s, LayoutWeekdayEnShort, sw, 1)
	} else {
		sw = TimeToWeekday(*g.time, LayoutWeekdayLong)
		s = strings.Replace(s, LayoutWeekdayLong, sw, 1)
	}

	isLongYear := strings.Contains(s, LayoutYearLong)
	isLongMonth := strings.Contains(s, LayoutMonthLong)
	isLongDay := strings.Contains(s, LayoutDayLong) && !isLongEnWeekday && !isShortEnWeekday

	y := g.time.Year() - startedAt.Year() + 1
	m := g.time.Month()
	d := g.time.Day()

	var sy string
	if isLongYear {
		sy = fmt.Sprintf("%02d", y)
		s = strings.Replace(s, LayoutYearLong, sy, 1)
	} else {
		sy = fmt.Sprintf("%d", y)
		s = strings.Replace(s, LayoutYearShort, sy, 1)
	}

	var sm string
	if isLongMonth {
		sm = fmt.Sprintf("%02d", m)
		s = strings.Replace(s, LayoutMonthLong, sm, 1)
	} else {
		sm = fmt.Sprintf("%d", m)
		s = strings.Replace(s, LayoutMonthShort, sm, 1)
	}

	var sd string
	if isLongDay {
		sd = fmt.Sprintf("%02d", d)
		s = strings.Replace(s, LayoutDayLong, sd, 1)
	} else {
		sd = fmt.Sprintf("%d", d)
		s = strings.Replace(s, LayoutDayShort, sd, 1)
	}

	return
}
