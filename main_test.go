package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// Define TestCases with various test cases.
var TestCases = []struct {
	name   string
	files  []string
	op     string
	col    int
	exp    string
	expErr error
}{
	{
		name:   "Test Mean",
		files:  []string{"housesInput.csv"},
		op:     "mean",
		col:    2,
		exp:    "3.8706710029070246\n",
		expErr: nil,
	},
}

func TestRun(t *testing.T) {

	// Define a tolerance for floating-point comparisons.
	const tolerance = .0000000000001

	for i := 0; i < 100; i++ {
		for _, tc := range TestCases {
			t.Run(fmt.Sprintf("%s_%d", tc.name, i), func(t *testing.T) {
				var res bytes.Buffer
				err := run(tc.files, tc.op, tc.col, &res)

				// Convert the expected result and the actual result to float64 for comparison.
				expected, _ := strconv.ParseFloat(tc.exp, 64)
				if err != nil {
					t.Errorf("Invalid expected value: %v", err)
					return
				}

				actual, _ := strconv.ParseFloat(res.String(), 64)
				if err != nil {
					t.Errorf("Invalid actual value: %v", err)
					return
				}

				if tc.expErr != nil {
					if err == nil {
						t.Errorf("Expected error. Got nil instead")
					}

					if !strings.Contains(err.Error(), tc.expErr.Error()) {
						t.Errorf("Expected error %q. Got %q instead", tc.expErr, err)
					}

					return
				}

				if err != nil {
					t.Errorf("Unexpected error %q", err)
					return
				}

				if math.Abs(actual-expected) > tolerance {
					t.Errorf("Expected %v, but got %v instead", expected, actual)
				}

				// if res.String() != tc.exp {
				// 	t.Errorf("Expected %q. Got %q instead", tc.exp, res.String())
				// }
			})
		}
	}
}
