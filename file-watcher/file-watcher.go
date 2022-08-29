package filewatcher

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func WatchDir(dir, fileToWatch string, c chan any) {
	watcher, err := fsnotify.NewWatcher()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					if strings.Contains(event.Name, fileToWatch) {
						c <- struct{}{}
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

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	for true {

	}
}
