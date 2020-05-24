package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func main() {

	for {

		// Load mirrors
		mirrors, err := loadMirrors()
		if err != nil {
			log.Fatalln(err.Error())
		}

		// Exit if there's no mirrors
		if len(mirrors) == 0 {
			log.Fatal("No mirrors were configured.")
		}

		// Create queue
		mirrorchan := make(chan *Mirror, runtime.NumCPU())
		go func() {
			for _, mirror := range mirrors {
				mirrorchan <- &mirror
			}
			close(mirrorchan)
		}()

		var wg sync.WaitGroup
		wg.Add(len(mirrors))

		// Mirror asynchronously
		for c := 1; c <= runtime.NumCPU(); c++ {
			go func() {
				for mirror := range mirrorchan {
					err := mirror.cycle()
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}()
		}

		wg.Wait()
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

		mirror.path = path
		mirror.info = info

		mirrors = append(mirrors, mirror)
		return nil
	})

	return mirrors, err
}
