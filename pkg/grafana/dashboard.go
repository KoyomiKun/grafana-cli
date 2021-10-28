package grafana

import (
	"encoding/json"

	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Dashboard struct {
	Meta      map[string]interface{} `json:"meta"`
	Dashboard map[string]interface{} `json:"dashboard"`
}

func NewDashboard(content []byte) (*Dashboard, error) {
	d := Dashboard{}
	if err := json.Unmarshal(content, &d); err != nil {
		log.Errorf("Fail unmarshaling %s: %v", content, err)
		return nil, err
	}
	return &d, nil
}

func (d *Dashboard) UpdateAlert(panelIdToTags map[int]map[string]string) {
	panels := d.Dashboard["panels"].([]interface{})
	for i := 0; i < len(panels); i++ {
		panel := panels[i].(map[string]interface{})
		if newTags, ok := panelIdToTags[int(panel["id"].(float64))]; ok {
			alert := panel["alert"].(map[string]interface{})
			tags := alert["alertRuleTags"].(map[string]interface{})
			for k, v := range newTags {
				if v == "" {
					delete(tags, k)
				} else {
					tags[k] = v
				}
			}
		}
	}
}

func (d *Dashboard) UpdateMetrics(panelIdToTags map[int]map[string]string) {
	panels := d.Dashboard["panels"].([]interface{})
	for i := 0; i < len(panels); i++ {
		panel := panels[i].(map[string]interface{})
		if newTags, ok := panelIdToTags[int(panel["id"].(float64))]; ok {
			alert := panel["alert"].(map[string]interface{})
			tags := alert["alertRuleTags"].(map[string]interface{})
			for k, v := range newTags {
				if v == "" {
					delete(tags, k)
				} else {
					tags[k] = v
				}
			}
		}
	}
}
