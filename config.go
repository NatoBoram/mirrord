package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

// UnmarshalConfig unmarshals a configuration.
func UnmarshalConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal turns a configuration into a JSON file.
func (c *Config) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

// Config is an instance of mirrord's configuration.
type Config struct {
	BeforeScript string `json:"before_script"`
	AfterScript  string `json:"after_script"`
}

func (c *Config) runBeforeScript() error {
	if c.BeforeScript == "" {
		return nil
	}

	cmd := exec.Command(c.BeforeScript)

	// Stream outputs to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}

func (c *Config) runAfterScript() error {
	if c.AfterScript == "" {
		return nil
	}

	cmd := exec.Command(c.AfterScript)

	// Stream outputs to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}
