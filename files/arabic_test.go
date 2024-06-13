package reverse_romans

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var arabicResults = []struct {
    roman  string
    arabic int
} {
    {"MXC", 1090},
    {"MMMCMXX", 3920},
    {"DCCXL", 740},
    {"DXC", 590},
    {"CDLXX", 470},
    {"CCCLIX", 359},
}

func TestArabic(t *testing.T) {
    for _, r := range arabicResults {
        assert.Equal(t, r.arabic, Arabic(r.roman), r.roman)
    }
}
