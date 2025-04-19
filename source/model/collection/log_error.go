package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogError struct {
	ID primitive.ObjectID `bson:"_id"`
	// Connection
	At     int64  `bson:"at"`
	Method string `bson:"method"`
	Path   string `bson:"path"`
	Query  string `bson:"query"`
	// Response status
	Status     int    `bson:"status"`
	StatusText string `bson:"status_text"`
}

func (LogError) CollectionName() string {
	return "log-error"
}
