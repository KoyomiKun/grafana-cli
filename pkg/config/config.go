package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Alert struct {
	AlertName string            `json:"alert_name"`
	Tags      map[string]string `json:"tags"`
}

type Target struct {
	RefId string `json:"ref_id,omitempty"`
	Expr  string `json:"expr,omitempty"`
}

type Config struct {
	APIKey  string   `json:"api_key"`
	BaseUrl string   `json:"base_url"`
	Targets []Target `json:"targets,omitempty"`
	Alerts  []Alert  `json:"alerts,omitempty"`
}

func NewConfig(cfgPath string) (*Config, error) {
	config := &Config{}
	configFile, err := os.Open(cfgPath)
	if err != nil {
		log.Errorf("Fail openning config file %s: %v\n", cfgPath, err)
		return nil, err
	}
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Errorf("Fail reading config file %s: %v\n", cfgPath, err)
		return nil, err
	}
	if err := json.Unmarshal(configBytes, config); err != nil {
		log.Errorf("Fail unmarshaling config file %s: %v\n", cfgPath, err)
		return nil, err
	}
	return nil, err
}
