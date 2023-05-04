package shared

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

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
	*mt = TimeToDateTime(t)
	return nil
}

func (mt *DateTime) MarshalJSON() ([]byte, error) {
	location, err := time.LoadLocation("")
	if err != nil {
		return nil, err
	}
	return json.Marshal(mt.Time.In(location).Format(dateFormat))
}

func (mt *DateTime) Scan(value interface{}) error {
	mt.Time = value.(time.Time)
	return nil
}

func (mt DateTime) Value() (driver.Value, error) {
	return mt.Time, nil
}

func TimeToDateTime(time time.Time) DateTime {
	return DateTime{time}
}

func StringToDate(date string) (*DateTime, error) {
	t, err := time.ParseInLocation(dateFormat, date, time.UTC)
	if err != nil {
		return nil, err
	}
	dateTime := TimeToDateTime(t)
	return &dateTime, nil
}

type DateTimeList []DateTime

func (dateTimeList DateTimeList) Len() int {
	return len(dateTimeList)
}

func (dateTimeList DateTimeList) Less(i, j int) bool {
	return dateTimeList[i].Time.Before(dateTimeList[j].Time)
}

func (dateTimeList DateTimeList) Swap(i, j int) {
	dateTimeList[i], dateTimeList[j] = dateTimeList[j], dateTimeList[i]
}
