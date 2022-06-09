package internal

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Environments struct {
	Path string
	Data map[interface{}]interface{}
}

func (e *Environments) LoadYaml() error {
	yamlString, err := ioutil.ReadFile(e.Path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlString, &e.Data)
	if err != nil {
		return err
	}
	return nil
}
