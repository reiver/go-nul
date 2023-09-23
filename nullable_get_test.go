package nul_test

import (
	"testing"

	"sourcecode.social/reiver/go-nul"
)

func TestOptional_Get_string(t *testing.T) {

	tests := []struct{
		Nullable nul.Nullable[string]
		ExpectedValue     string
		ExpectedSomething bool
	}{
		{
			Nullable: nul.Nothing[string](),
			ExpectedValue: "",
			ExpectedSomething: false,
		},



		{
			Nullable: nul.Something(""),
			ExpectedValue:          "",
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something("once twice thrice fource"),
			ExpectedValue:          "once twice thrice fource",
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something("😈"),
			ExpectedValue:          "😈",
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something("۰۱۲۳۴۵۶۷۸۹"),
			ExpectedValue:          "۰۱۲۳۴۵۶۷۸۹",
			ExpectedSomething: true,
		},
	}

	for testNumber, test := range tests {

		value, something := test.Nullable.Get()

		{
			expected := test.ExpectedSomething
			actual   := something

			if expected != actual {
				t.Errorf("For test #%d, the actual something-flag is not what was expected.", testNumber)
				t.Logf("EXPECTED FLAG: %t", expected)
				t.Logf("ACTUAL   FLAG: %t", actual)
				t.Logf("NULLABLE: %#v", test.Nullable)
	/////////////////////// CONTINUE
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := value

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("NULLABLE: %#v", test.Nullable)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_Get_int8(t *testing.T) {

	tests := []struct{
		Nullable nul.Nullable[int8]
		ExpectedValue     int8
		ExpectedSomething bool
	}{
		{
			Nullable: nul.Nothing[int8](),
			ExpectedValue: 0,
			ExpectedSomething: false,
		},



		{
			Nullable: nul.Something(int8(-127)),
			ExpectedValue:          int8(-127),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-126)),
			ExpectedValue:          int8(-126),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-125)),
			ExpectedValue:          int8(-125),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-124)),
			ExpectedValue:          int8(-124),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-123)),
			ExpectedValue:          int8(-123),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-122)),
			ExpectedValue:          int8(-122),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-121)),
			ExpectedValue:          int8(-121),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-120)),
			ExpectedValue:          int8(-120),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-9)),
			ExpectedValue:          int8(-9),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-8)),
			ExpectedValue:          int8(-8),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-7)),
			ExpectedValue:          int8(-7),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-6)),
			ExpectedValue:          int8(-6),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-5)),
			ExpectedValue:          int8(-5),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-4)),
			ExpectedValue:          int8(-4),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-3)),
			ExpectedValue:          int8(-3),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-2)),
			ExpectedValue:          int8(-2),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(-1)),
			ExpectedValue:          int8(-1),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(0)),
			ExpectedValue:          int8(0),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(1)),
			ExpectedValue:          int8(1),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(2)),
			ExpectedValue:          int8(2),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(3)),
			ExpectedValue:          int8(3),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(4)),
			ExpectedValue:          int8(4),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(5)),
			ExpectedValue:          int8(5),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(6)),
			ExpectedValue:          int8(6),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(7)),
			ExpectedValue:          int8(7),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(8)),
			ExpectedValue:          int8(8),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(9)),
			ExpectedValue:          int8(9),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(120)),
			ExpectedValue:          int8(120),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(121)),
			ExpectedValue:          int8(121),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(122)),
			ExpectedValue:          int8(122),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(123)),
			ExpectedValue:          int8(123),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(124)),
			ExpectedValue:          int8(124),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(125)),
			ExpectedValue:          int8(125),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(126)),
			ExpectedValue:          int8(126),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(int8(127)),
			ExpectedValue:          int8(127),
			ExpectedSomething: true,
		},
	}

	for testNumber, test := range tests {

		value, something := test.Nullable.Get()

		{
			expected := test.ExpectedSomething
			actual   := something

			if expected != actual {
				t.Errorf("For test #%d, the actual something-flag is not what was expected.", testNumber)
				t.Logf("EXPECTED FLAG: %t", expected)
				t.Logf("ACTUAL   FLAG: %t", actual)
				t.Logf("NULLABLE: %#v", test.Nullable)
	/////////////////////// CONTINUE
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := value

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("NULLABLE: %#v", test.Nullable)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_Get_uint8(t *testing.T) {

	tests := []struct{
		Nullable nul.Nullable[uint8]
		ExpectedValue     uint8
		ExpectedSomething bool
	}{
		{
			Nullable: nul.Nothing[uint8](),
			ExpectedValue: 0,
			ExpectedSomething: false,
		},



		{
			Nullable: nul.Something(uint8(0)),
			ExpectedValue:          uint8(0),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(uint8(1)),
			ExpectedValue:          uint8(1),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(uint8(2)),
			ExpectedValue:          uint8(2),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(uint8(254)),
			ExpectedValue:          uint8(254),
			ExpectedSomething: true,
		},
		{
			Nullable: nul.Something(uint8(255)),
			ExpectedValue:          uint8(255),
			ExpectedSomething: true,
		},
	}

	for testNumber, test := range tests {

		value, something := test.Nullable.Get()

		{
			expected := test.ExpectedSomething
			actual   := something

			if expected != actual {
				t.Errorf("For test #%d, the actual something-flag is not what was expected.", testNumber)
				t.Logf("EXPECTED FLAG: %t", expected)
				t.Logf("ACTUAL   FLAG: %t", actual)
				t.Logf("NULLABLE: %#v", test.Nullable)
	/////////////////////// CONTINUE
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := value

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("NULLABLE: %#v", test.Nullable)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}
