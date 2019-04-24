package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Boolean struct {
	b sql.NullBool
}

func NewBoolean(value bool) Boolean {
	var b sql.NullBool
	b.Valid = true
	b.Bool = value
	return Boolean{b: b}
}

func (b *Boolean) Valid() bool {
	return b.b.Valid
}

func (b *Boolean) Get() bool {
	return b.b.Bool
}

func (b Boolean) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	b.b.Valid = value != nil

	if value == nil {
		b.b.Bool = false
	} else {
		b.b.Bool = *value
	}

	return nil

}

func (b *Boolean) MarshalJSON() ([]byte, error) {
	if !b.Valid() {
		return []byte("null"), nil
	}

	return json.Marshal(b.b.Bool)
}

func (b *Boolean) Scan(value interface{}) error {
	return b.b.Scan(value)
}

func (b Boolean) Value() (driver.Value, error) {
	if !b.Valid() {
		return nil, nil
	}
	return b.b.Bool, nil
}
