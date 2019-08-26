package main

// https://github.com/howeyc/fsnotify

import (
	"log"
	"time"

	//"github.com/howeyc/fsnotify"
	"ffay/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

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

	for {
		//err = watcher.Watch("/tmp/")
		//err = watcher.Watch("./")
		err = watcher.Watch("./fsnotify/examples")
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second)
	}

}

//2019/08/21 22:59:39 event: "1": CREATE
//2019/08/21 22:59:50 event: "1": DELETE
//2019/08/21 22:59:55 event: "1": CREATE
//2019/08/21 22:59:56 event: "1": MODIFY|ATTRIB
//2019/08/21 23:00:00 event: "1": MODIFY|ATTRIB

//2019/08/21 23:01:18 event: "fsnotify/examples/1": CREATE
//2019/08/21 23:01:24 event: "fsnotify/examples/1": MODIFY|ATTRIB
//2019/08/21 23:01:24 event: "fsnotify/examples/1": MODIFY
//2019/08/21 23:01:25 event: "fsnotify/examples/1": MODIFY|ATTRIB
//2019/08/21 23:01:28 event: "fsnotify/examples/1": DELETE
//2019/08/21 23:01:35 event: "fsnotify/examples/dir": CREATE
//2019/08/21 23:01:40 event: "fsnotify/examples/dir": DELETE
