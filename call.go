package gen

type Call struct {
	name     string
	params   []string
	generics []string

	Calls []*Call

	g *Generator
}

func (c *Call) Param(name string) *Call {
	c.params = append(c.params, name)

	return c
}

func (c *Call) Call(indentifier ...string) *Call {
	newCall := &Call{
		name:  c.g.Id(indentifier...),
		Calls: c.Calls,
		g:     c.g,
	}

	newCall.Calls = append(newCall.Calls, newCall)

	return newCall
}

func (c *Call) Generic(name ...string) *Call {
	c.generics = append(c.generics, c.g.Id(name...))

	return c
}

func (c *Call) String() string {
	g := NewGenerator()

	for i, call := range c.Calls {
		if i > 0 {
			g.Lit(".")
		}

		g.Lit(call.name)

		if len(call.generics) > 0 {
			g.Lit("[")
			for i, gen := range call.generics {
				if i > 0 {
					g.Lit(",")
				}
				g.Lit(gen)
			}
			g.Lit("]")
		}

		g.Lit("(")
		for i, p := range call.params {
			if i > 0 {
				g.Lit(",")
			}
			g.Lit(p)
		}
		g.Lit(") ")
	}

	return g.String()
}

func (c *Call) Go() {
	for i, call := range c.Calls {
		if i > 0 {
			c.g.Lit(".")
		}
		c.g.Lit(call.name)

		if len(call.generics) > 0 {
			c.g.Lit("[")
			for i, gen := range call.generics {
				if i > 0 {
					c.g.Lit(",")
				}
				c.g.Lit(gen)
			}
			c.g.Lit("]")
		}

		c.g.Lit("(")
		for i, p := range call.params {
			if i > 0 {
				c.g.Lit(",")
			}
			c.g.Lit(p)
		}
		c.g.Lit(") ")
	}
	c.g.NewLine()
}
