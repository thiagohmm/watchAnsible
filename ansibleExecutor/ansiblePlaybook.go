package ansibleexecutor

import (
	"context"
	"fmt"
	"watch/loggwatch"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"go.uber.org/zap"
)

func executarPlaybookAnsible(playbookPath string, limit string) (result []string, err error) {
	logger, err := loggwatch.SetupLogger()

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Limit:         limit,
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
	}

	cmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbookPath),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(cmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
	)

	err = exec.Execute(context.TODO())

	if err != nil {
		fmt.Println(err)
		logger.Error("Erro ao executar o playbook ansible", zap.Error(err))
	}

	result = append(result, cmd.String())
	return result, nil
}

// func main() {
// 	// Exemplo de uso da função
// 	err := executarPlaybookAnsible("site.yml", "ctr101")
// 	if err != nil {
// 			log.Fatal(err)
// 	}
// }
