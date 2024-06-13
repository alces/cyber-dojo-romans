package reverse_romans

import (
    "strings"
)

func Arabic(roman string) (arabic int) {
    for strings.HasPrefix(roman, "M") {
        arabic += 1000
        roman = roman[1:]
    }
    
    if strings.HasPrefix(roman, "CM") {
        arabic += 900
        roman = roman[1:]
    }
    
    return
}
