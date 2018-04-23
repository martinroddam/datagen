package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatName_Basic(t *testing.T) {
	expected := "DR. YURI ZHIVAGO"
	pre := "Dr."
	first := "Yuri"
	last := "Zhivago"

	actual := formatName(pre, first, last)

	assert.Equal(t, expected, actual)
}

func TestFormatName_TableDriven(t *testing.T) {
	cases := []struct {
		Name, Prefix, First, Last, Expected string
	}{
		{"All lower", "mr", "david", "smith", "MR DAVID SMITH"},
		{"All upper", "MRS", "ANDREA", "JONES", "MRS ANDREA JONES"},
		{"Special characters", "MR", "GARY", "O'NEILL", "MR GARY O'NEILL"},
		{"Trimmed whitespace", " MISS ", "  SARAH JANE ", "PETERS   \n", "MISS SARAH JANE PETERS"},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := formatName(tc.Prefix, tc.First, tc.Last)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
