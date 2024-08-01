package nul_test

import (
	"testing"

	"fmt"

	"github.com/reiver/go-nul"
)

func TestNullable_GoString(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value:                           "",
			Expected: `nul.Something[string]("")`,
		},
		{
			Value:                           "once twice thrice fource",
			Expected: `nul.Something[string]("once twice thrice fource")`,
		},
		{
			Value:                           "apple banana cherry",
			Expected: `nul.Something[string]("apple banana cherry")`,
		},



		{
			Value:                   uint8 (0x0),
			Expected: `nul.Something[uint8](0x0)`,
		},
		{
			Value:                   uint8 (0x1),
			Expected: `nul.Something[uint8](0x1)`,
		},
		{
			Value:                   uint8 (0x2),
			Expected: `nul.Something[uint8](0x2)`,
		},
		{
			Value:                   uint8 (0xfe),
			Expected: `nul.Something[uint8](0xfe)`,
		},
		{
			Value:                   uint8 (0xff),
			Expected: `nul.Something[uint8](0xff)`,
		},



		{
			Value:                   uint16 (0x0),
			Expected: `nul.Something[uint16](0x0)`,
		},
		{
			Value:                   uint16 (0x1),
			Expected: `nul.Something[uint16](0x1)`,
		},
		{
			Value:                   uint16 (0x2),
			Expected: `nul.Something[uint16](0x2)`,
		},
		{
			Value:                   uint16 (0xfe),
			Expected: `nul.Something[uint16](0xfe)`,
		},
		{
			Value:                   uint16 (0xff),
			Expected: `nul.Something[uint16](0xff)`,
		},
		{
			Value:                   uint16 (0x100),
			Expected: `nul.Something[uint16](0x100)`,
		},
		{
			Value:                   uint16 (0x101),
			Expected: `nul.Something[uint16](0x101)`,
		},
		{
			Value:                   uint16 (0x102),
			Expected: `nul.Something[uint16](0x102)`,
		},
		{
			Value:                   uint16 (0xfffe),
			Expected: `nul.Something[uint16](0xfffe)`,
		},
		{
			Value:                   uint16 (0xffff),
			Expected: `nul.Something[uint16](0xffff)`,
		},



		{
			Value:                   struct { A string; B int }{A:"joeblow",B:7},
			Expected: `nul.Something[struct { A string; B int }](struct { A string; B int }{A:"joeblow", B:7})`,
		},
	}

	for testNumber, test := range tests {

		op := nul.Something(test.Value)
		gostring := op.GoString()


		{
			expected := test.Expected
			actual   := gostring

			if expected != actual {
				t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE-TYPE: %T", test.Value)
				t.Logf("VALUE: %#v", test.Value)
	//////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestNullable_GoString_nothing(t *testing.T) {

	tests := []struct{
		Nullable fmt.GoStringer
		Expected string
	}{
		{
			Nullable:  nul.Nothing[string](),
			Expected: `nul.Nothing[string]()`,
		},



		{
			Nullable:  nul.Nothing[int8](),
			Expected: `nul.Nothing[int8]()`,
		},
		{
			Nullable:  nul.Nothing[int16](),
			Expected: `nul.Nothing[int16]()`,
		},
		{
			Nullable:  nul.Nothing[int32](),
			Expected: `nul.Nothing[int32]()`,
		},
		{
			Nullable:  nul.Nothing[int64](),
			Expected: `nul.Nothing[int64]()`,
		},



		{
			Nullable:  nul.Nothing[uint8](),
			Expected: `nul.Nothing[uint8]()`,
		},
		{
			Nullable:  nul.Nothing[uint16](),
			Expected: `nul.Nothing[uint16]()`,
		},
		{
			Nullable:  nul.Nothing[uint32](),
			Expected: `nul.Nothing[uint32]()`,
		},
		{
			Nullable:  nul.Nothing[uint64](),
			Expected: `nul.Nothing[uint64]()`,
		},
	}

	for testNumber, test := range tests {

		expected := test.Expected
		actual   := test.Nullable.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual go-string value is not what was expected.", testNumber)
			t.Logf("EXPECTED GO-STRING: %q", expected)
			t.Logf("ACTUAL   GO-STRING: %q", actual)
			t.Logf("TYPE: %T", test.Nullable)
	/////////////// CONTINUE
			continue
		}
	}
}

func TestNullable_GoString_null(t *testing.T) {

	tests := []struct{
		Nullable fmt.GoStringer
		Expected string
	}{
		{
			Nullable:  nul.Null[string](),
			Expected: `nul.Null[string]()`,
		},



		{
			Nullable:  nul.Null[int8](),
			Expected: `nul.Null[int8]()`,
		},
		{
			Nullable:  nul.Null[int16](),
			Expected: `nul.Null[int16]()`,
		},
		{
			Nullable:  nul.Null[int32](),
			Expected: `nul.Null[int32]()`,
		},
		{
			Nullable:  nul.Null[int64](),
			Expected: `nul.Null[int64]()`,
		},



		{
			Nullable:  nul.Null[uint8](),
			Expected: `nul.Null[uint8]()`,
		},
		{
			Nullable:  nul.Null[uint16](),
			Expected: `nul.Null[uint16]()`,
		},
		{
			Nullable:  nul.Null[uint32](),
			Expected: `nul.Null[uint32]()`,
		},
		{
			Nullable:  nul.Null[uint64](),
			Expected: `nul.Null[uint64]()`,
		},
	}

	for testNumber, test := range tests {

		expected := test.Expected
		actual   := test.Nullable.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual go-string value is not what was expected.", testNumber)
			t.Logf("EXPECTED GO-STRING: %q", expected)
			t.Logf("ACTUAL   GO-STRING: %q", actual)
			t.Logf("TYPE: %T", test.Nullable)
	/////////////// CONTINUE
			continue
		}
	}
}
