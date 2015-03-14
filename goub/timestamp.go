package goub

import "time"

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

func (t1 Timestamp) Equal(t2 Timestamp) bool {
	return t1.Time.Equal(t2.Time)
}
