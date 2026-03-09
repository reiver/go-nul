package nul

import (
	"encoding/json"
	"reflect"

	"codeberg.org/reiver/go-erorr"
)

const errNilReceiver = erorr.Error("nul: nil receiver")

var _ json.Unmarshaler = new(Nullable[bool])
var _ json.Unmarshaler = new(Nullable[string])

// UnmarshalJSON makes it so json.Unmarshaler is implemented.
func (receiver *Nullable[T]) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var ok bool
	{
		switch interface{}(receiver.value).(type) {
		case
			bool,
			int,
			int8,
			int16,
			int32,
			int64,
			string,
			uint,
			uint8,
			uint16,
			uint32,
			uint64,
			json.Unmarshaler:
			ok = true
		}

                // Also check if *T implements json.Unmarshaler (pointer receiver methods).
                var ptr interface{} = &receiver.value
                if _, casted := ptr.(json.Unmarshaler); casted {
                        ok = true
                }

		if reflect.Struct == reflect.TypeOf(receiver.value).Kind() {
			ok = true
		}
	}

	if !ok {
		return erorr.Errorf("nul: cannot unmarshal into something of type %T from JSON because parameterized type is ‘%T’", receiver, receiver.value)
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
