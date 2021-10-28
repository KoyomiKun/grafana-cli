package grafana

import (
	"reflect"

	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Targets []Target
type Target struct {
	Expr         string `map:"expr"`
	LegendFormat string `map:"legendFormat"`
	RefId        string `map:"refId"`
}

func NewTargets(targets []map[string]string) Targets {

	metrics := make(Targets, 0, len(targets))

	for _, target := range targets {
		metric := Target{}
		ty := reflect.TypeOf(metric).Elem()
		value := reflect.ValueOf(metric).Elem()
		for i := 0; i < ty.NumField(); i++ {
			fieldName := ty.Field(i).Tag.Get("map")
			if v, ok := target[fieldName]; ok {
				value.Field(i).SetString(v)
			} else {
				log.Warnf("Target lack of field %s", fieldName)
			}
		}
		metrics = append(metrics, metric)
	}

	return metrics
}
