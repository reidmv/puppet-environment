package environment

import (
	"strings"
)

// Corresponds to an r10k yaml environment definition
type Environment struct {
	Type    string  `yaml:"type,omitempty"`
	Source  string  `yaml:"source,omitempty"`
	Version string  `yaml:"version,omitempty"`
	Modules Modules `yaml:"modules,omitempty"`
}

// Corresponds to an r10k yaml module definition
type Module struct {
	Type    string `yaml:"type,omitempty"`
	Source  string `yaml:"source,omitempty"`
	Version string `yaml:"version,omitempty"`
}

type Environments map[string]*Environment

type Modules map[string]*Module

// Puppet modules may be listed as <author>/<name>, but the author component is optional.
// The ModuleName function returns the name component of any input string which may or may not
// include an author component. This is important because r10k often requires modules be
// manipulated by name, excluding author.
func ModuleName(in string) string {
	fields := strings.FieldsFunc(in, func(c rune) bool {
		return (c == '/') || (c == '-')
	})
	return fields[len(fields)-1]
}

// Returns true if the Modules map includes a given key name (exact).
func (m Modules) Include(name string) (incl bool, key string) {
	key, incl = m.keyname(name)
	return incl, key
}

// Given an input string which may be in the form <author>/<name>, returns the map key name of
// an entry with a matching <name> component. If no match is found, the second return value
// will be false.
func (m Modules) keyname(in string) (string, bool) {
	name := ModuleName(in)
	for k := range m {
		if n := ModuleName(k); n == name {
			return k, true
		}
	}
	return "", false
}

// Set the map value at the given input key to the given Module. If the map already contains
// an entry whose <name> component matches the given input's <name> component, delete it before
// adding the new entry. There should only ever be one module per <name> component in the map.
func (m *Modules) Set(in string, mod *Module) {
	if *m == nil {
		*m = make(map[string]*Module)
	}
	if key, incl := m.keyname(in); incl {
		delete(*m, key)
	}
	(*m)[in] = mod
}
