package gen

import "github.com/elliotchance/orderedmap/v2"

type Method struct {
	name     string
	receiver string
	params   *orderedmap.OrderedMap[string, string]
	returns  *orderedmap.OrderedMap[string, string]

	g *Generator
}

func (m *Method) Receiver(ss ...string) *Method {
	for i, s := range ss {
		if i > 0 {
			m.receiver += " "
		}
		m.receiver += s
	}

	return m
}

func (m *Method) Param(name string, sType string) *Method {
	m.params.Set(name, sType)

	return m
}

func (m *Method) Returns(ss ...string) *Method {
	if len(ss) == 1 {
		m.returns.Set(ss[0], "")
	} else if len(ss) == 2 {
		m.returns.Set(ss[0], ss[1])
	} else {
		panic("Returns must receive 1 or 2 parameters")
	}

	return m
}

func (m *Method) Body(fn func()) {
	m.g.Lit("func ")

	if m.receiver != "" {
		m.g.Lit("(", m.receiver, ") ")
	}

	m.g.Lit(m.name, "(")
	i := 0
	for el := m.params.Front(); el != nil; el = el.Next() {
		if i > 0 {
			m.g.Lit(",")
		}
		m.g.Lit(el.Key, " ", el.Value)
		i++
	}
	m.g.Lit(") ")

	lenReturns := m.returns.Len()
	if lenReturns > 0 {
		if lenReturns > 1 {
			m.g.Lit("(")
		}
		i = 0
		for el := m.returns.Front(); el != nil; el = el.Next() {
			if i > 0 {
				m.g.Lit(",")
			}
			m.g.Lit(el.Key, " ", el.Value)
		}
		if lenReturns > 1 {
			m.g.Lit(")")
		}
	}

	m.g.Lit("{")
	m.g.NewLine()

	fn()

	m.g.Lit("}")
	m.g.NewLine()
	m.g.NewLine()
}
