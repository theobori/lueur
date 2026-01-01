package walker

import (
	"testing"

	"github.com/theobori/lueur/gophermap"
)

func TestWalkEmphasis(t *testing.T) {
	tests := []comparable{
		{
			source:   "**hello world**",
			expected: testEmptyGophermapLineString + "ihello world\t/\tlocalhost\t70\n",
		},
		{
			source:   "***hello *aa* world***",
			expected: testEmptyGophermapLineString + "ihello aa world\t/\tlocalhost\t70\n",
		},
		{
			source:   "_**ALINK**_ aa __c__",
			expected: testEmptyGophermapLineString + "iALINK aa c\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHeading(t *testing.T) {
	tests := []comparable{
		{
			source:   "# h1",
			expected: testEmptyGophermapLineString + "i# h1\t/\tlocalhost\t70\n",
		},
		{
			source:   "## h2",
			expected: testEmptyGophermapLineString + "i## h2\t/\tlocalhost\t70\n",
		},
		{
			source:   "### h3",
			expected: testEmptyGophermapLineString + "i### h3\t/\tlocalhost\t70\n",
		},
		{
			source:   "#### h4",
			expected: testEmptyGophermapLineString + "i#### h4\t/\tlocalhost\t70\n",
		},
		{
			source: `line1
line2
line3
-----`,
			expected: testEmptyGophermapLineString + `i## line1	/	localhost	70
i## line2	/	localhost	70
i## line3	/	localhost	70
`,
		},
	}

	localOptions := *testOptions
	localOptions.WriteFancyHeader = true

	testComparableMultipleHelper(t, tests, &localOptions)
}

func TestWalkAutoLink(t *testing.T) {
	tests := []comparable{
		{
			source: "https://a.com",
			expected: testEmptyGophermapLineString + `ihttps://a.com	/	localhost	70
hhttps://a.com	URL:https://a.com	a.com	443
`,
		},
		{
			source: "http://a.com",
			expected: testEmptyGophermapLineString + `ihttp://a.com	/	localhost	70
hhttp://a.com	URL:http://a.com	a.com	80
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkCodeBlock(t *testing.T) {
	tests := []comparable{
		{
			source: "```" + `
codeblock
codeblock
codeblock
codeblock
` + "```",
			expected: testEmptyGophermapLineString + `icodeblock	/	localhost	70
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
			expected: testEmptyGophermapLineString + `icodeblock	/	localhost	70
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

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkLink(t *testing.T) {
	tests := []comparable{
		{
			source: "[text link](https://example.com)",
			expected: testEmptyGophermapLineString + `itext link	/	localhost	70
htext link	URL:https://example.com	example.com	443
`,
		},
		{
			source: "[link](http://example.com)",
			expected: testEmptyGophermapLineString + `ilink	/	localhost	70
hlink	URL:http://example.com	example.com	80
`,
		},
		{
			source: "[multiple words link](https://test.org/path)",
			expected: testEmptyGophermapLineString + `imultiple words link	/	localhost	70
hmultiple words link	URL:https://test.org/path	test.org	443
`,
		},
		{
			source: "Some text [inline link](https://example.com) and more text",
			expected: testEmptyGophermapLineString + `iSome text inline link and more text	/	localhost	70
hinline link	URL:https://example.com	example.com	443
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkImage(t *testing.T) {
	tests := []comparable{
		{
			source: "![alt text](https://example.com/image.png)",
			expected: testEmptyGophermapLineString + `ialt text	/	localhost	70
halt text	URL:https://example.com/image.png	example.com	443
`,
		},
		{
			source: "![](/photo.jpg)",
			expected: testEmptyGophermapLineString + `i/photo.jpg	/	localhost	70
I/photo.jpg	/photo.jpg	localhost	70
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkBlockQuote(t *testing.T) {
	tests := []comparable{
		{
			source:   "> single line quote",
			expected: testEmptyGophermapLineString + "i“single line quote”\t/\tlocalhost\t70\n",
		},
		{
			source: `> first line
> second line
> third line`,
			expected: testEmptyGophermapLineString + `i“first line	/	localhost	70
isecond line	/	localhost	70
ithird line”	/	localhost	70
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkCodeSpan(t *testing.T) {
	tests := []comparable{
		{
			source:   "`hello world`",
			expected: testEmptyGophermapLineString + "ihello world\t/\tlocalhost\t70\n",
		},
		{
			source:   "`hello` `aa` `world` `a b c d`",
			expected: testEmptyGophermapLineString + "ihello aa world a b c d\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkParagraph(t *testing.T) {
	tests := []comparable{
		{
			source: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet eros nec interdum vestibulum. Sed elementum scelerisque euismod. Praesent pellentesque justo eu ex iaculis ullamcorper. Nulla suscipit purus quis sagittis sagittis. Sed eget tempus odio. Interdum et malesuada fames ac ante ipsum primis in faucibus. Suspendisse eget orci erat. Sed volutpat maximus urna eu commodo. Praesent tristique non nibh blandit ultricies. Etiam tempus nisi urna, non accumsan ante laoreet ac. Nam et lectus pharetra risus rhoncus facilisis. Suspendisse eu quam venenatis ipsum scelerisque scelerisque. `,
			expected: testEmptyGophermapLineString + `iLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet eros nec	/	localhost	70
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

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkDocument(t *testing.T) {
	tests := []comparable{
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

codeblockcodeblockcodeblockcodeblock
codeblockcodeblockcodeblock
codeblockcodeblockcodeblock
codeblockcodeblockcodeblockcodeblockcodeblock
codeblockcodeblock
codeblock

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
			expected: testEmptyGophermapLineString + `iConclusion	/	localhost	70
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
ihttps://example.com	/	localhost	70
hhttps://example.com	URL:https://example.com	example.com	443
i	/	localhost	70
ia	/	localhost	70
ia	/	localhost	70
ititle	/	localhost	70
ititle	/	localhost	70
ititle	/	localhost	70
ha	URL:https://a.com	a.com	443
ha	URL:https://a.com	a.com	443
Ititle	/dog.jpg	localhost	70
Ititle	/dog.jpg	localhost	70
Ititle	/dog.jpg	localhost	70
i	/	localhost	70
itelnet test	/	localhost	70
iphlog	/	localhost	70
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
iAlso see poison.gph	/	localhost	70
1Also see poison.gph	/poison.gph	localhost	70
i	/	localhost	70
iAs i was saying, i had a lot og pages like the b page	/	localhost	70
1the b page	/b.gophermap	localhost	70
i	/	localhost	70
iHello title, world.	/	localhost	70
Ititle	/dog.jpg	localhost	70
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkDocumentReferencesAfterAll(t *testing.T) {
	source := `[a](https://a.com) tttt uu vvvvv https://a.com

https://a.com
https://a.com
[a](https://a.com)

https://a.com

aaaa
aaaa

aaaa
aaaa
`

	testGopherMap := comparable{
		source: source, expected: testEmptyGophermapLineString + `i(a)[1] tttt uu vvvvv (https://a.com)[2]	/	localhost	70
i	/	localhost	70
i(https://a.com)[3]	/	localhost	70
i(https://a.com)[4]	/	localhost	70
i(a)[5]	/	localhost	70
i	/	localhost	70
i(https://a.com)[6]	/	localhost	70
i	/	localhost	70
iaaaa	/	localhost	70
iaaaa	/	localhost	70
i	/	localhost	70
iaaaa	/	localhost	70
iaaaa	/	localhost	70
h[1] a	URL:https://a.com	a.com	443
h[2] https://a.com	URL:https://a.com	a.com	443
h[3] https://a.com	URL:https://a.com	a.com	443
h[4] https://a.com	URL:https://a.com	a.com	443
h[5] a	URL:https://a.com	a.com	443
h[6] https://a.com	URL:https://a.com	a.com	443
`,
	}

	var localOptions Options

	localOptions = *testOptions
	localOptions.SetReferencePositionAndFileFormat(AfterTraverse, gophermap.FileFormatGophermap)
	testComparableHelper(t, testGopherMap, &localOptions)

	testTxt := comparable{
		source: source,
		expected: `
(a)[1] tttt uu vvvvv https://a.com

https://a.com
https://a.com
(a)[2]

https://a.com

aaaa
aaaa

aaaa
aaaa
[1] https://a.com
[2] https://a.com
`,
	}

	localOptions = *testOptions
	localOptions.SetReferencePositionAndFileFormat(AfterTraverse, gophermap.FileFormatTxt)
	testComparableHelper(t, testTxt, &localOptions)
}

func TestWalkList(t *testing.T) {
	tests := []comparable{
		{
			source: `- a
- a
- a
  - c
    - aa
    - d
      1. e
      2. e 
`,
			expected: testEmptyGophermapLineString + `i- a	/	localhost	70
i- a	/	localhost	70
i- a	/	localhost	70
i  - c	/	localhost	70
i    - aa	/	localhost	70
i    - d	/	localhost	70
i      1. e	/	localhost	70
i      2. e	/	localhost	70
`,
		},
		{
			source: `1. a
2. b Hello, world! aaa a aa [WW](https://google.fr). AA. d
3. c
   1. aa
      1. aa`,
			expected: `i	/	localhost	70
i1. a	/	localhost	70
i2. b Hello, world! aaa a aa WW. AA. d	/	localhost	70
i3. c	/	localhost	70
i  1. aa	/	localhost	70
i    1. aa	/	localhost	70
hWW	URL:https://google.fr	google.fr	443
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}
