package gjpd

import (
	"strings"
)

func ZenkakuNumbers() [][]string {
	return [][]string{
		{"０", "0"},
		{"１", "1"},
		{"２", "2"},
		{"３", "3"},
		{"４", "4"},
		{"５", "5"},
		{"６", "6"},
		{"７", "7"},
		{"８", "8"},
		{"９", "9"},
	}
}

func KanjiNumbers() [][]string {
	return [][]string{
		{"〇", "0"},
		{"一", "1"},
		{"二", "2"},
		{"三", "3"},
		{"四", "4"},
		{"五", "5"},
		{"六", "6"},
		{"七", "7"},
		{"八", "8"},
		{"九", "9"},
	}
}

func ZenkakuToNumber(str string) string {
	for _, p := range ZenkakuNumbers() {
		str = strings.Replace(str, p[0], p[1], -1)
	}
	return str
}
func NumberToZenkaku(str string) string {
	for _, p := range ZenkakuNumbers() {
		str = strings.Replace(str, p[1], p[0], -1)
	}
	return str
}

func KanjiToNumber(str string) string {
	for _, p := range KanjiNumbers() {
		str = strings.Replace(str, p[0], p[1], -1)
	}
	return str
}

func NumberToKanji(str string) string {
	for _, p := range KanjiNumbers() {
		str = strings.Replace(str, p[1], p[0], -1)
	}
	return str
}
