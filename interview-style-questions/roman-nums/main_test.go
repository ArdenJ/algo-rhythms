package main

import "testing"

func TestNumerals(t *testing.T) {
	testCases := []struct {
		Description string
		Integer     int
		Want        string
	}{
		{"1 returns 'I'", 1, "I"},
		{"2 returns 'II'", 2, "II"},
		{"3 returns 'III'", 3, "III"},
		{"4 returns 'IV' (can't repeat more than 3 times)", 4, "IV"},
		{"5 returns V", 5, "V"},
		{"6 returns VI", 6, "VI"},
		{"7 returns VII", 7, "VII"},
		{"8 returns VIII", 8, "VIII"},
		{"9 returns IX", 9, "IX"},
		{"10 returns X", 10, "X"},
		{"14 returns XIV", 14, "XIV"},
		{"18 returns XVIII", 18, "XVIII"},
		{"20 returns XX", 20, "XX"},
		{"39 returns XXXIX", 39, "XXXIX"},
		{"40 returns XL", 40, "XL"},
		{"47 returns XLVII", 47, "XLVII"},
		{"49 returns XLIX", 49, "XLIX"},
		{"50 returns L", 50, "L"},
		{"1984 returns MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 returns MMMCMXCIX", 3999, "MMMCMXCIX"},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			got := ToNumerals(test.Integer)
			assertConversion(t, got, test.Want)
		})
	}
}

func assertConversion(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
