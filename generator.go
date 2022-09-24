package gen

import (
	"strings"

	"github.com/elliotchance/orderedmap/v2"
)

var newLine = byte('\n')

type KV map[string]string

type Generator struct {
	buf strings.Builder
}

func NewGenerator() *Generator {
	return &Generator{}
}

// Line adds all params to the buffer and appends a new line.
func (g *Generator) Line(ss ...string) {
	g.Lit(ss...)
	g.NewLine()
}

// Go does the same as Line(). It is just nicer to look at, when generating Go code.
func (g *Generator) Go(ss ...string) {
	g.Lit(ss...)
	g.NewLine()
}

// Id adds an identifier to the buffer.
func (g *Generator) Id(ss ...string) string {
	s := ""
	if len(ss) == 1 {
		s = ss[0]
	} else if len(ss) == 2 {
		s = ss[0] + "." + ss[1]
	} else {
		panic("RetSlice must receive 1 or 2 parameters")
	}

	return s
}

// S joins all passed params. Convience function to make the template shorter.
func (g *Generator) S(ss ...string) string {
	return strings.Join(ss, "")
}

// Append is a shorthand for generating "append" code.
func (g *Generator) Append(name string, value string) {
	g.Lit(name, " = append(", name, ", ", value, ")")
	g.NewLine()
}

// Return adds a return stmt to the buffer.
func (g *Generator) Return(s string) {
	g.Lit("return ", s)
	g.NewLine()
}

// Panic adds a panic stmt to the buffer.
func (g *Generator) Panic(s string) {
	g.Lit("panic (", s, ")")
	g.NewLine()
}

// Package adds a packe stmt to the buffer.
func (g *Generator) Package(s string) {
	g.Lit("package ", s)
	g.NewLine()
	g.NewLine()
}

// Comment adds a line comment stmt to the buffer.
func (g *Generator) Comment(ss ...string) {
	g.Lit("// ")
	g.Lit(ss...)
	g.NewLine()
}

// Const adds a const block. Callback is used to make the template indentation correct and flexible.
func (g *Generator) ConstFn(fn func()) {
	g.Lit("const (")
	g.NewLine()

	fn()

	g.Lit(")")
	g.NewLine()
}

// Import adds a single grouped import to the buffer.
func (g *Generator) Import(ss ...string) {
	g.Lit("import (")
	g.NewLine()

	for i := range ss {
		g.Lit(`"`, ss[i], `"`)
		g.NewLine()
	}
	g.Lit(")")
	g.NewLine()
	g.NewLine()
}

// Struct definition with callback for indentation and flexibility.
func (g *Generator) Struct(name string, fn func()) {
	g.Lit("type ", name, " struct {")
	g.NewLine()

	fn()

	g.Lit("}")
	g.NewLine()
}

// If adds the params as a condition in an if stmt.
func (g *Generator) If(ss ...string) *IfStmt {
	g.Lit("if (")
	g.Lit(ss...)
	g.Lit(") {")
	g.NewLine()

	ifstmt := &IfStmt{}
	ifstmt.g = g
	return ifstmt
}

// Method returns a Method with params joined for the name.
func (g *Generator) Method(ss ...string) *Method {
	m := &Method{}

	m.name = strings.Join(ss, "")
	m.g = g

	m.params = orderedmap.NewOrderedMap[string, string]()
	m.returns = orderedmap.NewOrderedMap[string, string]()

	return m
}

// Func returns a Func with params joined for the name.
func (g *Generator) Func(ss ...string) *Func {
	f := &Func{}

	f.name = strings.Join(ss, "")
	f.g = g

	f.params = orderedmap.NewOrderedMap[string, string]()
	f.returns = orderedmap.NewOrderedMap[string, string]()

	return f
}

// Field returns a Field.
func (g *Generator) Field(name string) *Field {
	f := &Field{}

	f.name = name
	f.g = g

	f.tags = orderedmap.NewOrderedMap[string, string]()

	return f
}

// Lit adds passes params to the buffer.
func (g *Generator) Lit(ss ...string) {
	for i := range ss {
		g.buf.WriteString(ss[i])
	}
}

// NewLine added to the buffer.
func (g *Generator) NewLine() {
	g.buf.WriteByte(newLine)
}

// String returns buffer as string.
func (g *Generator) String() string {
	return g.buf.String()
}
