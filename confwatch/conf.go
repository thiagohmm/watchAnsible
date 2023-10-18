package confwatch

import (
	"encoding/json"
	"log"
	"os"
)

const configFile = "/etc/monitor/config.json"

type AppConfig struct {
	AnsiblePlaybookComand string `json:"playbook_command"`
	WatchFolder           string `json:"watch_folder"`
	LogErrorFile          string `json:"error_File"`
	LogAnsibleFile        string `json:log_ansible_path`
	MapAnsibleFile        string `json:map_ansible_path`
	LogExtension          string `json:log_extension`
	ConfigurationMap      string `json:config_map`
}

func RetornaConf(cn string) (string, error) {

	var config AppConfig
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do arquivo de configuração: %v", err)

	}

	switch cn {
	case "AnsiblePlaybookComand":
		return config.AnsiblePlaybookComand, nil

	case "WatchFolder":
		return config.WatchFolder, nil

	case "LogErrorFile":
		return config.LogErrorFile, nil

	case "LogAnsibleFile":
		return config.LogAnsibleFile, nil

	case "MapAnsibleFile":
		return config.MapAnsibleFile, nil

	case "LogExtension":
		return config.LogExtension, nil

	case "ConfigurationMap":
		return config.ConfigurationMap, nil

	default:
		return "Error", os.ErrExist
	}
}
