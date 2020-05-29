package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readConfig() (*Config, error) {

	dir, err := configDir()
	if err != nil {
		return nil, err
	}

	path := dir + string(os.PathSeparator) + "config.json"

	// Read configuration file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Bytes->JSON
	config, err := UnmarshalConfig(bytes)
	if err != nil {
		return nil, err
	}

	return &config, err
}

func readMirrors() ([]Mirror, error) {

	dir, err := mirrorDir()
	if err != nil {
		return nil, err
	}

	// Get all mirror configuration files
	var mirrors []Mirror
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		} else if !strings.HasSuffix(info.Name(), ".json") {
			return ErrNotJSON
		}

		// Read mirror configuration file
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Bytes->JSON
		mirror, err := UnmarshalMirror(bytes)
		if err != nil {
			return err
		}

		mirror.path = path
		mirror.info = info

		mirrors = append(mirrors, mirror)
		return nil
	})

	return mirrors, err
}
