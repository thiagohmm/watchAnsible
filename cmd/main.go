package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"watch/ansibleExecutor"
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

	conf, err := confwatch.LoadConfig("../config.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Recebendo o arquivo", arquivo)

	logger, err := loggwatch.SetupLogger()

	playbook, error := mapconfiguration.FindValueForKey(conf.MapAnsibleFile, arquivo)
	if error != nil {
		logger.Error("Arquivo map ansible com error")
		removeFile(arquivo)
	}

	playbookPath := conf.AnsibleLocation + "/" + playbook
	fmt.Println(conf.MapAnsibleFile, conf.LogAnsibleFile, conf.AnsibleLocation, arquivo, conf.LogExtension, playbookPath)

	//var ansible = ansibleexecutor.Ansible{}

	resultPlaybook, err := ansibleExecutor.ExecutarPlaybookAnsible(playbookPath, arquivo)

	if err != nil {
		logger.Error("Erro ao executar o playbook ansible")
		removeFile(arquivo)
	}
	fmt.Println("Playbook executado com sucesso", resultPlaybook)
	logger.Info("Playbook executado com sucesso" + strings.Join(resultPlaybook, " "))

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
