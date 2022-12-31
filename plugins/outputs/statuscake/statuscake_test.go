package statuscake

import (
	"github.com/influxdata/telegraf/testutil"
)

func NewStatusCake() *StatusCake {
	return &StatusCake{
		Log: testutil.Logger{},
	}
}
