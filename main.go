package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	src := flag.String("src", "./dist/index.js", "index.js source")
	dest := flag.String("dest", "", "index.js destination")

	flag.Parse()

	if flag.NFlag() == 0 {
		// print usage
		fmt.Println("Usage: directus-extension-copy -src -dest")
		fmt.Println("\t-src  extension entrypoint source (optional). Default: ./dist/index.js")
		fmt.Println("\t-dest extension entrypoint destination (required)")
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
					fmt.Printf("🕐copying %s to %s\n", *src, *dest)
					_, err := copyFile(*src, *dest)
					if err != nil {
						return
					}
					fmt.Printf("✅copied %s to %s\n", *src, *dest)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	fmt.Printf("👁️ watching file `%s` for changes...\n", *src)

	err = watcher.Add(*src)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func copyFile(src, dest string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	return io.Copy(destFile, srcFile)
}
