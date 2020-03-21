package minimumswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwap(t *testing.T) {
	expected := 3
	actual := MinimumSwap([]int{4, 3, 1, 2})
	assert.Equal(t, expected, actual, "Not matched")
}
