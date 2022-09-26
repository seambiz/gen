package gen

func Sprintf(fmt string, ss ...string) string {
	g := NewGenerator()

	g.Lit("fmt.Sprintf(", `"`, fmt, `",`)
	g.Lit(ss...)
	g.Lit(")")

	return g.String()
}

func Ptr(ss ...string) string {
	g := NewGenerator()

	g.Lit("*")
	g.Lit(ss...)

	return g.String()
}

func Slice(ss ...string) string {
	g := NewGenerator()

	g.Lit("[]")
	g.Lit(ss...)

	return g.String()
}

func PtrIf(b bool, ss ...string) string {
	g := NewGenerator()
	if b {
		g.Lit("*")
	}
	g.Lit(ss...)

	return g.String()
}

// NewStruct generates the code for struct initialization.
// It uses a map, so the template looks closer to go code.
func NewStruct(name string, kv *KV) string {
	g := NewGenerator()

	g.Lit(name, "{")
	g.NewLine()

	for key, value := range *kv {
		g.Lit(key, ":", value, ",")
		g.NewLine()
	}
	g.Lit(name, "}")
	g.NewLine()

	return g.String()
}
