package problems

import "testing"

func TestCheckPermutation(t *testing.T) {
	testCases := []struct {
		Name     string
		S1       string
		S2       string
		Expected bool
	}{
		{"permutation", "abcd", "dbac", true},
		{"different length", "abcdf", "dbac", false},
		{"not a permutation", "abcdf", "dbakc", false},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := CheckPermutation(tc.S1, tc.S2)
			if actual != tc.Expected {
				t.Errorf("Expect %t, actual %t", tc.Expected, actual)
			}
		})
	}
}

func TestCheckPermutationUtf8(t *testing.T) {
	testCases := []struct {
		Name     string
		S1       string
		S2       string
		Expected bool
	}{
		{"permutation", "абєшУ", "бєаУш", true},
		{"different length", "БрушЗ", "БруФшЗ", false},
		{"not a permutation", "биухн", "бицхн", false},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := CheckPermutation(tc.S1, tc.S2)
			if actual != tc.Expected {
				t.Errorf("Expect %t, actual %t", tc.Expected, actual)
			}
		})
	}
}
