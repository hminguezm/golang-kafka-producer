package entity

type HealthCheck struct {
	App     string `json:"name"`
	Version string `json:"version"`
	Env     string `json:"environment"`
	Author  string `json:"author"`
}
