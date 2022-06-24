package jsonrpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ID represents the identifier of a Request or Response.
type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) Int() (int, error) {
	return strconv.Atoi(string(i))
}

func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(i) + `"`), nil
}

func (i *ID) UnmarshalJSON(b []byte) error {
	var out any
	if err := json.Unmarshal(b, &out); err != nil {
		return err
	}

	switch v := out.(type) {
	case string:
		*i = ID(v)
	case int, int16, int32, int64:
		*i = ID(strconv.Itoa(v.(int)))
	default:
		return errors.New("id: unmarshal: id must be either a string or a number")
	}

	return nil
}
