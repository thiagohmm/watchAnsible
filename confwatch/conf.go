package confwatch

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFile = "config.json"

type AppConfig struct {
	AnsiblePlaybookComand string `json:"playbook_command"`
	WatchFolder           string `json:"watch_folder"`
	LogErrorFile          string `json:"error_File"`
	LogAnsibleFile        string `json:"log_ansible_path"`
	MapAnsibleFile        string `json:"map_ansible_path"`
	LogExtension          string `json:"log_extension"`
	ConfigurationMap      string `json:"config_map"`
	AnsibleLocation       string `json:"ansible_location"`
}

func RetornaConf(cn string) (string, error) {

	var config AppConfig
	configFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do arquivo de configuração: %v", err)

	}

	switch cn {
	case "AnsiblePlaybookComand":
		if config.AnsiblePlaybookComand == "" {
			return "", fmt.Errorf("Chave Nula")
		}
		return config.AnsiblePlaybookComand, nil

	case "WatchFolder":
		if config.WatchFolder == "" {
			return "", fmt.Errorf("Chave Nula")
		}

		return config.WatchFolder, nil

	case "LogErrorFile":
		if config.LogErrorFile == "" {
			return "", fmt.Errorf("Chave Nula")
		}

		return config.LogErrorFile, nil

	case "LogAnsibleFile":
		if config.LogAnsibleFile == "" {
			return "", fmt.Errorf("Chave Nula")
		}
		return config.LogAnsibleFile, nil

	case "MapAnsibleFile":
		if config.MapAnsibleFile == "" {
			return "", fmt.Errorf("Chave Nula")
		}
		return config.MapAnsibleFile, nil

	case "LogExtension":
		if config.LogExtension == "" {
			return "", fmt.Errorf("Chave Nula")
		}
		return config.LogExtension, nil

	case "ConfigurationMap":
		if config.ConfigurationMap == "" {
			return "", fmt.Errorf("Chave Nula")
		}
		return config.ConfigurationMap, nil

	case "AnsibleLocation":
		if config.AnsibleLocation == "" {
			return "", fmt.Errorf("Chave Nula")
		}
		return config.AnsibleLocation, nil

	default:
		return "Error", os.ErrExist
	}
}

func CheckIfFileIsBlank(arquivo string) (string, error) {
	file, err := os.ReadFile(arquivo)
		
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
		return(string(file), nil);
	
}
