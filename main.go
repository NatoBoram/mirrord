package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	_, err := loadMirrors()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func loadMirrors() ([]Mirror, error) {

	// Get user config dir
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	mirrorConfigDir := configDir + string(os.PathSeparator) + name + string(os.PathSeparator) + "mirrors"

	// Create `mirror` folder
	err = os.MkdirAll(mirrorConfigDir, osPrivateDirectory)
	if err != nil {
		return nil, err
	}

	// Get all mirror configuration files
	var mirrors []Mirror
	err = filepath.Walk(mirrorConfigDir, func(path string, info os.FileInfo, err error) error {
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

		mirrors = append(mirrors, mirror)
		return nil
	})

	return mirrors, err
}
