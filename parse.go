package gjpd

import (
	"errors"
	"regexp"
	"slices"
	"strings"
	"time"
)

const (
	LayoutSlashShort = "2006/1/2"
	LayoutSlash      = "2006/01/02"
)

var (
	RegexpGengAbbr  = regexp.MustCompile(`^([\S]{1})([0-9]{1,2})\.([0-9]{1,2})\.([0-9]{1,2})`)
	RegexpGengoFull = regexp.MustCompile(`^([\S]{2})([0-9]{1,2})年([0-9]{1,2})月([0-9]{1,2})日`)
)

func Trim(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "/0", "/", -1)
	str = strings.Replace(str, PrefixFirstYear, "1", 1)
	str = ZenkakuToNumber(str)
	str = KanjiToNumber(str)

	return str
}

func Parse(str string) (t time.Time, e error) {
	if len(str) < 1 {
		e = errors.New(ErrorInvalidArg)
		return
	}
	str = Trim(str)

	t, e = time.Parse(time.DateOnly, str)
	if e == nil {
		return
	}

	t, e = time.Parse(LayoutSlashShort, str)
	if e == nil {
		return
	}
	e = nil

	firstString := string(str[0])

	if slices.Contains(GetAbbrs(), firstString) {
		return ParseAsAbbr(str)
	}

	return ParseAsName(str)
}

func ParseAsAbbr(str string) (t time.Time, e error) {
	str = Trim(str)

	matches := RegexpGengAbbr.FindAllStringSubmatch(str, -1)
	if len(matches) < 1 {
		e = errors.New(ErrorAbbrMatches)
		return
	}
	match := matches[0]
	if len(match) < 5 {
		e = errors.New(ErrorAbbrMatch)
		return
	}
	abbr := match[1]
	gengo := GetGengoByAbbr(abbr)

	sy := match[2]
	sm := match[3]
	sd := match[4]

	t, e = gengo.SetDate(sy, sm, sd)
	return
}

func ParseAsName(str string) (t time.Time, e error) {
	str = Trim(str)

	matches := RegexpGengoFull.FindAllStringSubmatch(str, -1)
	if len(matches) < 1 {
		e = errors.New(ErrorNameMatches)
		return
	}

	match := matches[0]
	if len(match) < 3 {
		e = errors.New(ErrorNameMatches)
		return
	}
	name := match[1]
	gengo := GetGengoByName(name)

	sy := match[2]
	sm := match[3]
	sd := match[4]

	t, e = gengo.SetDate(sy, sm, sd)
	return
}
