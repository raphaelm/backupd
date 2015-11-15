package model_test

import (
	"testing"

	"github.com/raphaelm/backupd/backupd/model"
	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	ct := model.ClockTime{}
	assert.Nil(t, ct.UnmarshalJSON([]byte(`"13:37"`)))
	assert.Equal(t, 13, ct.Hour)
	assert.Equal(t, 37, ct.Minute)
	out, err := ct.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`"13:37"`), out)

	assert.NotNil(t, ct.UnmarshalJSON([]byte(`"13:XX`)))
	assert.NotNil(t, ct.UnmarshalJSON([]byte(`"13:70"`)))

	ct = model.ClockTime{Hour: 22, Minute: 42}
	assert.Equal(t, 22, ct.Hour)
	assert.Equal(t, 42, ct.Minute)
}
