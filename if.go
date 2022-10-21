package gen

import "strings"

type IfStmt struct {
	g *Generator
}

func (i *IfStmt) Body(fn func()) *IfStmt {
	fn()

	i.g.Lit("}")
	i.g.NewLine()

	return i
}

func (i *IfStmt) Else(fn func()) {
	// TODO cleanup
	buf := i.g.String()
	buf = strings.TrimSuffix(buf, "\n")
	i.g.buf.Reset()
	i.g.buf.WriteString(buf)

	i.g.Go(" else {")

	fn()

	i.g.Lit("}")
	i.g.NewLine()
}
