package gen

import "github.com/elliotchance/orderedmap/v2"

type Func struct {
	name    string
	params  *orderedmap.OrderedMap[string, string]
	returns *orderedmap.OrderedMap[string, string]

	g *Generator
}

func (f *Func) Param(name string, sType string) *Func {
	f.params.Set(name, sType)

	return f
}

func (f *Func) Returns(ss ...string) *Func {
	if len(ss) == 1 {
		f.returns.Set(ss[0], "")
	} else if len(ss) == 2 {
		f.returns.Set(ss[0], ss[1])
	} else {
		panic("Returns must receive 1 or 2 parameters")
	}

	return f
}

func (f *Func) String(fn func()) string {
	g := NewGenerator()

	g.Lit("func ")

	g.Lit(f.name, "(")
	i := 0
	for el := f.params.Front(); el != nil; el = el.Next() {
		if i > 0 {
			g.Lit(",")
		}
		g.Lit(el.Key, " ", el.Value)
		i++
	}
	g.Lit(") ")

	lenReturns := f.returns.Len()
	if lenReturns > 0 {
		if lenReturns > 1 {
			g.Lit("(")
		}
		i = 0
		for el := f.returns.Front(); el != nil; el = el.Next() {
			if i > 0 {
				g.Lit(",")
			}
			g.Lit(el.Key, " ", el.Value)
			i++
		}
		if lenReturns > 1 {
			g.Lit(")")
		}
	}

	g.Lit("{")
	g.NewLine()

	fn()

	g.Lit("}")
	return g.String()
}

func (f *Func) Body(fn func()) {
	f.g.Lit(f.String(fn))
	f.g.NewLine()
	f.g.NewLine()
}
