package scheduler

import (
	"time"

	"github.com/raphaelm/backupd/backupd/model"
)

const RetryInterval = time.Duration(30 * time.Minute)

// NextDate calculates the next date the given job should be executed.
// In order to do this correctly, it also needs information about the
// last backup that has been run
func NextDate(job model.Job, lastBackup model.Backup) time.Time {
	return NextDateRelative(job, lastBackup, time.Now())
}

// NextDateRelative is the real implementation of NextDate, but calculates
// relative to a configurable time instead of the current time. This is
// mainly used for having deterministic unit tests.
func NextDateRelative(job model.Job, lastBackup model.Backup, now time.Time) time.Time {
	nextTime := lastBackup.Start.Add(job.Interval)
	lastTime := lastBackup.Start
	if lastBackup.Result != model.BackupCompleted {
		nextTime = lastBackup.End.Add(RetryInterval)
		lastTime = lastBackup.End
	}

	preferredOnNextDay := time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(),
		job.PreferredTime.Hour, job.PreferredTime.Minute, 0, 0, time.UTC)

	if lastTime.Before(preferredOnNextDay) && nextTime.After(preferredOnNextDay) {
		nextTime = preferredOnNextDay
	}

	if now.After(nextTime) {
		nextTime = now
	}

	return nextTime
}
