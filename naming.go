package raymond

import "path"

var (
	namer Namer = NamerFunc(func(filepath string) string {
		// fileBase returns base file name
		// example: /foo/bar/baz.png => baz
		fileName := path.Base(filepath)
		fileExt := path.Ext(filepath)
		return fileName[:len(fileName)-len(fileExt)]
	})
)

// Namer make a new name from an old name.
type Namer interface {
	Naming(string) string
}

// NamerFunc wrap a func as Namer
type NamerFunc func(string) string

// Naming rename give name to a new name.
func (f NamerFunc) Naming(name string) string {
	return f(name)
}

// RegisterNamer register a new namer replace default.
// Note: this function is not concurrent safe.
func RegisterNamer(n Namer) {
	namer = n
}
