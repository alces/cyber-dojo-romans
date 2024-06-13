package reverse_romans

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestArabic(t *testing.T) {
    assert.Equal(t, 1000, Arabic("M"), "M")
}
