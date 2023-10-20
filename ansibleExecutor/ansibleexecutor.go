// ansibleexecutor/ansibleexecutor.go

package ansibleexecutor

import (
	"bytes"
	"os/exec"
)

type Ansible struct {
	inventoryFile string
	user          string
	privateKey    string
}

// 构造函数，构造 ansible 运行所需参数
func NewAnsible(inventoryFile, user, privateKey string) *Ansible {
	return &Ansible{
		inventoryFile: inventoryFile,
		user:          user,
		privateKey:    privateKey,
	}
}

// 运行接口
func (a *Ansible) RunPlaybook(playbookPath string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("ansible-playbook",
		"-i", a.inventoryFile,
		"-u", a.user,
		"--private-key="+a.privateKey,
		playbookPath)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// 返回结果
	return stdout.String() + stderr.String(), err
}
