package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"watch/confwatch"
	"watch/loggwatch"

	"github.com/fsnotify/fsnotify"
)

func fileFunc(arquivo string) error {
	logger, err := loggwatch.SetupLogger()
	if err != nil {
		panic("Erro ao configurar o logger: " + err.Error())
	}
	defer logger.Sync()

	fmt.Println("Procurando o arquivo", arquivo)

	// Verifica se o arquivo existe
	if _, err := os.Stat(arquivo); os.IsNotExist(err) {
		err1 := errors.New("Arquivo não encontrado:" + err.Error())
		logger.Error(err1.Error())
		return fmt.Errorf("Arquivo não encontrado: %v", err)
	}

	res := strings.ReplaceAll(arquivo, "/home/thiagohmm/watchTESTE/", "")
	fmt.Println(res)
	app := "find"

	arg0 := "/"
	arg1 := "-name"
	arg2 := res

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()
	if err != nil {
		err1 := errors.New("Erro ao executar o comando")
		logger.Error(err1.Error())
		return fmt.Errorf("Erro ao executar o comando 'find': %v", err)
	}

	fmt.Print(string(stdout))

	// Obter a data atual
	dataAtual := time.Now().Format("2006-01-02")
	outputFileName := fmt.Sprintf("/tmp/%s_%s.txt", res, dataAtual)

	// Escrever a saída no arquivo
	file, err := os.Create(outputFileName)
	if err != nil {
		err1 := errors.New("Erro ao criar o arquivo de saída")
		logger.Error(err1.Error())
		return fmt.Errorf("Erro ao criar o arquivo de saída: %v", err)

	}
	defer file.Close()

	_, err = file.WriteString(string(stdout))
	if err != nil {
		err1 := errors.New("Erro ao criar o arquivo de saída")
		logger.Error(err1.Error())
		return fmt.Errorf("Erro ao escrever no arquivo de saída: %v", err)
	}

	fmt.Printf("Saída salva em %s\n", outputFileName)

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
		err1 := errors.New("Diretório inexistente")
		logger.Error(err1.Error())
		panic("Diretório inexistente")

	}

	<-done
}
