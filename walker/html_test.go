package walker

import "testing"

func TestWalkHTMLH1(t *testing.T) {
	test := comparable{
		source:   "<h1>Header 1</h1>",
		expected: "iHeader 1\t/\tlocalhost\t70\n",
	}

	testComparableHelper(t, test, testOptions)
}

func TestWalkHTMLH2(t *testing.T) {
	tests := []comparable{
		{
			source:   "<h2>Header 2</h2>",
			expected: "iHeader 2\t/\tlocalhost\t70\n",
		},
		{
			source:   "<h2>Head<h2>er</h2> 2</h2>",
			expected: "iHeader 2\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLH3(t *testing.T) {
	tests := []comparable{
		{
			source:   "<h3>Header 3</h3>",
			expected: "iHeader 3\t/\tlocalhost\t70\n",
		},
		{
			source:   "<h3>Head<h3>er</h3> 3</h3>",
			expected: "iHeader 3\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLH4(t *testing.T) {
	tests := []comparable{
		{
			source:   "<h4>Header 4</h4>",
			expected: "iHeader 4\t/\tlocalhost\t70\n",
		},
		{
			source:   "<h4>Head<h4>er</h4> 4</h4>",
			expected: "iHeader 4\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLH5(t *testing.T) {
	tests := []comparable{
		{
			source:   "<h5>Header 5</h5>",
			expected: "iHeader 5\t/\tlocalhost\t70\n",
		},
		{
			source:   "<h5>Head<h5>er</h5> 5</h5>",
			expected: "iHeader 5\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLH6(t *testing.T) {
	tests := []comparable{
		{
			source:   "<h6>Header 6</h6>",
			expected: "iHeader 6\t/\tlocalhost\t70\n",
		},
		{
			source:   "<h6>Head<h6>er</h6> 6</h6>",
			expected: "iHeader 6\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLB(t *testing.T) {
	tests := []comparable{
		{
			source:   "<b>Bold</b>",
			expected: testEmptyGophermapLineString + "iBold\t/\tlocalhost\t70\n",
		},
		{
			source:   "<b>Bo<b>ld</b></b>",
			expected: testEmptyGophermapLineString + "iBold\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLP(t *testing.T) {
	tests := []comparable{
		{
			source:   "<p>A simple paragraph.</p>",
			expected: testEmptyGophermapLineString + "iA simple paragraph.\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLDiv(t *testing.T) {
	tests := []comparable{
		{
			source:   "<div>A simple div.</div>",
			expected: "iA simple div.\t/\tlocalhost\t70\n",
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLImg(t *testing.T) {
	tests := []comparable{
		{
			source: `<img src="https://a.com" alt="a">`,
			expected: `ia	/	localhost	70
ha	URL:https://a.com	a.com	443
`,
		},
		{
			source: `<img src="https://a.com" alt="a"></img>`,
			expected: testEmptyGophermapLineString + `ia	/	localhost	70
ha	URL:https://a.com	a.com	443
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}

func TestWalkHTMLHTML(t *testing.T) {
	tests := []comparable{
		{
			source: `<div>
a
<p>aa <b>AAA</b> aa</p>
<b>AAA</b>
<img src="https://google.fr"></img>
b
</div>`,
			expected: testEmptyGophermapLineString + `ia	/	localhost	70
i	/	localhost	70
iaa AAA aa	/	localhost	70
i	/	localhost	70
iAAA	/	localhost	70
ihttps://google.fr	/	localhost	70
ib	/	localhost	70
hhttps://google.fr	URL:https://google.fr	google.fr	443
`,
		},
	}

	testComparableMultipleHelper(t, tests, testOptions)
}
