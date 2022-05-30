package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	src := flag.String("src", "./dist/index.js", "index.js source")
	onChange := flag.String("onChange", "", "command to be executed after file change")

	flag.Parse()

	if flag.NFlag() == 0 {
		// print usage
		fmt.Println("Usage: wacp -src -dest")
		fmt.Println("\t-src  extension entrypoint source (optional). Default: ./dist/index.js")
		fmt.Println("\t-onChange command to be executed after file change")
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					args := strings.Split(*onChange, " ")
					cmd := exec.Command(args[0], args[1:]...)
					cmd.Stdout = os.Stdout
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	log.Println("watching file for changes...")

	err = watcher.Add(*src)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
