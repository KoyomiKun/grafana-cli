package grafana

import (
	"encoding/json"

	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Alerts []Alert
type Alert struct {
	Id             int                    `json:"id"`
	DashboardId    int                    `json:"dashboardId"`
	DashboardUid   string                 `json:"dashboardUid"`
	DashboardSlug  string                 `json:"dashboardSlug"`
	PanelId        int                    `json:"panelId"`
	Name           string                 `json:"name"`
	State          string                 `json:"state"`
	NewStateDate   string                 `json:"newStateDate"`
	EvalDate       string                 `json:"evalDate"`
	EvalData       map[string]interface{} `json:"evalData"`
	ExecutionError string                 `json:"executionError"`
	Url            string                 `json:"url"`
}

func NewAlerts(content []byte) (Alerts, error) {

	a := Alerts{}
	if err := json.Unmarshal(content, &a); err != nil {
		log.Errorf("Fail unmarshaling %s: %v", content, err)
		return nil, err
	}

	return a, nil
}
