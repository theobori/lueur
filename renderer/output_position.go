package renderer

import "fmt"

type OutputPosition int

const (
	// The node will be outputed after blocks at depth 0
	AfterBlocks OutputPosition = iota
	// The node will be outputed after the AST has been evaluated
	AfterTraverse
)

func NewOutputPositionFromString(s string) (OutputPosition, error) {
	switch s {
	case "after-block":
		return AfterBlocks, nil
	case "after-all":
		return AfterTraverse, nil
	default:
		return AfterBlocks, fmt.Errorf("unsupported string value: %s", s)
	}
}
