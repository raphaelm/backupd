package model

import "time"

type Job struct {
	ID            int64         `json:"id"`
	RemoteID      int64         `json:"remote_id"`
	JobName       string        `json:"job_name"`
	Interval      time.Duration `json:"interval"`
	PreferredTime time.Time     `json:"preferred_time"`
}
type Jobs []Job
