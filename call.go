package gen

type Call struct {
	name   string
	params []string
}

func (f *Call) Param(name string) *Call {
	f.params = append(f.params, name)

	return f
}

func (c *Call) Go() string {
	g := NewGenerator()

	g.Lit(c.name, "(")
	for i, p := range c.params {
		if i > 0 {
			g.Lit(",")
		}
		g.Lit(p)
	}
	g.Lit(") ")

	g.NewLine()

	return g.String()
}
