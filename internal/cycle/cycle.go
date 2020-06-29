package cycle

import "fmt"

type Cycle struct {
	back []byte

	usedFront int64
	usedBack int64
}

func (c *Cycle) extendIfNeccessary() {
	if c.usedFront + c.usedBack == int64(len(c.back)) {
		if len(c.back) == 0 {
			c.back = []byte{0}
		} else {
			c.back = append(c.back, c.back...)
		}
	}
}

func (c *Cycle) getInternalIndex(i int64) int64 {
	if i < 0 {
		i += int64(len(c.back))
	}
	return i
}

func (c *Cycle) Len() int64 {
	return int64(c.usedFront + c.usedBack)
}

func (c *Cycle) LenPos() int64 {
	return c.usedFront
}

func (c *Cycle) LenNeg() int64 {
	return c.usedBack
}

func (c *Cycle) AppendFront(v byte) {
	c.extendIfNeccessary()
	c.back[c.usedFront] = v
	c.usedFront++
}

func (c *Cycle) AppendBack(v byte) {
	c.extendIfNeccessary()
	c.usedBack++
	c.back[len(c.back) - int(c.usedBack)] = v
}

func (c *Cycle) checkBounary (i int64) {
	if c.usedFront <= i || i < -c.usedBack {
		panic (
			fmt.Sprintf(
				"[cycle] Out of range: accessed [%d], while %d ~ %d available",
				i,
				-c.usedBack,
				c.usedFront,
			),
		)
	}
}

func (c *Cycle) Set(i int64, v byte) {
	c.checkBounary(i)
	c.back[c.getInternalIndex(i)] = v
}

func (c *Cycle) Get(i int64) byte {
	c.checkBounary(i)
	return c.back[c.getInternalIndex(i)]
}
