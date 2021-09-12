package tools

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type MyTime struct {
	time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatTime := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(formatTime), nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v", v)
}