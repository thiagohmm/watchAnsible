// ansibleexecutor/ansibleexecutor.go

package ansibleexecutor

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

type Command struct {
	command string
}

func NewCommand(cmd string) *Command {
	return &Command{
		command: cmd,
	}
}

func (a *Command) RunCMD(command string, host string, path string, extension string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	logfile := path + "/" + host + extension
	// Crie um arquivo para direcionar a saída
	outputFile, err := os.Create(logfile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Defina o arquivo como a saída do comando
	cmd.Stdout = outputFile

	err = cmd.Run()

	if err != nil {
		return err
	}

	if cmd.ProcessState.ExitCode() != 0 {
		return errors.New("ansible-playbook failed" + err.Error())
	}

	return nil
}
