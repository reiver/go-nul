package nul_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-nul"
)

func TestNullable_UnmarshalJSON_string(t *testing.T) {

	tests := []struct{
		JSON string
		Expected nul.Nullable[string]
	}{
		{
			JSON: `null`,
			Expected: nul.Null[string](),
		},



		{
			JSON: `""`,
			Expected: nul.Something(""),
		},



		{
			JSON: `"apple"`,
			Expected: nul.Something("apple"),
		},
		{
			JSON: `"banana"`,
			Expected: nul.Something("banana"),
		},
		{
			JSON: `"cherry"`,
			Expected: nul.Something("cherry"),
		},



		{
			JSON: `"ONCE"`,
			Expected: nul.Something("ONCE"),
		},
		{
			JSON: `"TWICE"`,
			Expected: nul.Something("TWICE"),
		},
		{
			JSON: `"THRICE"`,
			Expected: nul.Something("THRICE"),
		},
		{
			JSON: `"FOURCE"`,
			Expected: nul.Something("FOURCE"),
		},



		{
			JSON: `"🙂"`,
			Expected: nul.Something("🙂"),
		},
		{
			JSON: `"😈"`,
			Expected: nul.Something("😈"),
		},
		{
			JSON: `"❤️"`,
			Expected: nul.Something("❤️"),
		},



		{
			JSON: `"٠١٢٣۴۵۶٧٨٩"`,
			Expected: nul.Something("٠١٢٣۴۵۶٧٨٩"),
		},



		{
			JSON: `"𐏑𐏓𐏕"`,
			Expected: nul.Something("𐏑𐏓𐏕"),
		},
	}

	for testNumber, test := range tests {

		var actual nul.Nullable[string]

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
				t.Errorf("For test #%d, the actual value of the nullable optional type is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON: %q", test.JSON)
				continue
			}
		}
	}
}

func TestNullable_UnmarshalJSON_string_fail(t *testing.T) {

	var op nul.Nullable[string]

	var jason string = "12345"

	err := json.Unmarshal([]byte(jason), &op)
	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}
}
