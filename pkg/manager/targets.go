package manager

import (
	"github.com/KoyomiKun/grafana-cli/pkg/config"
	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type TargetManager struct {
	dbManager *DashboardManager
	targets   []config.Target
}

func NewTargetManager(
	dbManager *DashboardManager,
	targets []config.Target) *TargetManager {

	return &TargetManager{
		dbManager,
		targets,
	}
}

func (tm *TargetManager) Update() {

	for _, target := range tm.targets {
		dashboard, err := tm.dbManager.GetByUid(target.DashboardUid)
		if err != nil {
			log.Warnf("Fail getting dashboard by uid %s: %v", target.DashboardUid, err)
			continue
		}

		dashboard.UpdateTargets(target.Panels)

		tm.dbManager.Update(dashboard)
	}

}
