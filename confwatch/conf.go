package confwatch

import (
	"github.com/spf13/viper"
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

func LoadConfig(path string) (*AppConfig, error) {
	var c AppConfig
	viper.SetConfigFile(path)
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("confwatch")
	viper.Unmarshal(&c)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &c, nil
}
