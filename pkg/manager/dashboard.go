package manager

import (
	"github.com/KoyomiKun/grafana-cli/pkg/client"
	"github.com/KoyomiKun/grafana-cli/pkg/grafana"
	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type DashboardManager struct {
	client *client.Client
}

func NewDashboardManager(client *client.Client) *DashboardManager {
	return &DashboardManager{
		client: client,
	}
}

func (dm *DashboardManager) GetByUid(uid string) (*grafana.Dashboard, error) {
	content, err := dm.client.GetApi(
		"/api/dashboards/uid/"+uid,
		map[string]string{},
		map[string]string{})
	if err != nil {
		log.Errorf("Fail getting dashboard %s: %v", uid, err)
		return nil, err
	}
	dashboard, err := grafana.NewDashboard(content)
	if err != nil {
		log.Errorf("Fail creating dashbaord %s: %v", uid, err)
		return nil, err
	}
	return dashboard, err
}
