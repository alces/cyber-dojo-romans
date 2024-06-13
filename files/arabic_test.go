package reverse_romans

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var arabicResults = []struct {
    roman  string
    arabic int
} {
    {"MXCIV", 1094},
    {"MMMCMXX", 3920},
    {"DCCXLV", 745},
    {"DXCIII", 593},
    {"CDLXX", 470},
    {"CCCLIX", 359},
}

func TestArabic(t *testing.T) {
    for _, r := range arabicResults {
        assert.Equal(t, r.arabic, Arabic(r.roman), r.roman)
    }
}
