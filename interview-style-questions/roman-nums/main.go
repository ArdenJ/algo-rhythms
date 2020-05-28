package main

import "strings"

// Decoupling data from behaviour makes code more declarative
// we've declared some rules about the numerals as DATA and applied
// that to the algorithm
type Numeral struct {
	Value  int
	Symbol string
}

var allNumerals = []Numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToNumerals(arg int) string {

	var result strings.Builder

	for _, numeral := range allNumerals {
		for arg >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arg -= numeral.Value
		}
	}

	return result.String()

}
