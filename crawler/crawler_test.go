package crawler

import (
	"fmt"
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestDoogleCrawler_analyze(t *testing.T) {
	crawler, _ := NewCrawler()
	cr := crawler.(*doogleCrawler)

	for i, cc := range []struct {
		target    string
		expTitle  string
		expEdges  []string
		expTokens []string
	}{
		{
			target: `
<!DOCTYPE html><html>
	<header>
		<title>title1</title>
	</header>
	<body>
		<a href="https://www.google.com">
	</body>
</html>`,
			expTitle:  "title1",
			expEdges:  []string{"https://www.google.com"},
			expTokens: []string{"title1"},
		},
		{
			target: `
<!DOCTYPE html><html>
	<header>
		<title>This is a pen</title>
	</header>
	<body>
		<a href="https://www.google.com"> 123456 </a>
		<a href="https://www.doogle.com"> 123456 </a>
	</body>
</html>`,
			expTitle:  "This is a pen",
			expEdges:  []string{"https://www.google.com", "https://www.doogle.com"},
			expTokens: []string{"This", "is", "a", "pen", "123456", "123456"},
		},
		{
			target: `
<!DOCTYPE html><html>
	<header>
		<title>This is a pen 100yen</title>
	</header>
	<body>
		<a href="https://www.google.com"> 123456 </a>
		<a href="https://www.doogle.com"> 123456 </a>
		<p> this is first text field</p>
	</body>
</html>`,
			expTitle:  "This is a pen 100yen",
			expEdges:  []string{"https://www.google.com", "https://www.doogle.com"},
			expTokens: []string{"This", "is", "a", "pen", "100yen", "123456", "123456", "this", "is", "first", "text", "field"},
		},
		{
			target: `
<!DOCTYPE html><html>
	<header>
		<title>This is a pen 100yen</title>
	</header>
	<body>
		<a href="https://www.google.com"> 123456 </a>
		<a href="htt://www.doogle.com"> 123456 </a>
		<a href="/img/cat.jpg"></a>
		<p> this is first text field</p>
	</body>
</html>`,
			expTitle:  "This is a pen 100yen",
			expEdges:  []string{"https://www.google.com"},
			expTokens: []string{"This", "is", "a", "pen", "100yen", "123456", "123456", "this", "is", "first", "text", "field"},
		},
	} {
		c := cc
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			body := strings.NewReader(c.target)
			aTitle, aTokens, aEdgeURLs, err := cr.analyze(body)
			if err != nil {
				panic(err)
			}
			assert.Equal(t, c.expTitle, aTitle)
			assert.Equal(t, len(c.expEdges), len(aEdgeURLs))

			for i := range c.expEdges {
				assert.Equal(t, c.expEdges[i], aEdgeURLs[i])
			}

			assert.Equal(t, len(c.expTokens), len(aTokens))

			for i := range c.expTokens {
				assert.Equal(t, c.expTokens[i], aTokens[i])
			}
		})
	}
}
