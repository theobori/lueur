package markdown

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/walker"
)

var (
	options walker.Options = walker.Options{
		WordWrapLimit:     80,
		ReferencePosition: walker.AfterBlocks,
		Domain:            "localhost",
		Port:              70,
		WriteFancyHeader:  false,
		WriteGPHFormat:    false,
	}

	emptyGophermapLine, _ = gophermap.NewLine(
		gophermap.ItemTypeInlineText,
		"",
		"/",
		options.Domain,
		options.Port,
	)
	emptyGophermapLineString string = emptyGophermapLine.String() + "\n"

	dmp *diffmatchpatch.DiffMatchPatch = diffmatchpatch.New()
)

type Comparable struct {
	source   string
	expected string
}

func testCompsHelper(t *testing.T, comps []Comparable, options *walker.Options) {
	for _, comp := range comps {
		w := NewWalkerWithOptions([]byte(comp.source), options)

		s, err := w.WalkFromRoot()
		if err != nil {
			t.Fatal(err)
		}

		if s != comp.expected {
			diff := dmp.DiffMain(s, comp.expected, false)
			prettyDiffString := dmp.DiffPrettyText(diff)

			t.Fatal(prettyDiffString)
		}
	}
}

func TestWalkEmphasis(t *testing.T) {
	tests := []Comparable{
		{
			source:   "**hello world**",
			expected: emptyGophermapLineString + "ihello world\t/\tlocalhost\t70\n",
		},
		{
			source:   "***hello *aa* world***",
			expected: emptyGophermapLineString + "ihello aa world\t/\tlocalhost\t70\n",
		},
		{
			source:   "_**ALINK**_ aa __c__",
			expected: emptyGophermapLineString + "iALINK aa c\t/\tlocalhost\t70\n",
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkHeading(t *testing.T) {
	tests := []Comparable{
		{
			source:   "# h1",
			expected: emptyGophermapLineString + "i# h1\t/\tlocalhost\t70\n",
		},
		{
			source:   "## h1",
			expected: emptyGophermapLineString + "i## h1\t/\tlocalhost\t70\n",
		},
		{
			source:   "### h1",
			expected: emptyGophermapLineString + "i### h1\t/\tlocalhost\t70\n",
		},
		{
			source:   "#### h1",
			expected: emptyGophermapLineString + "i#### h1\t/\tlocalhost\t70\n",
		},
		{
			source: `line1
line2
line3
-----`,
			expected: emptyGophermapLineString + `i## line1	/	localhost	70
i## line2	/	localhost	70
i## line3	/	localhost	70
`,
		},
	}

	localOptions := options
	localOptions.WriteFancyHeader = true

	testCompsHelper(t, tests, &localOptions)
}

func TestWalkAutoLink(t *testing.T) {
	tests := []Comparable{
		{
			source:   "https://a.com",
			expected: emptyGophermapLineString + "hhttps://a.com\tURL:https://a.com\ta.com\t443\n",
		},
		{
			source:   "http://a.com",
			expected: emptyGophermapLineString + "hhttp://a.com\tURL:http://a.com\ta.com\t80\n",
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkCodeBlock(t *testing.T) {
	tests := []Comparable{
		{
			source: "```" + `
codeblock
codeblock
codeblock
codeblock
` + "```",
			expected: emptyGophermapLineString + `icodeblock	/	localhost	70
icodeblock	/	localhost	70
icodeblock	/	localhost	70
icodeblock	/	localhost	70
`,
		},
		{
			source: "```" + `
codeblock



codeblock



codeblock



codeblock
` + "```",
			expected: emptyGophermapLineString + `icodeblock	/	localhost	70
i	/	localhost	70
i	/	localhost	70
i	/	localhost	70
icodeblock	/	localhost	70
i	/	localhost	70
i	/	localhost	70
i	/	localhost	70
icodeblock	/	localhost	70
i	/	localhost	70
i	/	localhost	70
i	/	localhost	70
icodeblock	/	localhost	70
`,
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkLink(t *testing.T) {
	tests := []Comparable{
		{
			source:   "[text link](https://example.com)",
			expected: emptyGophermapLineString + "htext link\tURL:https://example.com\texample.com\t443\n",
		},
		{
			source:   "[link](http://example.com)",
			expected: emptyGophermapLineString + "hlink\tURL:http://example.com\texample.com\t80\n",
		},
		{
			source:   "[multiple words link](https://test.org/path)",
			expected: emptyGophermapLineString + "hmultiple words link\tURL:https://test.org/path\ttest.org\t443\n",
		},
		{
			source: "Some text [inline link](https://example.com) and more text",
			expected: emptyGophermapLineString + `iSome text inline link and more text	/	localhost	70
hinline link	URL:https://example.com	example.com	443
`,
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkImage(t *testing.T) {
	tests := []Comparable{
		{
			source:   "![alt text](https://example.com/image.png)",
			expected: emptyGophermapLineString + "halt text\tURL:https://example.com/image.png\texample.com\t443\n",
		},
		{
			source:   "![](/photo.jpg)",
			expected: emptyGophermapLineString + "I\t/photo.jpg\tlocalhost\t70\n",
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkBlockQuote(t *testing.T) {
	tests := []Comparable{
		{
			source:   "> single line quote",
			expected: emptyGophermapLineString + "i“single line quote”\t/\tlocalhost\t70\n",
		},
		{
			source: `> first line
> second line
> third line`,
			expected: emptyGophermapLineString + `i“first line	/	localhost	70
isecond line	/	localhost	70
ithird line”	/	localhost	70
`,
		},
	}

	testCompsHelper(t, tests, &options)
}
