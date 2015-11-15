package model

import (
	"errors"
	"time"
)

var timeLayout = "15:04"
var TimeParseError = errors.New(`TimeParseError: should be a string formatted as "15:04:05"`)

type ClockTime struct {
	Hour, Minute int
}

func (t ClockTime) toTime() time.Time {
	return time.Date(0, 0, 0, t.Hour, t.Minute, 0, 0, time.UTC)
}

func (t ClockTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.toTime().Format(timeLayout) + `"`), nil
}

func (t *ClockTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// len(`"23:59"`) == 7
	if len(s) != 7 {
		return TimeParseError
	}
	ret, err := time.Parse(timeLayout, s[1:6])
	if err != nil {
		return err
	}
	t.Hour = ret.Hour()
	t.Minute = ret.Minute()
	return nil
}
