package r10k

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var confTemplate = `---
sources:
  puppet-environment:
    type: yaml
    config: %s
    basedir: %s
`

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
	if _, err = f.Write([]byte(r10kYamlStr(config, dir))); err != nil {
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

func r10kYamlStr(config, dir string) string {
	return fmt.Sprintf(confTemplate, config, dir)
}
