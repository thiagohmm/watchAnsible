// ansibleexecutor/ansibleexecutor.go

package ansibleexecutor

import (
	"bytes"
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

// 运行接口
func (a *Ansible) RunPlaybook(playbookPath string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("ansible-playbook",
		"-i", a.inventoryFile,
		"--limit "+a.limit,
		playbookPath)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// 返回结果
	return stdout.String() + stderr.String(), err
}
