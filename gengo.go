package gjpd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	PrefixFirstYear = "元"
)

type Gengo struct {
	StartedAt string
	EndedAt   string
	Name      string
	Short     string
	Key       string
	Abbr      string
	//
	time *time.Time
}

var (
	gengoMeiji = Gengo{
		StartedAt: "1868-10-23",
		EndedAt:   "1912-07-29",
		Name:      "明治",
		Short:     "明",
		Key:       "Meiji",
		Abbr:      "M",
	}
	gengoTaisho = Gengo{
		StartedAt: "1912-07-30",
		EndedAt:   "1926-12-24",
		Name:      "大正",
		Short:     "大",
		Key:       "Taisho",
		Abbr:      "T",
	}
	gengoShowa = Gengo{
		StartedAt: "1926-12-25",
		EndedAt:   "1989-01-07",
		Name:      "昭和",
		Short:     "昭",
		Key:       "Showa",
		Abbr:      "S",
	}
	gengoHeisei = Gengo{
		StartedAt: "1989-01-08",
		EndedAt:   "2019-04-30",
		Name:      "平成",
		Short:     "平",
		Key:       "Heisei",
		Abbr:      "H",
	}
	gengoReiwa = Gengo{
		StartedAt: "2019-05-01",
		Name:      "令和",
		Short:     "令",
		Key:       "Reiwa",
		Abbr:      "R",
	}
	defaultGengo = gengoReiwa
)

func GetGengo(key string) Gengo {
	m := GetGengos()
	if v, ok := m[key]; ok {
		return v
	}
	return defaultGengo
}

func GetGengos() (m map[string]Gengo) {
	m = make(map[string]Gengo)

	m["meiji"] = gengoMeiji
	m["taisho"] = gengoTaisho
	m["showa"] = gengoShowa
	m["heisei"] = gengoHeisei
	m["reiwa"] = gengoReiwa

	return
}

func GetAbbrs() (abbrs []string) {
	for _, g := range GetGengos() {
		abbrs = append(abbrs, g.Abbr)
	}
	return
}

func GetGengoByTime(date time.Time) (g Gengo, e error) {
	rachet := time.Time{}
	for _, gengo := range GetGengos() {
		startedAt, err := time.Parse(time.DateOnly, gengo.StartedAt)
		if err != nil {
			e = err
			continue
		}
		if date.Compare(startedAt) >= 0 && startedAt.Compare(rachet) > 0 {
			g = gengo
			rachet = startedAt
		}
	}
	g.time = &date
	return
}

func GetGengoByKey(key string) Gengo {
	key = strings.ToLower(key)
	for _, g := range GetGengos() {
		gKey := strings.ToLower(g.Key)
		if gKey == key {
			return g
		}
	}
	return defaultGengo
}

func GetGengoByAbbr(abbr string) Gengo {
	abbr = strings.ToUpper(abbr)
	for _, g := range GetGengos() {
		if g.Abbr == abbr {
			return g
		}
	}
	return defaultGengo
}

func GetGengoByName(name string) Gengo {
	for _, g := range GetGengos() {
		if g.Name == name {
			return g
		}
	}
	return defaultGengo
}

func (g *Gengo) SetDate(y, m, d string) (t time.Time, e error) {
	startedAt, err := time.Parse(time.DateOnly, g.StartedAt)
	if err != nil {
		e = err
		return
	}

	var ny int
	if y == PrefixFirstYear || y == "1" {
		ny = 0
	} else {
		ny, _ = strconv.Atoi(y)
		ny -= 1
	}

	nm, _ := strconv.Atoi(m)
	nd, _ := strconv.Atoi(d)

	t = time.Date(startedAt.Year()+ny, time.Month(nm), nd, 0, 0, 0, 0, time.UTC)
	if t.Compare(startedAt) < 0 {
		newG, err := GetGengoByTime(t)
		if err != nil {
			e = err
			return
		}

		newY := fmt.Sprintf("%d", newG.Years()+ny)
		return newG.SetDate(newY, m, d)
	}
	g.time = &t

	return
}

func (g *Gengo) Years() (n int) {
	startedAt, _ := time.Parse(time.DateOnly, g.StartedAt)
	if startedAt.Month() != 1 && startedAt.Day() != 1 {
		n = 1
	}
	if len(g.EndedAt) < 1 {
		n = 999
		return
	}
	endedAt, _ := time.Parse(time.DateOnly, g.EndedAt)
	n += (endedAt.Year() - startedAt.Year())
	if endedAt.Month() != 1 && endedAt.Day() != 1 {
		n += 1
	}
	return
}
