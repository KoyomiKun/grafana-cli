package manager

import (
	"bytes"
	"encoding/json"

	"github.com/KoyomiKun/grafana-cli/pkg/client"
	"github.com/KoyomiKun/grafana-cli/pkg/config"
	"github.com/KoyomiKun/grafana-cli/pkg/grafana"
	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type TagManager struct {
	dbManager *DashboardManager
	alerts    []config.Alert
}

func NewTagManager(
	client *client.Client,
	alerts []config.Alert) *TagManager {

	return &TagManager{
		dbManager: NewDashboardManager(client),
		alerts:    alerts,
	}
}

func (tm *TagManager) Update() {

	client := tm.dbManager.client
	// group config.alerts by dashboard && panel
	dashboardUidToPanelIdToTags := make(map[string]map[int]map[string]string)

	for _, alert := range tm.alerts {

		content, err := client.GetApi(
			"/api/alerts",
			map[string]string{},
			map[string]string{
				"query": alert.AlertName,
			})
		if err != nil {
			log.Warnf("Fail get alerts: %v", err)
			continue
		}

		alerts, err := grafana.NewAlerts(content)
		if err != nil {
			log.Warnf("Fail create alerts: %v", err)
			continue
		}

		for _, a := range alerts {
			_, ok := dashboardUidToPanelIdToTags[a.DashboardUid]
			if !ok {
				dashboardUidToPanelIdToTags[a.DashboardUid] = make(map[int]map[string]string)
			}
			_, ok = dashboardUidToPanelIdToTags[a.DashboardUid][a.PanelId]
			if !ok {
				dashboardUidToPanelIdToTags[a.DashboardUid][a.PanelId] = make(map[string]string)
			}
			for k, v := range alert.Tags {
				dashboardUidToPanelIdToTags[a.DashboardUid][a.PanelId][k] = v
			}
		}
	}

	// modify dashboard
	for dashboardUid, panelIdToTags := range dashboardUidToPanelIdToTags {
		dashboard, err := tm.dbManager.GetByUid(dashboardUid)
		if err != nil {
			log.Warnf("Fail getting dashboard by uid %s: %v", dashboardUid, err)
			continue
		}

		dashboard.UpdateAlert(panelIdToTags)

		body, err := json.Marshal(dashboard)
		if err != nil {
			log.Warnf("Fail marshaling dashboard %s: %v", dashboardUid, err)
			continue
		}
		_, err = client.PostApi(
			"/api/dashboards/db",
			map[string]string{},
			bytes.NewBuffer(body),
		)
		if err != nil {
			log.Warnf("Fail post %s: %v", dashboardUid, err)
			continue
		}
	}
}
