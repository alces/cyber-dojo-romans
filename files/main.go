package reverse_romans

import (
    "strings"
)

var conversions = []struct {
    roman  string
    arabic int
} {
    {"M", 1000},
    {"CM", 900},
    {"D", 500},
    {"CD", 400},
    {"C", 100},
    {"XC", 90},
    {"L", 50},
    {"XL", 40},
    {"X", 10},
    {"IX", 9},
    {"V", 5},
    {"IV", 4},
    {"I", 1},
}

func Arabic(roman string) (arabic int) {
    for _, c := range conversions {
        for strings.HasPrefix(roman, c.roman) {
            arabic += c.arabic
            roman = roman[len(c.roman):]
        }
    }
    
    return
}

func Roman(arabic int) (roman string) {
    for _, c := range conversions {
        for arabic/c.arabic > 0 {
            roman += c.roman
            arabic -= c.arabic
        }
    }
    
    return
}
