package gen

import "github.com/elliotchance/orderedmap/v2"

type Field struct {
	name        string
	isPtr       bool
	isSlice     bool
	typePointer bool
	typeName    string
	tags        *orderedmap.OrderedMap[string, string]

	g *Generator
}

func (f *Field) Go() {
	f.g.Lit(f.name)

	if f.isPtr {
		f.g.Lit(" *")
	} else {
		f.g.Lit(" ")
	}

	if f.isSlice {
		f.g.Lit("[]")
	}

	if f.typePointer {
		f.g.Lit("*")
	}

	f.g.Lit(f.typeName)

	if f.tags.Len() > 0 {
		f.g.Lit(" `")

		i := 0
		for el := f.tags.Front(); el != nil; el = el.Next() {

			if i > 0 {
				f.g.Lit(" ")
			}
			f.g.Lit(el.Key, `:"`, el.Value, `"`)
			i++
		}

		f.g.Lit("`")
	}
	f.g.NewLine()
}

func (f *Field) Tag(key string, value string) *Field {
	f.tags.Set(key, value)

	return f
}

func (f *Field) Ptr() *Field {
	f.isPtr = true

	return f
}

func (f *Field) Slice(b bool) *Field {
	f.isSlice = b

	return f
}

func (f *Field) Type(name string) *Field {
	f.typeName = name

	return f
}

func (f *Field) PointerType(name string) *Field {
	f.typePointer = true
	f.typeName = name

	return f
}
