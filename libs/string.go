package libs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

//String 实现 database/sql/driver null 无法问题
type String struct {
	s sql.NullString
}

//IsNULL 判断字符串是否为NULL
func (s *String) IsNULL() bool {
	return s.s.Valid
}

//Get 获取数据
func (s *String) Get() string {
	return s.s.String
}

func (s *String) Valid() bool {
	return s.s.Valid
}

//创建String
func NewString(str string) String {
	var s sql.NullString
	s.String = str
	s.Valid = str != ""
	return String{s: s}
}

//Scan 实现 Scan，sql/driver 的数据 调用此方法获取值
func (s *String) Scan(value interface{}) error {
	return s.s.Scan(value)
}

//Value 实现 sql/driver Value ,调用此方法设置值
func (s String) Value() (driver.Value, error) {
	if !s.s.Valid {
		return nil, nil
	}
	return s.s.String, nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte(""), nil
	}

	return json.Marshal(s.s.String)
}

func (s *String) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	s.s.Valid = value != nil
	if value == nil {
		s.s.String = ""
	} else {
		s.s.String = *value
	}
	return nil
}
