package nul_test

import (
	"testing"

	"sourcecode.social/reiver/go-nul"
)

func TestOptional_Filter_int(t *testing.T) {

	tests := []struct{
		Nullable nul.Nullable[int]
		Expected nul.Nullable[int]
	}{
		{
			Nullable: nul.Nothing[int](),
			Expected: nul.Nothing[int](),
		},



		{
			Nullable: nul.Something[int](-2),
			Expected: nul.Something[int](-2),
		},
		{
			Nullable: nul.Something[int](-1),
			Expected: nul.Nothing[int](),
		},
		{
			Nullable: nul.Something[int](0),
			Expected: nul.Something[int](0),
		},
		{
			Nullable: nul.Something[int](1),
			Expected: nul.Nothing[int](),
		},
		{
			Nullable: nul.Something[int](2),
			Expected: nul.Something[int](2),
		},
	}

	for testNumber, test := range tests {

		fn := func(value int) bool {
			return 0 == (value % 2)
		}

		actual   := test.Nullable.Filter(fn)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
	/////////////// CONTINUE
			continue
		}
	}
}
