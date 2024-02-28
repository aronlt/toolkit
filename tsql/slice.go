package tsql

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

// Strings db column value type.
type Strings []string

// Scan implements the sql.Scanner interface.
func (s *Strings) Scan(value interface{}) error {
	return JSONScan(value, s)
}

// Value implements the driver.Valuer interface.
func (s Strings) Value() (driver.Value, error) {
	return JSONArray(s, s == nil)
}

// IntStringMap db column value type.
type IntStringMap map[int64]string

// Scan implements the sql.Scanner interface.
func (s *IntStringMap) Scan(value interface{}) error {
	return JSONScan(value, s)
}

// Value implements the driver.Valuer interface.
func (s IntStringMap) Value() (driver.Value, error) {
	return JSONObject(s, s == nil)
}

// JSONScan parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func JSONScan(value, v interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("invalid value types")
	}
	if len(data) == 0 {
		v = make([]string, 0)
		return nil
	}
	return json.Unmarshal(data, v)
}

// JSONArray returns the JSON encoding of v array.
func JSONArray(v interface{}, isNil bool) (driver.Value, error) {
	if isNil {
		return []byte{'[', ']'}, nil
	}
	return json.Marshal(v)
}

// JSONObject returns the JSON encoding of v object.
func JSONObject(v interface{}, isNil bool) (driver.Value, error) {
	if isNil {
		return []byte{'{', '}'}, nil
	}
	return json.Marshal(v)
}
