package nul_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-nul"
)

func TestNullable_UnmarshalJSON_int(t *testing.T) {

	tests := []struct{
		JSON string
		Expected  nul.Nullable[int]
	}{
		{
			JSON: `null`,
			Expected: nul.Null[int](),
		},



		{
			JSON: `-65536`,
			Expected: nul.Something[int](-65536),
		},
		{
			JSON: `-65535`,
			Expected: nul.Something[int](-65535),
		},



		{
			JSON: `-256`,
			Expected: nul.Something[int](-256),
		},
		{
			JSON: `-255`,
			Expected: nul.Something[int](-255),
		},



		{
			JSON: `-5`,
			Expected: nul.Something[int](-5),
		},
		{
			JSON: `-4`,
			Expected: nul.Something[int](-4),
		},
		{
			JSON: `-3`,
			Expected: nul.Something[int](-3),
		},
		{
			JSON: `-2`,
			Expected: nul.Something[int](-2),
		},
		{
			JSON: `-1`,
			Expected: nul.Something[int](-1),
		},
		{
			JSON: `0`,
			Expected: nul.Something[int](0),
		},
		{
			JSON: `1`,
			Expected: nul.Something[int](1),
		},
		{
			JSON: `2`,
			Expected: nul.Something[int](2),
		},
		{
			JSON: `3`,
			Expected: nul.Something[int](3),
		},
		{
			JSON: `4`,
			Expected: nul.Something[int](4),
		},
		{
			JSON: `5`,
			Expected: nul.Something[int](5),
		},



		{
			JSON: `255`,
			Expected: nul.Something[int](255),
		},
		{
			JSON: `256`,
			Expected: nul.Something[int](256),
		},



		{
			JSON: `65535`,
			Expected: nul.Something[int](65535),
		},
		{
			JSON: `65536`,
			Expected: nul.Something[int](65536),
		},
	}

	for testNumber, test := range tests {

		var actual nul.Nullable[int]

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON: %q", test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual nullable-int is not what was expected.", testNumber)
			t.Logf("EXPECTED: #%v", expected)
			t.Logf("ACTUAL:   #%v", actual)
			t.Logf("JSON: %q", test.JSON)
			continue
		}
	}
}

func TestNullable_UnmarshalJSON_int_fail(t *testing.T) {

	var op nul.Nullable[int]

	var jason string = "true"

	err := json.Unmarshal([]byte(jason), &op)
	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}
}
