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
