package numeral

import (
  "testing"
  "fmt"
  "testing/quick"
)

var cases = []struct{
    Arabic      int
    Roman       string
  }{
    {1, "I"},
    {2, "II"},
    {3, "III"},
    {4, "IV"},
    {5, "V"},
    {6, "VI"},
    {7, "VII"},
    {8, "VIII"},
    {9, "IX"},
    {10, "X"},
    {14, "XIV"},
    {18, "XVIII"},
    {20, "XX"},
    {39, "XXXIX"},
    {40, "XL"},
    {47, "XLVII"},
    {49, "XLIX"},
    {50, "L"},
    {Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
  }


func TestConvertingToRoman(t *testing.T) {
  for _, test := range cases {
    t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
      got := ConvertToRoman(test.Arabic)

      if got != test.Roman {
        t.Errorf("got %q want %q", got, test.Roman)
      }
    })
  }
}

func TestConvertingToArabic(t *testing.T) {
  for _, test := range cases[:] {
    t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
      got := ConvertToArabic(test.Roman)

      if got != test.Arabic {
        t.Errorf("got %d want %d", got, test.Arabic)
      }
    })
  }
}

func TestPropertiesOfConversion(t *testing.T) {
  assertion := func(arabic uint16) bool {
    if arabic > 3999 {
      return true
    }
    roman := ConvertToRoman(int(arabic))
    fromRoman := ConvertToArabic(roman)
    return fromRoman == int(arabic)
  }

  if err := quick.Check(assertion, &quick.Config{
      MaxCount: 1000,
  }); err != nil {
    t.Error("failed checks", err)
  }
}


