package counter_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mercadolibre/advanced-go/2018-05-advanced-go/day2/example/mocking"
	"github.com/stretchr/testify/require"
)

func TestCountReader(t *testing.T) {
	tt := []struct {
		Name     string
		Input    string
		Expected int64
	}{
		{
			Name:     "Empty String",
			Input:    "",
			Expected: 0,
		},
		{
			Name:     "One line",
			Input:    "This is my one and only line",
			Expected: 1,
		},
		{
			Name:     "Three Lines",
			Input:    "This is my first\nsecond\nand third lines",
			Expected: 2,
		},
		{
			Name:     "Four Continuos Lines",
			Input:    "This is my first\n\n\n\nand last lines",
			Expected: 5,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			r := bytes.NewReader([]byte(tc.Input))
			out := counter.CountReader(r)

			require.Equal(t, tc.Expected, out)
		})
	}
}

func ExampleCountReader() {
	r := bytes.NewReader([]byte("a\n\nb"))
	out := counter.CountReader(r)

	fmt.Printf("found %d lines", out)
	// Output: found 2 lines
}
