package environment

import (
	"strings"
)

type Environment struct {
	Type    string  `yaml:"type,omitempty"`
	Source  string  `yaml:"source,omitempty"`
	Version string  `yaml:"version,omitempty"`
	Modules Modules `yaml:"modules,omitempty"`
}

type Module struct {
	Type    string `yaml:"type,omitempty"`
	Source  string `yaml:"source,omitempty"`
	Version string `yaml:"version,omitempty"`
}

type Environments map[string]*Environment

type Modules map[string]*Module

func ModuleName(in string) string {
	fields := strings.FieldsFunc(in, func(c rune) bool {
		return (c == '/') || (c == '-')
	})
	return fields[len(fields)-1]
}

func (m Modules) Include(name string) (incl bool, key string) {
	key, incl = m.keyname(name)
	return incl, key
}

func (m Modules) keyname(in string) (string, bool) {
	name := ModuleName(in)
	for k := range m {
		if n := ModuleName(k); n == name {
			return k, true
		}
	}
	return "", false
}

func (m Modules) Set(in string, mod *Module) {
	if key, incl := m.keyname(in); incl {
		delete(m, key)
	}
	m[in] = mod
}
