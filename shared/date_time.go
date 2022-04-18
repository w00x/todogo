package shared

import (
	"encoding/json"
	"time"
)

type DateTime time.Time

var _ json.Unmarshaler = &DateTime{}

const dateFormat = "2006-01-02 15:04:05"

func (mt *DateTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(dateFormat, s, time.UTC)
	if err != nil {
		return err
	}
	*mt = DateTime(t)
	return nil
}

func (mt *DateTime) MarshalJSON() ([]byte, error) {
	location, err := time.LoadLocation("")
	if err != nil {
		return nil, err
	}
	return json.Marshal(time.Time(*mt).In(location).Format(dateFormat))
}