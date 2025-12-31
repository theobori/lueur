package walker

import (
	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/internal/common"
)

type Context struct {
	Depth           *common.Counter
	ReferencesQueue []gophermap.Line
	Indentation     *common.Indentation
}

func NewDefaultContext() *Context {
	return &Context{
		Depth:           common.NewDefaultCounter(),
		ReferencesQueue: []gophermap.Line{},
		Indentation:     common.NewDefaultIndentation(),
	}
}

func (c *Context) Reset() {
	c.ClearQueues()
	c.Depth.Reset()
	c.Indentation.Reset()
}

func (c *Context) ClearQueues() {
	c.ReferencesQueue = nil
}
