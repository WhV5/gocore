package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Integer struct {
	i sql.NullInt64
}

func NewInteger(value int64) Integer {
	var s sql.NullInt64
	s.Valid = true
	s.Int64 = value
	return Integer{i: s}
}

func (i *Integer) Get() int64 {
	return i.i.Int64
}

func (i *Integer) IsNULL() bool {
	return i.i.Valid
}

func (i *Integer) Set(value int64) *Integer {
	i.i.Valid = true
	i.i.Int64 = value
	return i
}

func (i *Integer) Scan(value interface{}) error {
	return i.i.Scan(value)
}

func (i Integer) Value() (driver.Value, error) {
	if !i.i.Valid {
		return nil, nil
	}
	return i.i.Int64, nil
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.i.Valid = value != nil
	if value == nil {
		i.i.Int64 = 0
	} else {
		i.i.Int64 = *value
	}

	return nil

}

func (i Integer) MarshalJSON() ([]byte, error) {
	if !i.i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.i.Int64)
}
