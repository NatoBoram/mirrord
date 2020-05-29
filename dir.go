package main

import "os"

func configDir() (string, error) {

	// Get user config dir
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	dir := userConfigDir + string(os.PathSeparator) + "mirrord"

	// Create directory
	err = os.MkdirAll(dir, osPrivateDirectory)
	return dir, err
}

func mirrorDir() (string, error) {

	// Get config dir
	configDir, err := configDir()
	if err != nil {
		return "", err
	}

	dir := configDir + string(os.PathSeparator) + "mirrors"

	// Create directory
	err = os.MkdirAll(dir, osPrivateDirectory)
	return dir, err
}
