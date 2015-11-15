package model

import (
	"errors"
	"time"
)

var timeLayout = "15:04"
var TimeParseError = errors.New(`TimeParseError: should be a string formatted as "15:04:05"`)

type ClockTime struct {
	time.Time
}

func (t ClockTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(timeLayout) + `"`), nil
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
	t.Time = ret
	return nil
}
