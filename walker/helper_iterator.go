package walker

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

func (w *Walker) walkIteratorHelper(node ast.Node) (string, error) {
	// a good optimization here is to use a string builder when iterating instead
	// of an immutable string
	//
	// With an immutable string on a Markdown file with >60k lines:
	//________________________________________________________
	// Executed in    4.52 secs    fish           external
	// usr time   18.02 secs    0.32 millis   18.02 secs
	// sys time    1.03 secs    1.77 millis    1.03 secs
	//
	// With a string builder
	//
	// ________________________________________________________
	// Executed in  197.16 millis    fish           external
	// usr time  311.73 millis    1.13 millis  310.60 millis
	// sys time   57.08 millis    1.02 millis   56.06 millis
	//
	s := strings.Builder{}

	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		res, err := w.Walk(c)
		if err != nil {
			return "", err
		}

		_, err = s.WriteString(res)
		if err != nil {
			return "", err
		}
	}

	return s.String(), nil
}
