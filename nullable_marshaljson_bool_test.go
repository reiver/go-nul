package nul_test

import (
	"testing"

	"github.com/reiver/go-nul"
)

func TestNullable_MarshalJSON_bool(t *testing.T) {

	tests := []struct{
		Value nul.Nullable[bool]
		Expected string
	}{
		{
			Value: nul.Null[bool](),
			Expected: "null",
		},
		{
			Value: nul.Something(false),
			Expected: "false",
		},
		{
			Value: nul.Something(true),
			Expected: "true",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for the JSON marshaling is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}

func TestNullable_MarshalJSON_bool_fail(t *testing.T) {

	var nothing nul.Nullable[bool]

	actualBytes, err := nothing.MarshalJSON()
	if nil == err {
		t.Error("Expected an error but did not actually get one.")
		t.Logf("ACTUAL: %q", actualBytes)
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
	if nil != actualBytes {
		t.Error("Expected not bytes but actually get some.")
		t.Logf("ACTUAL: %q", actualBytes)
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
}
