package environment

type Environments map[string]Environment

type Environment struct {
	Type    string
	Source  string
	Version string
	Modules map[string]Module
}
