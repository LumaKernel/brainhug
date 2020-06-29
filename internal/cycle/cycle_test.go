package cycle_test
import (
	"github.com/LumaKernel/brainhug/internal/cycle"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	c := new(cycle.Cycle)
	c.AppendFront(0)
	c.AppendFront(1)
	c.AppendFront(2)

	assert.Equal(t, c.Get(0), byte(0), "")
	assert.Equal(t, c.Get(1), byte(1), "")
	assert.Equal(t, c.Get(2), byte(2), "")

	c.Set(1, c.Get(0) + c.Get(2) + 10)
	assert.Equal(t, c.Get(1), byte(12), "")

	c.AppendBack(3)
	assert.Equal(t, c.Get(-1), byte(3), "")
	c.AppendFront(4)
	assert.Equal(t, c.Get(3), byte(4), "")
	c.AppendBack(5)
	assert.Equal(t, c.Get(-2), byte(5), "")
	c.AppendFront(6)
	assert.Equal(t, c.Get(4), byte(6), "")

	c.Set(-2, c.Get(-1) + c.Get(1))
	assert.Equal(t, c.Get(-2), byte(3 + 12), "")

	c.Set(-2, 10)
	c.AppendBack(0)
	c.AppendBack(0)
	c.AppendBack(0)
	assert.Equal(t, c.Get(-2), byte(10), "")
}
