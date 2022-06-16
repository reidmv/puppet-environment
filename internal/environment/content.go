package environment

type Environment struct {
	Type    string
	Source  string
	Version string
	Modules Modules
}

type Module struct {
	Type    string
	Source  string
	Version string
}

type Environments map[string]*Environment

type Modules map[string]*Module
