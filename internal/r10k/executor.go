package r10k

import (
	"bufio"
	"fmt"
	"log"
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
	r10kYaml, err := os.CreateTemp("", "puppet-environment")
	if err != nil {
		return err
	}
	defer func() {
		if err = r10kYaml.Close(); err != nil {
			ret = err
		}
	}()

	if _, err = r10kYaml.Write([]byte(r10kYamlStr(config, dir))); err != nil {
		return err
	}

	cmd := exec.Command("echo", "r10k", "--config", r10kYaml.Name(), "deploy", "environment", name, "-m")
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

func DeployModule(environment, name, config, dir string) error {
	log.Fatal("Not implemented")
	return nil
}

func r10kYamlStr(config, dir string) string {
	return fmt.Sprintf(confTemplate, config, dir)
}
