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
    {"MMCM", 2900},
    {"DCCXL", 740},
    {"DXC", 590},
    {"CDLXX", 470},
    {"CCCL", 350},
}

func TestArabic(t *testing.T) {
    for _, r := range arabicResults {
        assert.Equal(t, r.arabic, Arabic(r.roman), r.roman)
    }
}
