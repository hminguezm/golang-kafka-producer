package event

type MessageBase struct {
	EventId         string                 `json:"eventId"`
	EventName       string                 `json:"eventName"`
	EventDataFormat string                 `json:"eventDataFormat"`
	Type            string                 `json:"type"`
	Timestamp       string                 `json:"timestamp"`
	Version         string                 `json:"version"`
	Country         string                 `json:"country"`
	Origin          string                 `json:"origin"`
	Payload         map[string]interface{} `json:"payload"`
}
