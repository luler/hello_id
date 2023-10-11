package type_helper

import "time"

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(timeFormat) + `"`), nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
