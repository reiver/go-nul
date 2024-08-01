package nul_test

import (
	"testing"

	"github.com/reiver/go-nul"
)

func TestNullable_MarshalJSON_string(t *testing.T) {

	tests := []struct{
		Value nul.Nullable[string]
		Expected string
	}{
		{
			Value: nul.Null[string](),
			Expected: `null`,
		},



		{
			Value: nul.Something(""),
			Expected: `""`,
		},



		{
			Value: nul.Something("apple"),
			Expected: `"apple"`,
		},
		{
			Value: nul.Something("banana"),
			Expected: `"banana"`,
		},
		{
			Value: nul.Something("cherry"),
			Expected: `"cherry"`,
		},



		{
			Value: nul.Something("ONCE"),
			Expected: `"ONCE"`,
		},
		{
			Value: nul.Something("TWICE"),
			Expected: `"TWICE"`,
		},
		{
			Value: nul.Something("THRICE"),
			Expected: `"THRICE"`,
		},
		{
			Value: nul.Something("FOURCE"),
			Expected: `"FOURCE"`,
		},



		{
			Value: nul.Something("🙂"),
			Expected: `"🙂"`,
		},
		{
			Value: nul.Something("😈"),
			Expected: `"😈"`,
		},
		{
			Value: nul.Something("❤️"),
			Expected: `"❤️"`,
		},



		{
			Value: nul.Something("٠١٢٣۴۵۶٧٨٩"),
			Expected: `"٠١٢٣۴۵۶٧٨٩"`,
		},



		{
			Value: nul.Something("𐏑𐏓𐏕"),
			Expected: `"𐏑𐏓𐏕"`,
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

func TestNullable_MarshalJSON_string_fail(t *testing.T) {

	var nothing nul.Nullable[string]

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
