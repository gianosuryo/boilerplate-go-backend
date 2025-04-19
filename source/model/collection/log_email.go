package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogEmail struct {
	ID primitive.ObjectID `bson:"_id"`

	// Main information
	At      int64  `bson:"at"`
	To      string `bson:"to"`
	Cc      string `bson:"cc"`
	Bcc     string `bson:"cc"`
	Subject string `bson:"subject"`

	// Sending status
	IsSent bool   `bson:"is_sent"`
	Error  string `bson:"error"`
}

func (LogEmail) CollectionName() string {
	return "log-email"
}
