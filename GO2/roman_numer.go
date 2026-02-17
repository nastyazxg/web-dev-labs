package main

import (
	"errors"
	"strings"
)

type RomanNumeralParser struct {
	mapping map[rune]int
}

func NewRomanParser() *RomanNumeralParser {
	return &RomanNumeralParser{
		mapping: map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		},
	}
}

func (p *RomanNumeralParser) ParseRoman(input string) (int, error) {
	if input == "" {
		return 0, errors.New("входная строка пуста")
	}

	input = strings.ToUpper(input)

	for _, ch := range input {
		if _, exists := p.mapping[ch]; !exists {
			return 0, errors.New("обнаружен недопустимый символ: " + string(ch))
		}
	}

	total := 0
	previous := 0

	for i := len(input) - 1; i >= 0; i-- {
		current := p.mapping[rune(input[i])]

		if current < previous {
			total -= current
		} else {
			total += current
		}
		previous = current
	}

	if !p.isValidCombination(input, total) {
		return 0, errors.New("некорректная комбинация римских цифр")
	}

	return total, nil
}

func (p *RomanNumeralParser) FormatRoman(number int) (string, error) {
	if number <= 0 || number > 3999 {
		return "", errors.New("число должно находиться в диапазоне 1-3999")
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	output := ""
	remaining := number

	for i := 0; i < len(nums); i++ {
		for remaining >= nums[i] {
			output += letters[i]
			remaining -= nums[i]
		}
	}

	return output, nil
}

func (p *RomanNumeralParser) isValidCombination(roman string, expected int) bool {
	converted, err := p.FormatRoman(expected)
	if err != nil {
		return false
	}
	return converted == roman
}

func (p *RomanNumeralParser) GetExamplePairs() map[string]int {
	return map[string]int{
		"I":    1,
		"IV":   4,
		"V":    5,
		"IX":   9,
		"X":    10,
		"XL":   40,
		"L":    50,
		"XC":   90,
		"C":    100,
		"CD":   400,
		"D":    500,
		"CM":   900,
		"M":    1000,
		"MMXX": 2020,
		"MCML": 1950,
	}
}
