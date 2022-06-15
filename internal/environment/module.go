package environment

type Modules map[string]Module

type Module struct {
	Type    string
	Source  string
	Version string
}
