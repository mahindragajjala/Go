// models/audit.go
package models

import "time"

type AuditLog struct {
	Email     string    `bson:"email"`
	IP        string    `bson:"ip"`
	Action    string    `bson:"action"`
	Timestamp time.Time `bson:"timestamp"`
}
