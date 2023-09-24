package nul_test

import (
	"testing"

	"sourcecode.social/reiver/go-nul"
)

func TestNullable_WhenNothing(t *testing.T) {


	tests := []struct{
		Nullable interface{WhenNothing(func())}
	}{
		{
			Nullable: nul.Nothing[string](),
		},



		{
			Nullable: nul.Nothing[int](),
		},
		{
			Nullable: nul.Nothing[int8](),
		},
		{
			Nullable: nul.Nothing[int16](),
		},
		{
			Nullable: nul.Nothing[int32](),
		},
		{
			Nullable: nul.Nothing[int64](),
		},



		{
			Nullable: nul.Nothing[uint](),
		},
		{
			Nullable: nul.Nothing[uint8](),
		},
		{
			Nullable: nul.Nothing[uint16](),
		},
		{
			Nullable: nul.Nothing[uint32](),
		},
		{
			Nullable: nul.Nothing[uint64](),
		},
	}

	for testNumber, test := range tests {

		var worked bool = false

		test.Nullable.WhenNothing(func(){

			worked = true
		})

		if !worked {
			t.Errorf("For test #%d, the call to the method did not seem work.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("NULLABLE: (%T) %#v", test.Nullable, test.Nullable)
	//////////////// CONTINUE
			continue
		}
	}
}

func TestNullable_WhenNull_something(t *testing.T) {


	tests := []struct{
		Nullable interface{WhenNothing(func())}
	}{
		{
			Nullable: nul.Null[string](),
		},



		{
			Nullable: nul.Null[int](),
		},
		{
			Nullable: nul.Null[int8](),
		},
		{
			Nullable: nul.Null[int16](),
		},
		{
			Nullable: nul.Null[int32](),
		},
		{
			Nullable: nul.Null[int64](),
		},



		{
			Nullable: nul.Null[uint](),
		},
		{
			Nullable: nul.Null[uint8](),
		},
		{
			Nullable: nul.Null[uint16](),
		},
		{
			Nullable: nul.Null[uint32](),
		},
		{
			Nullable: nul.Null[uint64](),
		},
	}

	for testNumber, test := range tests {

		var worked bool = false

		test.Nullable.WhenNothing(func(){

			worked = true
		})

		if worked {
			t.Errorf("For test #%d, the call to the method worked, but it should not have.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("NULLABLE: (%T) %#v", test.Nullable, test.Nullable)
	//////////////// CONTINUE
			continue
		}
	}
}

func TestNullable_WhenNothing_something(t *testing.T) {


	tests := []struct{
		Nullable interface{WhenNothing(func())}
	}{
		{
			Nullable: nul.Something[string]("once twice thrice fource"),
		},



		{
			Nullable: nul.Something[int](-1),
		},
		{
			Nullable: nul.Something[int8](-101),
		},
		{
			Nullable: nul.Something[int16](-10101),
		},
		{
			Nullable: nul.Something[int32](-1010101),
		},
		{
			Nullable: nul.Something[int64](-101010101),
		},



		{
			Nullable: nul.Something[uint](1),
		},
		{
			Nullable: nul.Something[uint8](101),
		},
		{
			Nullable: nul.Something[uint16](10101),
		},
		{
			Nullable: nul.Something[uint32](1010101),
		},
		{
			Nullable: nul.Something[uint64](101010101),
		},
	}

	for testNumber, test := range tests {

		var worked bool = false

		test.Nullable.WhenNothing(func(){

			worked = true
		})

		if worked {
			t.Errorf("For test #%d, the call to the method worked, but it should not have.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("NULLABLE: (%T) %#v", test.Nullable, test.Nullable)
	//////////////// CONTINUE
			continue
		}
	}
}
