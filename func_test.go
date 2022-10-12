package gen

import (
	"testing"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestFunc_Body(t *testing.T) {
	type fields struct {
		name    string
		params  *orderedmap.OrderedMap[string, string]
		returns *orderedmap.OrderedMap[string, string]
		g       *Generator
	}
	type args struct {
		fn func()
	}
	tests := []struct {
		name     string
		res      func(*Generator) string
		expected string
	}{
		{
			name: "Test",
			res: func(g *Generator) string {
				g.Func("Test").Param("data", "Type").Param("two", "Params").Body(func() {})
				return g.String()
			},
			expected: "func Test(data Type,two Params) {\n}\n\n",
		},
		{
			name: "Test",
			res: func(g *Generator) string {
				g.Func("Test").Param("data", "Type").Param("two", "Params").Body(func() {})
				return g.String()
			},
			expected: "func Test(data Type,two Params) {\n}\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator()
			res := tt.res(g)
			if res != tt.expected {
				t.Error(res, tt.expected)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.expected, res, false)
				t.Error("Diff:", dmp.DiffPrettyText(diffs))
			}
		})
	}
}
