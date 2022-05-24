package event_store

type EventStore struct {
	ID           string `json:"id"`
	Aggregate    string `json:"aggregate"`
	AggregateId  string `json:"aggregateId"`
	Action       string `json:"action"`
	Payload      any    `json:"payload"`
	IsSyncNeeded bool   `json:"isSyncNeeded"`
}
