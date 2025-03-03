package gjpd

import (
	"testing"
)

func TestZenkakuToNumber(t *testing.T) {

	str := "０１２３４５６７８９"

	actual := ZenkakuToNumber(str)
	expected := "0123456789"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}

func TestNumberToZenkaku(t *testing.T) {

	str := "0123456789"
	actual := NumberToZenkaku(str)
	expected := "０１２３４５６７８９"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}

func TestKanjiToNumber(t *testing.T) {

	str := "〇一二三四五六七八九"
	actual := KanjiToNumber(str)
	expected := "0123456789"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
func TestNumberToKanji(t *testing.T) {

	str := "0123456789"
	actual := NumberToKanji(str)
	expected := "〇一二三四五六七八九"
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}
