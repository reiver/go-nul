package nul_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-nul"
)

func TestNullable_UnmarshalJSON_bool(t *testing.T) {

	tests := []struct{
		JSON string
		Expected nul.Nullable[bool]
	}{
		{
			JSON: `null`,
			Expected: nul.Null[bool](),
		},
		{
			JSON: `false`,
			Expected: nul.Something(false),
		},
		{
			JSON: `true`,
			Expected: nul.Something(true),
		},
	}

	for testNumber, test := range tests {

		var actual nul.Nullable[bool]

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value of the optional type is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON: %q", test.JSON)
				continue
			}
		}
	}
}

func TestNullable_UnmarshalJSON_bool_fail(t *testing.T) {

	var op nul.Nullable[bool]

	var jason string = "12345"

	err := json.Unmarshal([]byte(jason), &op)
	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}
}
