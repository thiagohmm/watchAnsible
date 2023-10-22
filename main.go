package main

import (
	"errors"
	"fmt"

	"watch/confwatch"
	"watch/loggwatch"

	"github.com/fsnotify/fsnotify"
)

func fileFunc(arquivo string) error {

	fmt.Print("Recebendo o arquivo", arquivo)

	logger, err := loggwatch.SetupLogger()
	if err != nil {
		panic("Erro ao configurar o logger: " + err.Error())
	}
	defer logger.Sync()

	MapAnsibleFile, error := confwatch.RetornaConf("MapAnsibleFile")

	fmt.Println("Printando", MapAnsibleFile)
	if error == nil {

		fmt.Print("Printando2", MapAnsibleFile)
		//playbook, erro := mapconfiguration.FindValueForKey(MapAnsibleFile, arquivo)

		//ansibleexecutor.NewAnsible({})

	} else {
		logger.Error(arquivo + " Sem playbook ou host nao encontrado")
		return error
	}
	return nil
}

func main() {
	logger, err := loggwatch.SetupLogger()
	if err != nil {
		panic("Erro ao configurar o logger: " + err.Error())
	}
	defer logger.Sync()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)

	} else {
		fmt.Println("Watcher adicionado com sucesso")
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
				if event.Has(fsnotify.Create) {
					if err := fileFunc(event.Name); err != nil {
						fmt.Println("Erro:", err)

					}
				}

			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)

			}
		}
	}()

	var watchDirectory, _ = confwatch.RetornaConf("WatchFolder")
	if err == nil {
		err = watcher.Add(watchDirectory)
	}

	if err == nil {
		fmt.Println("Watcher adicionado com sucesso")

	} else {
		fmt.Println("ERROR", err)
		err1 := errors.New("Diretório " + watchDirectory + " inexistente")
		logger.Error(err1.Error())
		panic("Diretório inexistente")

	}

	<-done
}
