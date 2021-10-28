package manager

import "github.com/KoyomiKun/grafana-cli/pkg/config"

type TargetManager struct {
	dbManager *DashboardManager
	targets   []config.Target
}

func NewTargetManager(
	dbManager *DashboardManager,
	targets []config.Target) {
}
