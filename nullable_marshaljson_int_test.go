package nul_test

import (
	"testing"

	"github.com/reiver/go-nul"
)

func TestNullable_MarshalJSON_int(t *testing.T) {

	tests := []struct{
		Value nul.Nullable[int]
		Expected string
	}{
		{
			Value: nul.Null[int](),
			Expected: "null",
		},



		{
			Value: nul.Something(-65536),
			Expected:           `-65536`,
		},
		{
			Value: nul.Something(-65535),
			Expected:           `-65535`,
		},



		{
			Value: nul.Something(-256),
			Expected:           `-256`,
		},
		{
			Value: nul.Something(-255),
			Expected:           `-255`,
		},



		{
			Value: nul.Something(-5),
			Expected:           `-5`,
		},
		{
			Value: nul.Something(-4),
			Expected:           `-4`,
		},
		{
			Value: nul.Something(-3),
			Expected:           `-3`,
		},
		{
			Value: nul.Something(-2),
			Expected:           `-2`,
		},
		{
			Value: nul.Something(-1),
			Expected:           `-1`,
		},
		{
			Value: nul.Something(0),
			Expected:           `0`,
		},
		{
			Value: nul.Something(1),
			Expected:           `1`,
		},
		{
			Value: nul.Something(2),
			Expected:           `2`,
		},
		{
			Value: nul.Something(3),
			Expected:           `3`,
		},
		{
			Value: nul.Something(4),
			Expected:           `4`,
		},
		{
			Value: nul.Something(5),
			Expected:           `5`,
		},



		{
			Value: nul.Something(255),
			Expected:           `255`,
		},
		{
			Value: nul.Something(256),
			Expected:           `256`,
		},



		{
			Value: nul.Something(65535),
			Expected:           `65535`,
		},
		{
			Value: nul.Something(65536),
			Expected:           `65536`,
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

func TestNullable_MarshalJSON_int_fail(t *testing.T) {

	var nothing nul.Nullable[int]

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
