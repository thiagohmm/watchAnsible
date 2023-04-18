package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// func fileFunc(arquivo string) {
// 	fmt.Println("removendo arquivo", arquivo)
// 	e := os.Remove(arquivo)
// 	if e != nil {
// 		log.Fatal(e)

// 	}
// }

func fileFunc(arquivo string) {

	fmt.Println("procurando o arquivo", arquivo)

	res := strings.ReplaceAll(arquivo, "/home/thiagohmm/watchTESTE/", "")
	fmt.Println(res)
	app := "find"

	arg0 := "/"
	arg1 := "-name"
	arg2 := res

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, _ := cmd.Output()

	fmt.Print(string(stdout))

}

// main
func main() {

	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event, ok := <-watcher.Events:
				//fmt.Printf("EVENT! %#v\n", event.Name)
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					//log.Println("modified file:", event.Name)
					go fileFunc(event.Name)
				}

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("/home/thiagohmm/watchTESTE"); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
