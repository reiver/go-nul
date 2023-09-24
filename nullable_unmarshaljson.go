package nul

import (
	"fmt"
	"encoding/json"
)

var _ json.Unmarshaler = new(Nullable[bool])
var _ json.Unmarshaler = new(Nullable[string])

// UnmarshalJSON makes it so json.Unmarshaler is implemented.
func (receiver *Nullable[T]) UnmarshalJSON(data []byte) error {
	switch interface{}(receiver.value).(type) {
	case bool, string:
		// these are OK.
	default:
		return fmt.Errorf("cannot unmarshal into something of type %T from JSON because parameterized type is ‘%T’ rather than ‘bool’ or ‘string’", receiver, receiver.value)
	}

	if 4 == len(data) && 'n' == data[0] && 'u' == data[1] && 'l' == data[2] && 'l' == data[3] {
		*receiver = Null[T]()
		return nil

	}

	{
		var dst T

		err := json.Unmarshal(data, &dst)
		if nil != err {
			return err
		}

		*receiver = Something(dst)
	}

	return nil
}
