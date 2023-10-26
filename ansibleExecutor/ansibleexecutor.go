// ansibleexecutor/ansibleexecutor.go

package ansibleexecutor

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

type Ansible struct {
	inventoryFile string
	limit         string
}

func NewAnsible(inventoryFile, limit string) *Ansible {
	return &Ansible{
		inventoryFile: inventoryFile,
		limit:         limit,
	}
}

func (a *Ansible) RunPlaybook(playbookPath string, host string, path string, extension string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("ansible-playbook",
		"-i", a.inventoryFile,
		"--limit", a.limit,
		playbookPath)

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
		return errors.New("ansible-playbook failed")
	}

	return nil
}
