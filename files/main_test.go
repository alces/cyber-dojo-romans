package reverse_romans

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var testResults = []struct {
    roman  string
    arabic int
} {
    {"MXCIV", 1094},
    {"MMMCMXX", 3920},
    {"DCCXLV", 745},
    {"DXCIII", 593},
    {"CDLXXVIII", 478},
    {"CCCLIX", 359},
}

func TestArabic(t *testing.T) {
    for _, r := range testResults {
        assert.Equal(t, r.arabic, Arabic(r.roman), r.roman)
    }
}

func TestRoman(t *testing.T) {
    for _, r := range testResults {
        assert.Equal(t, r.roman, Roman(r.arabic), r.arabic)
    }
}