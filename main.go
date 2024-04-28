package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"watch/confwatch"
	"watch/loggwatch"
	"watch/mapconfiguration"

	"github.com/fsnotify/fsnotify"
)

func removeFile(arquivo string) {
	e := os.Remove(arquivo)
	if e != nil {
		log.Fatal(e)
	}
}

func fileFunc(arquivo string) error {

	fmt.Println("Recebendo o arquivo", arquivo)

	checkIfFIleisBlank, err := confwatch.CheckIfFileIsBlank(arquivo)
	if err != nil {
		log.Fatal(err)
	}

	if checkIfFIleisBlank == "" {

		logger, err := loggwatch.SetupLogger()
		if err != nil {
			panic("Erro ao configurar o logger: " + err.Error())
		}
		defer logger.Sync()

		MapAnsibleFile, error := confwatch.RetornaConf("MapAnsibleFile")

		if error != nil {
			logger.Error("MapAnsibleFile chave de configuraçao não encontrado")
			removeFile(arquivo)

		}

		log_ansible_path, error := confwatch.RetornaConf("LogAnsibleFile")
		if error != nil {
			logger.Error("log_ansible_path chave de configuraçao não encontrado")
			removeFile(arquivo)

		}

		ansible_location, error := confwatch.RetornaConf("AnsibleLocation")
		if error != nil {
			logger.Error("ansible_location chave de configuraçao não encontrado")
			removeFile(arquivo)

		}

		var watchDirectory, _ = confwatch.RetornaConf("WatchFolder")
		if err != nil {
			logger.Error("WatchFolder chave de configuraçao não encontrado")
			removeFile(arquivo)
		}

		host := strings.Split(arquivo, watchDirectory)[1]
		fmt.Println("printando o host", host)

		log_extension, error := confwatch.RetornaConf("LogExtension")
		if error != nil {
			logger.Error("ansible_location chave de configuraçao não encontrado")
			removeFile(arquivo)

		}

		playbook, error := mapconfiguration.FindValueForKey(MapAnsibleFile, host)
		if error != nil {
			logger.Error("Arquivo map ansible com error")
			removeFile(arquivo)

		}
		playbookPath := ansible_location + "/" + playbook
		fmt.Println(MapAnsibleFile, log_ansible_path, ansible_location, host, log_extension, playbookPath)

		return nil
	} else {
		return errors.New("Arquivo em branco")
	}
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
