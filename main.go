package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/KoyomiKun/grafana-cli/pkg/alert"
	"github.com/KoyomiKun/grafana-cli/pkg/client"
	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Config struct {
	APIKey  string        `json:"api_key,omitempty"`
	BaseUrl string        `json:"base_url,omitempty"`
	Alerts  []alert.Alert `json:"alerts"`
}

var (
	cfgPath string
)

func main() {
	currentDir, _ := os.Getwd()
	app := &cli.App{
		Name:        "grafana-cli",
		Description: "grafana cmd client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       path.Join(currentDir, "/cfg/config.json"),
				Destination: &cfgPath,
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "alert",
				Description: "alert management.",
				Subcommands: []*cli.Command{
					{
						Name:        "tags",
						Description: "Update tags.",
						Action: func(c *cli.Context) error {
							return startTag()
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func startTag() error {

	// init config
	config := &Config{}
	configFile, err := os.Open(cfgPath)
	if err != nil {
		log.Errorf("Fail openning config file %s: %v\n", cfgPath, err)
		return err
	}
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Errorf("Fail reading config file %s: %v\n", cfgPath, err)
		return err
	}
	if err := json.Unmarshal(configBytes, config); err != nil {
		log.Errorf("Fail unmarshaling config file %s: %v\n", cfgPath, err)
		return err
	}

	client := client.NewClient(
		2*time.Second,
		config.BaseUrl,
		config.APIKey,
	)

	tagManager := alert.NewTagManager(
		client,
		config.Alerts,
	)

	tagManager.Update()
	return nil

}
