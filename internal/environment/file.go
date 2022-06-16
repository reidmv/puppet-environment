package environment

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type EnvironmentsFile struct {
	Path         string
	Environments Environments
}

func (f *EnvironmentsFile) Read() error {
	yamlString, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlString, &f.Environments)
	if err != nil {
		return err
	}
	return nil
}

func (f *EnvironmentsFile) Write() error {
	yamlData, err := yaml.Marshal(&f.Environments)
	if err != nil {
		return err
	}
	err = os.WriteFile(f.Path, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}
