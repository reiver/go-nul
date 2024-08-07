package nul

import (
	"encoding/json"
	"reflect"

	"github.com/reiver/go-erorr"
)

var _ json.Marshaler = Nothing[bool]()
var _ json.Marshaler = Nothing[string]()

// MarshalJSON makes it so json.Marshaler is implemented.
func (receiver Nullable[T]) MarshalJSON() ([]byte, error) {
	switch interface{}(receiver.value).(type) {
	case bool, string,json.Marshaler:
		// these are OK.
	default:
		if reflect.Struct != reflect.TypeOf(receiver.value).Kind() {
			return nil, erorr.Errorf("nul: cannot marshal something of type %T into JSON because parameterized type is ‘%T’ rather than ‘bool’, ‘string’, or ‘json.Marshaler’", receiver, receiver.value)
		}
	}

	if receiver.isnothing() {
		return nil, erorr.Errorf("nul: cannot marshal nul.Nothing[%T]() into JSON", receiver.value)
	}

	if receiver.isnull {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}
