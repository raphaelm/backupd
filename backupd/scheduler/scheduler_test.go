package scheduler_test

import (
	"testing"
	"time"

	"github.com/raphaelm/backupd/backupd/model"
	"github.com/raphaelm/backupd/backupd/scheduler"
	"github.com/stretchr/testify/assert"
)

var now = time.Date(2015, 10, 29, 14, 0, 0, 0, time.UTC)

func TestSimpleAddition(t *testing.T) {
	assert := assert.New(t)

	job := model.Job{
		Interval:      time.Duration(1 * time.Hour),
		PreferredTime: time.Date(0, 0, 0, 2, 0, 0, 0, time.UTC),
	}
	last := model.Backup{
		Start:  time.Date(2015, 10, 29, 13, 50, 0, 0, time.UTC),
		End:    time.Date(2015, 10, 29, 13, 52, 0, 0, time.UTC),
		Result: model.BackupCompleted,
	}
	next := scheduler.NextDateRelative(job, last, now)
	assert.Equal(next, time.Date(2015, 10, 29, 14, 50, 0, 0, time.UTC))
}

func TestOnlyInFuture(t *testing.T) {
	job := model.Job{
		Interval:      time.Duration(1 * time.Hour),
		PreferredTime: time.Date(0, 0, 0, 2, 0, 0, 0, time.UTC),
	}
	last := model.Backup{
		Start:  time.Date(2015, 10, 29, 11, 50, 0, 0, time.UTC),
		End:    time.Date(2015, 10, 29, 11, 52, 0, 0, time.UTC),
		Result: model.BackupCompleted,
	}
	next := scheduler.NextDateRelative(job, last, now)
	assert.Equal(t, next, now)
}

func TestPreferredTimeInPast(t *testing.T) {
	job := model.Job{
		Interval:      time.Duration(3 * time.Hour),
		PreferredTime: time.Date(0, 0, 0, 13, 55, 0, 0, time.UTC),
	}
	last := model.Backup{
		Start:  time.Date(2015, 10, 29, 13, 50, 0, 0, time.UTC),
		End:    time.Date(2015, 10, 29, 13, 52, 0, 0, time.UTC),
		Result: model.BackupCompleted,
	}
	next := scheduler.NextDateRelative(job, last, now)
	assert.Equal(t, next, now)
}

func TestPreferredTime(t *testing.T) {
	job := model.Job{
		Interval:      time.Duration(3 * time.Hour),
		PreferredTime: time.Date(0, 0, 0, 15, 0, 0, 0, time.UTC),
	}
	last := model.Backup{
		Start:  time.Date(2015, 10, 29, 13, 50, 0, 0, time.UTC),
		End:    time.Date(2015, 10, 29, 13, 52, 0, 0, time.UTC),
		Result: model.BackupCompleted,
	}
	next := scheduler.NextDateRelative(job, last, now)
	assert.Equal(t, next, time.Date(2015, 10, 29, 15, 00, 0, 0, time.UTC))
}

func TestErroredRetry(t *testing.T) {
	job := model.Job{
		Interval:      time.Duration(3 * time.Hour),
		PreferredTime: time.Date(0, 0, 0, 15, 0, 0, 0, time.UTC),
	}
	last := model.Backup{
		Start:  time.Date(2015, 10, 29, 13, 50, 0, 0, time.UTC),
		End:    time.Date(2015, 10, 29, 13, 52, 0, 0, time.UTC),
		Result: model.BackupFailed,
	}
	next := scheduler.NextDateRelative(job, last, now)
	assert.Equal(t, next, last.End.Add(scheduler.RetryInterval))
}

func TestErroredRetryAtPreferredTime(t *testing.T) {
	job := model.Job{
		Interval:      time.Duration(3 * time.Hour),
		PreferredTime: time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC),
	}
	last := model.Backup{
		Start:  time.Date(2015, 10, 29, 13, 50, 0, 0, time.UTC),
		End:    time.Date(2015, 10, 29, 13, 52, 0, 0, time.UTC),
		Result: model.BackupFailed,
	}
	next := scheduler.NextDateRelative(job, last, now)
	assert.Equal(t, next, time.Date(2015, 10, 29, 14, 0, 0, 0, time.UTC))
}

func TestRelToNow(t *testing.T) {
	job := model.Job{
		Interval:      time.Duration(time.Hour),
		PreferredTime: time.Now().Add(time.Duration(-5 * time.Hour)),
	}
	last := model.Backup{
		Start:  time.Now().Add(time.Duration(-20 * time.Minute)),
		End:    time.Now().Add(time.Duration(-15 * time.Minute)),
		Result: model.BackupCompleted,
	}
	next := scheduler.NextDate(job, last)
	assert.InDelta(t, next.Unix(),
		time.Now().Add(time.Duration(40*time.Minute)).Unix(), 20)
}
