package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func main() {

	for {

		// Reload configurations at each loop
		config, err := readConfig()
		if err != nil {
			log.Fatalln(err.Error())
		}

		mirrors, err := readMirrors()
		if err != nil {
			log.Fatalln(err.Error())
		}

		// Exit if there's no mirrors
		if len(mirrors) == 0 {
			log.Fatalln("No mirrors were configured.")
		}

		err = config.runBeforeScript()
		if err != nil {
			fmt.Println(err.Error())
		}

		// Create queue
		mirrorchan := make(chan *Mirror, runtime.NumCPU())
		go func() {
			for _, mirror := range mirrors {
				mirrorchan <- &mirror
			}
			close(mirrorchan)
		}()

		// Mirror asynchronously
		var wg sync.WaitGroup
		wg.Add(len(mirrors))
		for c := 1; c <= runtime.NumCPU(); c++ {
			go func() {
				for mirror := range mirrorchan {
					err := mirror.cycle()
					if err != nil {
						fmt.Println(err.Error())
					}

					wg.Done()
				}
			}()
		}

		// Minimum time between cycles
		var wgSleep sync.WaitGroup
		wgSleep.Add(1)
		go func() {
			time.Sleep(time.Hour)
			wgSleep.Done()
		}()

		wg.Wait()

		fmt.Println("Running garbage collection...")
		cmd := exec.Command("ipfs", "repo", "gc", "--stream-errors")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		err = config.runAfterScript()
		if err != nil {
			fmt.Println(err.Error())
		}

		wgSleep.Wait()
	}
}
