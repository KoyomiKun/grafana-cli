package manager

import (
	"bytes"
	"encoding/json"

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

func (dm *DashboardManager) List() (grafana.Dashboards, error) {
	content, err := dm.client.GetApi("/api/search/", map[string]string{}, map[string]string{})
	if err != nil {
		log.Errorf("fail searching all dashboards: %v", err)
		return nil, err
	}
	resps := []map[string]interface{}{}
	err = json.Unmarshal(content, &resps)
	if err != nil {
		log.Errorf("fail unmarshaling response in searching dashboards: %v", err)
		return nil, err
	}

	dbs := make(grafana.Dashboards, 0, len(resps))
	for _, resp := range resps {
		if resp["type"].(string) != "dash-db" {
			continue
		}

		uid := resp["uid"].(string)
		db, err := dm.GetByUid(uid)
		if err != nil {
			log.Warnf("fail getting dashboard by uid %s: %v", uid, err)
			continue
		}
		if fuid, ok := resp["folderUid"]; ok {
			db.FolderUid = fuid.(string)
		}
		dbs = append(dbs, db)
	}

	return dbs, nil
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

func (dm *DashboardManager) Update(dashboard *grafana.Dashboard) error {

	// modify dashboard
	body, err := json.Marshal(dashboard)
	if err != nil {
		log.Errorf("Fail marshaling dashboard : %v", err)
		return err
	}
	_, err = dm.client.PostApi(
		"/api/dashboards/db",
		map[string]string{},
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Errorf("Fail post : %v", err)
		return err
	}
	return nil
}
