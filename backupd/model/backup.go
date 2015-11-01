package model

import "time"

const (
	BackupCompleted = 0
	BackupFailed    = 1
)

type Backup struct {
	ID     int64     `json:"id"`
	JobID  int64     `json:"job_id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Result int64     `json:"result"`
	Log    string    `json:"log"`
}
