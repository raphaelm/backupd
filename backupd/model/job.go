package model

import "time"

type Job struct {
	ID            int64
	RemoteId      int64
	JobName       string
	Interval      time.Duration
	PreferredTime time.Time
}
