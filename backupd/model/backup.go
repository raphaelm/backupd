package model

import "time"

const (
	BackupCompleted = 0
	BackupFailed    = 1
)

type Backup struct {
	ID     int64
	JobId  int64
	Start  time.Time
	End    time.Time
	Result int64
	Log    string
}
