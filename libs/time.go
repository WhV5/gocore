package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Time struct {
	t time.Time
	v bool
}

func NewTime(t time.Time) Time {
	return Time{
		t: t,
		v: true,
	}
}

func (t *Time) IsNULL() bool {
	return t.v
}

func (t *Time) Get() time.Time {
	if !t.v {
		return time.Unix(0, 0)
	}
	return t.t
}

func (t Time) MarshalJSON() ([]byte, error) {
	if !t.v {
		return []byte("null"), nil
	}
	return json.Marshal(t.t.Format("2006/01/02 15:04:05"))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var ts *time.Time
	if err := json.Unmarshal(data, &ts); err != nil {
		return err
	}

	t.v = ts != nil

	if ts == nil {
		t.t = time.Unix(0, 0)
	} else {
		t.t = *ts
	}

	return nil

}

func (t *Time) Scan(value interface{}) error {
	t.t, t.v = value.(time.Time)
	if t.v {
		return nil
	}
	var ns sql.NullString

	if err := ns.Scan(value); err != nil {
		return err
	}

	if !ns.Valid {
		return nil
	}

	for _, tf := range timestampFormats {
		if tt, err := time.Parse(tf, ns.String); err == nil {
			t.t = tt
			t.v = true
			return nil
		}
	}

	return nil
}

func (t Time) Value() (driver.Value, error) {
	if !t.v {
		return nil, nil
	}

	return t.t, nil
}

var timestampFormats = []string{
	"2006-01-02 15:04:05.999999999-07:00",
	"2006-01-02T15:04:05.999999999-07:00",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
}
