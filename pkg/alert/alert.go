package alert

type Alert struct {
	AlertName string            `json:"alert_name"`
	Tags      map[string]string `json:"tags"`
}
