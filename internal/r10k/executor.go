package r10k

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

var confTemplate = `---
sources:
  puppet-environment:
    type: yaml
    config: %s
    basedir: %s
`

var configPath string

func SetConfigPath(path string) {
	configPath = path
}

func DeployEnvironment(name, config, dir string) (ret error) {
	r10kYaml, err := createTempConfig(config, dir)
	if err != nil {
		return err
	}
	defer func() {
		ret = closeTempConfig(r10kYaml)
	}()

	err = run("r10k", "--config", r10kYaml.Name(), "deploy", "environment", name, "-v", "-m")
	if err != nil {
		return err
	}
	return nil
}

func DeployModule(environment, name, config, dir string) (ret error) {
	r10kYaml, err := createTempConfig(config, dir)
	if err != nil {
		return err
	}
	defer func() {
		ret = closeTempConfig(r10kYaml)
	}()

	err = run("r10k", "--config", r10kYaml.Name(), "deploy", "module", name, "--environment", environment, "-v")
	if err != nil {
		return err
	}
	return nil
}

func closeTempConfig(f *os.File) (err error) {
	if err = f.Close(); err != nil {
		return err
	}
	if err = os.Remove(f.Name()); err != nil {
		return err
	}
	return nil
}

func createTempConfig(config, dir string) (*os.File, error) {
	f, err := os.CreateTemp("", "puppet-environment")
	if err != nil {
		return f, err
	}
	configStr, err := r10kYamlStr(config, dir)
	if err != nil {
		return f, err
	}
	if _, err = f.Write([]byte(configStr)); err != nil {
		return f, err
	}
	return f, nil
}

func run(arg0 string, args ...string) error {
	cmd := exec.Command(arg0, args...)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(pipe)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
	return nil
}

func r10kYamlStr(envYaml, envDir string) (string, error) {
	var config map[string]interface{}

	yamlSource := map[string]string{
		"type":    "yaml",
		"config":  envYaml,
		"basedir": envDir,
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		config = map[string]interface{}{
			"sources": map[string]interface{}{
				"puppet-environment": &yamlSource,
			},
		}
		if bytes, err := yaml.Marshal(config); err != nil {
			return "", err
		} else {
			return string(bytes), nil
		}
	}

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return "", err
	}

	if srcs, ok := config["sources"]; !ok {
		config["sources"] = map[string]interface{}{
			"puppet-environment": &yamlSource,
		}
	} else {
		srcs.(map[string]interface{})["puppet-environment"] = &yamlSource
	}

	if bytes, err = yaml.Marshal(config); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
