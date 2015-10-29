package model

import "time"

const (
	BackupCompleted = 0
	BackupFailed    = 1
)

type Backup struct {
	ID     int64
	JobId  int64
	start  time.Time
	end    time.Time
	result int64
	log    string
}
