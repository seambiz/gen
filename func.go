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

func (f *Func) Body(fn func()) {
	f.g.Lit("func ")

	f.g.Lit(f.name, "(")
	i := 0
	for el := f.params.Front(); el != nil; el = el.Next() {
		if i > 0 {
			f.g.Lit(",")
		}
		f.g.Lit(el.Key, " ", el.Value)
	}
	f.g.Lit(") ")

	lenReturns := f.returns.Len()
	if lenReturns > 0 {
		if lenReturns > 1 {
			f.g.Lit("(")
		}
		i = 0
		for el := f.returns.Front(); el != nil; el = el.Next() {
			if i > 0 {
				f.g.Lit(",")
			}
			f.g.Lit(el.Key, " ", el.Value)
		}
		if lenReturns > 1 {
			f.g.Lit(")")
		}
	}

	f.g.Lit("{")
	f.g.NewLine()

	fn()

	f.g.Lit("}")
	f.g.NewLine()
	f.g.NewLine()
}
