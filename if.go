package gen

type IfStmt struct {
	g *Generator
}

func (i *IfStmt) Body(fn func()) {
	fn()

	i.g.Lit("}")
}
