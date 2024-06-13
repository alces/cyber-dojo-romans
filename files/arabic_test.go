package reverse_romans

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var arabicResults = []struct {
    roman  string
    arabic int
} {
    {"M", 1000},
    {"CM", 900},
}

func TestArabic(t *testing.T) {
    for _, r := range arabicResults {
        assert.Equal(t, r.arabic, Arabic(r.roman), r.roman)
    }
}
