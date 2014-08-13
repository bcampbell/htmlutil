package htmlutil

import (
	"code.google.com/p/go.net/html"
	//"fmt"
	"strings"
	"testing"
)

// tests for renderNode()
func TestRenderNode(t *testing.T) {
	testData := []struct {
		htmlFragment string
		expected     string // raw text
	}{
		{"<p>Hello.</p>", "Hello.\n"},
		{"<p>Hello</p><p>there.</p>", "Hello\nthere.\n"},
		{"<p><em>Hello</em> there.</p>", "Hello there.\n"},
		{`<p><span class="cap">T</span>he cat sat on the mat.</p>`, "The cat sat on the mat.\n"},
	}

	for _, dat := range testData {
		nodes, err := html.ParseFragment(strings.NewReader(dat.htmlFragment), nil)
		if err != nil {
			panic(err)
		}
		got := ""
		for _, node := range nodes {
			got += RenderNode(node)
		}
		if got != dat.expected {
			t.Errorf("RenderNode(`%s`) got `%s` (expected `%s`)", dat.htmlFragment, got, dat.expected)
		}
	}
}
