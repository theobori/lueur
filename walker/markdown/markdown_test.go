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
		FileFormat:        gophermap.FileFormatGophermap,
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

func TestWalkCodeSpan(t *testing.T) {
	tests := []Comparable{
		{
			source:   "`hello world`",
			expected: emptyGophermapLineString + "ihello world\t/\tlocalhost\t70\n",
		},
		{
			source:   "`hello` `aa` `world` `a b c d`",
			expected: emptyGophermapLineString + "ihello aa world a b c d\t/\tlocalhost\t70\n",
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkParagraph(t *testing.T) {
	tests := []Comparable{
		{
			source: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet eros nec interdum vestibulum. Sed elementum scelerisque euismod. Praesent pellentesque justo eu ex iaculis ullamcorper. Nulla suscipit purus quis sagittis sagittis. Sed eget tempus odio. Interdum et malesuada fames ac ante ipsum primis in faucibus. Suspendisse eget orci erat. Sed volutpat maximus urna eu commodo. Praesent tristique non nibh blandit ultricies. Etiam tempus nisi urna, non accumsan ante laoreet ac. Nam et lectus pharetra risus rhoncus facilisis. Suspendisse eu quam venenatis ipsum scelerisque scelerisque. `,
			expected: emptyGophermapLineString + `iLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet eros nec	/	localhost	70
iinterdum vestibulum. Sed elementum scelerisque euismod. Praesent pellentesque	/	localhost	70
ijusto eu ex iaculis ullamcorper. Nulla suscipit purus quis sagittis sagittis.	/	localhost	70
iSed eget tempus odio. Interdum et malesuada fames ac ante ipsum primis in	/	localhost	70
ifaucibus. Suspendisse eget orci erat. Sed volutpat maximus urna eu commodo.	/	localhost	70
iPraesent tristique non nibh blandit ultricies. Etiam tempus nisi urna, non	/	localhost	70
iaccumsan ante laoreet ac. Nam et lectus pharetra risus rhoncus facilisis.	/	localhost	70
iSuspendisse eu quam venenatis ipsum scelerisque scelerisque.	/	localhost	70
`,
		},
	}

	testCompsHelper(t, tests, &options)
}

func TestWalkDocument(t *testing.T) {
	tests := []Comparable{
		{
			source: `# Conclusion

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed [_**ALINK**_](https://a.com) eros nec interdum vestibulum. Sed elementum scelerisque euismod. Praesent pellentesque justo eu ex iaculis ullamcorper. Nulla suscipit purus quis sagittis sagittis. Sed eget tempus odio. Interdum et malesuada fames ac ante https://example.com primis in faucibus. Suspendisse eget orci erat. Sed volutpat maximus [**BLINK**](https://a.com) eu commodo. Praesent tristique non nibh blandit [**BLINK**](https://a.com). Etiam tempus nisi urna, non accumsan ante laoreet ac. Nam et lectus pharetra risus rhoncus facilisis. Suspendisse eu quam venenatis ipsum ![hello][id] scelerisque. 

**a**

*bb*
__c__
_c_

**aa
ddd**

a  
b  
c

## H2

` + "```" +
				`
codeblockcodeblockcodeblockcodeblock
codeblockcodeblockcodeblock
codeblockcodeblockcodeblock
codeblockcodeblockcodeblockcodeblockcodeblock
codeblockcodeblock
codeblock
` + "```" +
				`



hello
## HELLO a a\

aa
d
a
a
-------

Hello world https://example.com [**my-link**](https://a.com)


https://example.com

[a](https://a.com)
[a](https://a.com)
![alt][id]
![alt][id]
![alt][id]

[telnet test](telnet://a.com)
[phlog](phlog)


> block quote tes d d a dd
> bb
> b
> b
> 
> C c c c **AA**
> # a a
> # d
> 
> > c
> > c
> # aa

aaa
a

Hello ![alt](cat.png), world.

[Also see poison.gph](/poison.gph)

As i was saying, i had a lot _og_ pages like [the b page](/b.gophermap)

Hello ![alt][id], world.

[id]: dog.jpg "title"
`,
			expected: emptyGophermapLineString + `iConclusion	/	localhost	70
i	/	localhost	70
iLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ALINK eros nec	/	localhost	70
iinterdum vestibulum. Sed elementum scelerisque euismod. Praesent pellentesque	/	localhost	70
ijusto eu ex iaculis ullamcorper. Nulla suscipit purus quis sagittis sagittis.	/	localhost	70
iSed eget tempus odio. Interdum et malesuada fames ac ante https://example.com	/	localhost	70
iprimis in faucibus. Suspendisse eget orci erat. Sed volutpat maximus BLINK eu	/	localhost	70
icommodo. Praesent tristique non nibh blandit BLINK. Etiam tempus nisi urna, non	/	localhost	70
iaccumsan ante laoreet ac. Nam et lectus pharetra risus rhoncus facilisis.	/	localhost	70
iSuspendisse eu quam venenatis ipsum title scelerisque.	/	localhost	70
hALINK	URL:https://a.com	a.com	443
hhttps://example.com	URL:https://example.com	example.com	443
hBLINK	URL:https://a.com	a.com	443
hBLINK	URL:https://a.com	a.com	443
Ititle	/dog.jpg	localhost	70
i	/	localhost	70
ia	/	localhost	70
i	/	localhost	70
ibb	/	localhost	70
ic	/	localhost	70
ic	/	localhost	70
i	/	localhost	70
iaa	/	localhost	70
iddd	/	localhost	70
i	/	localhost	70
ia	/	localhost	70
ib	/	localhost	70
ic	/	localhost	70
i	/	localhost	70
iH2	/	localhost	70
i	/	localhost	70
icodeblockcodeblockcodeblockcodeblock	/	localhost	70
icodeblockcodeblockcodeblock	/	localhost	70
icodeblockcodeblockcodeblock	/	localhost	70
icodeblockcodeblockcodeblockcodeblockcodeblock	/	localhost	70
icodeblockcodeblock	/	localhost	70
icodeblock	/	localhost	70
i	/	localhost	70
ihello	/	localhost	70
iHELLO a a	/	localhost	70
i	/	localhost	70
iaa	/	localhost	70
id	/	localhost	70
ia	/	localhost	70
ia	/	localhost	70
i	/	localhost	70
iHello world https://example.com my-link	/	localhost	70
hhttps://example.com	URL:https://example.com	example.com	443
hmy-link	URL:https://a.com	a.com	443
i	/	localhost	70
hhttps://example.com	URL:https://example.com	example.com	443
i	/	localhost	70
ha	URL:https://a.com	a.com	443
ha	URL:https://a.com	a.com	443
Ititle	/dog.jpg	localhost	70
Ititle	/dog.jpg	localhost	70
Ititle	/dog.jpg	localhost	70
i	/	localhost	70
8telnet test	user	a.com	23
1phlog	/phlog	localhost	70
i	/	localhost	70
i“block quote tes d d a dd	/	localhost	70
ibb	/	localhost	70
ib	/	localhost	70
ib	/	localhost	70
i	/	localhost	70
iC c c c AA	/	localhost	70
ia a	/	localhost	70
id	/	localhost	70
i	/	localhost	70
i“c	/	localhost	70
ic”	/	localhost	70
iaa”	/	localhost	70
i	/	localhost	70
iaaa	/	localhost	70
ia	/	localhost	70
i	/	localhost	70
iHello alt, world.	/	localhost	70
Ialt	/cat.png	localhost	70
i	/	localhost	70
1Also see poison.gph	/poison.gph	localhost	70
i	/	localhost	70
iAs i was saying, i had a lot og pages like the b page	/	localhost	70
1the b page	/b.gophermap	localhost	70
i	/	localhost	70
iHello title, world.	/	localhost	70
Ititle	/dog.jpg	localhost	70
i	/	localhost	70
`,
		},
	}

	testCompsHelper(t, tests, &options)
}
