package r10k

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

// Path to look at for a pre-existing r10k.yaml file
var configPath string

// Sets what path the package will look for r10k.yaml at
func SetConfigPath(path string) {
	configPath = path
}

func DeployEnvironment(name, config, dir string) error {
	r10kYaml, err := createTempConfig(config, dir)
	if err != nil {
		return err
	}
	defer func() {
		r10kYaml.Close()
		os.Remove(r10kYaml.Name())
	}()

	err = run("r10k", "--config", r10kYaml.Name(), "deploy", "environment", name, "-v", "-m")
	if err != nil {
		return err
	}
	return nil
}

func DeployModule(environment, name, config, dir string) error {
	r10kYaml, err := createTempConfig(config, dir)
	if err != nil {
		return err
	}
	defer func() {
		r10kYaml.Close()
		os.Remove(r10kYaml.Name())
	}()

	err = run("r10k", "--config", r10kYaml.Name(), "deploy", "module", name, "--environment", environment, "-v")
	if err != nil {
		return err
	}
	return nil
}

// Create and write an r10k.yaml file to a new temporary location.
// Return the file object.
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

// Run a command, stream the output to stdout
func run(arg0 string, args ...string) error {
	cmd := exec.Command(arg0, args...)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout

	fmt.Println("Running command:", arg0, strings.Join(args, " "))
	err = cmd.Start()
	if err != nil {
		return err
	}

	// Stream the output from r10k as it is generated
	scanner := bufio.NewScanner(pipe)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

// Return a string containing r10k.yaml content appropriate for use in invoking r10k.
// The content generated will:
//   - Be based on an existing r10k.yaml file, if found
//   - Overwrite the sources.puppet-environment config from a base file, if present
//   - Include sources.puppet-environment configured with the given envYaml and envDir values
func r10kYamlStr(envYaml, envDir string) (string, error) {
	var config map[string]interface{}

	yamlSource := map[string]string{
		"type":    "yaml",
		"config":  envYaml,
		"basedir": envDir,
	}

	// If no pre-existing r10k.yaml file is found, generate and return just the yaml we need.
	// No need to read in and merge any additional config.
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

	// Read in the pre-existing r10k.yaml file and unmarshal it.
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return "", err
	}

	// Set or merge in the appropriate sources.puppet-environment value
	if srcs, ok := config["sources"]; !ok {
		config["sources"] = map[string]interface{}{
			"puppet-environment": &yamlSource,
		}
	} else {
		srcs.(map[string]interface{})["puppet-environment"] = &yamlSource
	}

	// Marshal the full result back to a yaml string to return
	if bytes, err = yaml.Marshal(config); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
