package euclidean_algorithm_test

import (
	"testing"

	ea "euclidean_algorithm"
)

func TestGCD(t *testing.T) {
	t.Parallel()

	cases := map[string]struct{ a, b, expected int }{
		"gcd(42, 56) = 14":              {a: 42, b: 56, expected: 14},
		"gcd(461952, 116298) = 18":      {a: 461952, b: 116298, expected: 18},
		"gcd(7966496, 314080416) = 32":  {a: 7966496, b: 314080416, expected: 32},
		"gcd(24826148, 45296490) = 526": {a: 24826148, b: 45296490, expected: 526},
		"gcd(12, 0) = 12":               {a: 12, b: 0, expected: 12},
		"gcd(0, 9) = 9":                 {a: 0, b: 9, expected: 9},
		"gcd(0, 0) = 0":                 {a: 0, b: 0, expected: 0},
		"gcd(1, 9) = 1":                 {a: 1, b: 9, expected: 1},
		"gcd(9, 1) = 1":                 {a: 9, b: 1, expected: 1},
		"gcd(-42, 56) = 14":             {a: -42, b: 56, expected: 14},
		"gcd(42, -56) = 14":             {a: 42, b: -56, expected: 14},
		"gcd(-42, -56) = 14":            {a: -42, b: -56, expected: 14},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := ea.GCD(tc.a, tc.b)
			if actual != tc.expected {
				t.Logf("FAIL: %s:\n\tactual:   %v\n\texpected: %v", t.Name(), actual, tc.expected)
				t.Fail()
			}
		})
	}
}
