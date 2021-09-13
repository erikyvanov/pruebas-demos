package models

type Event struct {
	EntityName string                 `json:"entity_name"`
	EventName  string                 `json:"event_name"`
	Data       map[string]interface{} `json:"data"`
}
