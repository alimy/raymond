package raymond

import (
	"path"
	"testing"
)

func TestNaming(t *testing.T) {
	for origin, expect := range map[string]string{
		"templates/a.tmpl":   "a",
		"templates/b/c.tmpl": "c",
		"templates/d/e.tmpl": "e",
		"templates/d/f.tmpl": "f",
	} {
		if name := namer.Naming(origin); name != expect {
			t.Errorf("expect: %s got %s", expect, name)
		}
	}
}

func TestRegisterNamer(t *testing.T) {
	results := make(map[string]string)
	oldNamer := namer
	RegisterNamer(NamerFunc(func(filepath string) string {
		results[oldNamer.Naming(filepath)] = partialName(filepath)
		return oldNamer.Naming(filepath)
	}))
	for origin, expect := range map[string]string{
		"templates/a.tmpl":   "templates/a",
		"templates/b/c.tmpl": "templates/b/c",
		"templates/d/e.tmpl": "templates/d/e",
		"templates/d/f.tmpl": "templates/d/f",
	} {
		name := namer.Naming(origin)
		if value, exist := results[name]; exist && value != expect {
			t.Errorf("expect: %s got %s", expect, value)
		}
	}
}

func partialName(filepath string) string {
	ext := path.Ext(filepath)
	return filepath[:len(filepath)-len(ext)]
}
