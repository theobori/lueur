package walker

import (
	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/internal/common"
)

type Context struct {
	Depth           common.Counter
	Indentation     common.Counter
	ReferencesQueue []gophermap.Line
}

func (c *Context) Reset() {
	c.ClearQueues()
	c.Depth.Reset()
	c.Indentation.Reset()
}

func (c *Context) ClearQueues() {
	c.ReferencesQueue = nil
}

func NewDefaultContext() *Context {
	return &Context{
		Depth:           *common.NewDefaultCounter(),
		Indentation:     *common.NewDefaultCounter(),
		ReferencesQueue: []gophermap.Line{},
	}
}
