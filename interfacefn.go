package gen

import "github.com/elliotchance/orderedmap/v2"

type InterfaceFn struct {
	name    string
	params  *orderedmap.OrderedMap[string, string]
	returns *orderedmap.OrderedMap[string, string]

	g *Generator
}

func (f *InterfaceFn) Param(name string, sType string) *InterfaceFn {
	f.params.Set(name, sType)

	return f
}

func (f *InterfaceFn) Return(ss ...string) *InterfaceFn {
	if len(ss) == 1 {
		f.returns.Set(ss[0], "")
	} else if len(ss) == 2 {
		f.returns.Set(ss[0], ss[1])
	} else {
		panic("Returns must receive 1 or 2 parameters")
	}

	return f
}

func (f *InterfaceFn) Go() {
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
	f.g.NewLine()
}
