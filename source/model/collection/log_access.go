package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogAccess struct {
	ID primitive.ObjectID `bson:"_id"`
	// Connection
	At     int64             `bson:"at"`
	Method string            `bson:"method"`
	Path   string            `bson:"path"`
	Query  map[string]string `bson:"query"`
	Ip     string            `bson:"ip"`
	// Response status
	Status     int    `bson:"status"`
	StatusText string `bson:"status_text"`
	// Latency number
	LatencyMs   uint64
	LatencyText string `bson:"latency_text"`
	// User
	UserId   uint64   `bson:"user_id"`
	UserName string   `bson:"user_name"`
	Role     []string `bson:"role"`
}

func (LogAccess) CollectionName() string {
	return "log-access"
}
