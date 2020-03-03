package inst

import "time"

type ClusterUserMessage struct {
	ID          int64
	ClusterName string
	Level       string
	Message     string
	CreatedAt   time.Time
}
