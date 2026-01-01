package walker

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/theobori/lueur/gophermap"
)

var (
	testOptions, _ = NewOptions(
		80,
		AfterBlocks,
		"localhost",
		70,
		false,
		gophermap.FileFormatGophermap,
	)

	testEmptyGophermapLine, _ = gophermap.NewLine(
		gophermap.ItemTypeInlineText,
		"",
		"/",
		testOptions.Domain(),
		testOptions.Port(),
	)
	testEmptyGophermapLineString string = testEmptyGophermapLine.String() + "\n"

	testDmp *diffmatchpatch.DiffMatchPatch = diffmatchpatch.New()
)

type comparable struct {
	source   string
	expected string
}

func testComparableHelper(t *testing.T, comp comparable, options *Options) {
	w := NewWalkerWithOptions([]byte(comp.source), options)

	s, err := w.WalkFromRoot()
	if err != nil {
		t.Fatal(err)
	}

	if s != comp.expected {
		diff := testDmp.DiffMain(s, comp.expected, false)
		prettyDiffString := testDmp.DiffPrettyText(diff)

		t.Fatal(prettyDiffString)
	}
}

func testComparableMultipleHelper(t *testing.T, comps []comparable, options *Options) {
	for _, comp := range comps {
		testComparableHelper(t, comp, options)
	}
}
