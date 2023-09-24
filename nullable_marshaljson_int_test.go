package nul_test

import (
	"testing"

	"sourcecode.social/reiver/go-nul"
)

func TestNullable_MarshalJSON_int(t *testing.T) {

	tests := []struct{
		Value nul.Nullable[int]
	}{
		{
			Value: nul.Null[int](),
		},



		{
			Value: nul.Something(-65536),
		},
		{
			Value: nul.Something(-65535),
		},



		{
			Value: nul.Something(-256),
		},
		{
			Value: nul.Something(-255),
		},



		{
			Value: nul.Something(-5),
		},
		{
			Value: nul.Something(-4),
		},
		{
			Value: nul.Something(-3),
		},
		{
			Value: nul.Something(-2),
		},
		{
			Value: nul.Something(-1),
		},
		{
			Value: nul.Something(0),
		},
		{
			Value: nul.Something(1),
		},
		{
			Value: nul.Something(2),
		},
		{
			Value: nul.Something(3),
		},
		{
			Value: nul.Something(4),
		},
		{
			Value: nul.Something(5),
		},



		{
			Value: nul.Something(255),
		},
		{
			Value: nul.Something(256),
		},



		{
			Value: nul.Something(65535),
		},
		{
			Value: nul.Something(65536),
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one." , testNumber)
			t.Logf("ACTUAL BYTES: %q", actualBytes)
			t.Logf("ERROR: (%T) %s", err, err)
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
