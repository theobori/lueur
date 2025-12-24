package common

type Counter struct {
	value int
	delta int
}

func NewCounter(delta int) *Counter {
	if delta < 0 {
		delta = -delta
	}

	return &Counter{
		value: 0,
		delta: delta,
	}
}

func NewDefaultCounter() *Counter {
	return NewCounter(1)
}

func (c *Counter) Add() {
	c.value += c.delta
}

func (c *Counter) Remove() {
	c.value -= c.delta

	if c.value < 0 {
		c.value = 0
	}
}

func (c *Counter) Reset() {
	c.value = 0
}

func (c *Counter) Value() int {
	return c.value
}
