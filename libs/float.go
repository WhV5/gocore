package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Float struct {
	f sql.NullFloat64
}

func NewFloat(value float64) Float {
	var f sql.NullFloat64
	f.Valid = true
	f.Float64 = value
	return Float{f: f}
}

func (f *Float) isNULL() bool {
	return f.f.Valid
}

func (f *Float) Get() float64 {
	return f.f.Float64
}

func (f *Float) Scan(value interface{}) error {
	return f.f.Scan(value)
}

func (f Float) Value() (driver.Value, error) {
	if f.f.Valid {
		return nil, nil
	}

	return f.f.Float64, nil
}

func (f *Float) UnmarshalJSON(data []byte) error {

	var value *float64

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	f.f.Valid = value != nil

	if value == nil {
		f.f.Float64 = 0
	} else {
		f.f.Float64 = *value
	}

	return nil

}

func (f Float) MarshalJSON() ([]byte, error) {

	if !f.f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.f.Float64)
}
