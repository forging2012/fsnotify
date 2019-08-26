package main

import (
	"log"

	//"github.com/howeyc/fsnotify"
	"ffay/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	//err = watcher.Watch("testDir")
	err = watcher.Watch("./fsnotify/examples")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	//watcher.Close()
	defer watcher.Close()

}
