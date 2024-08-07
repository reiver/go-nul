package nul_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-nul"
)

func TestNullable_UnmarshalJSON_int(t *testing.T) {

	tests := []struct{
		JSON string
	}{
		{
			JSON: `null`,
		},



		{
			JSON: `-65536`,
		},
		{
			JSON: `-65535`,
		},



		{
			JSON: `-256`,
		},
		{
			JSON: `-255`,
		},



		{
			JSON: `-5`,
		},
		{
			JSON: `-4`,
		},
		{
			JSON: `-3`,
		},
		{
			JSON: `-2`,
		},
		{
			JSON: `-1`,
		},
		{
			JSON: `0`,
		},
		{
			JSON: `1`,
		},
		{
			JSON: `2`,
		},
		{
			JSON: `3`,
		},
		{
			JSON: `4`,
		},
		{
			JSON: `5`,
		},



		{
			JSON: `-255`,
		},
		{
			JSON: `-256`,
		},



		{
			JSON: `65535`,
		},
		{
			JSON: `65536`,
		},
	}

	for testNumber, test := range tests {

		var actual nul.Nullable[int]

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}

func TestNullable_UnmarshalJSON_int_fail(t *testing.T) {

	var op nul.Nullable[int]

	var jason string = "12345"

	err := json.Unmarshal([]byte(jason), &op)
	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}
}
