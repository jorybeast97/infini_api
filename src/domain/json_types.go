package domain

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
)

func (l Location) Value() (driver.Value, error) {
    b, err := json.Marshal(l)
    if err != nil { return nil, err }
    return string(b), nil
}

func (l *Location) Scan(src interface{}) error {
    switch v := src.(type) {
    case []byte:
        return json.Unmarshal(v, l)
    case string:
        return json.Unmarshal([]byte(v), l)
    default:
        return fmt.Errorf("unsupported type %T", src)
    }
}

func (s Social) Value() (driver.Value, error) {
    b, err := json.Marshal(s)
    if err != nil { return nil, err }
    return string(b), nil
}

func (s *Social) Scan(src interface{}) error {
    switch v := src.(type) {
    case []byte:
        return json.Unmarshal(v, s)
    case string:
        return json.Unmarshal([]byte(v), s)
    default:
        return fmt.Errorf("unsupported type %T", src)
    }
}

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
    b, err := json.Marshal(a)
    if err != nil { return nil, err }
    return string(b), nil
}

func (a *StringArray) Scan(src interface{}) error {
    switch v := src.(type) {
    case []byte:
        return json.Unmarshal(v, a)
    case string:
        return json.Unmarshal([]byte(v), a)
    default:
        return fmt.Errorf("unsupported type %T", src)
    }
}

