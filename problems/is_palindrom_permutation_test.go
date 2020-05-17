package problems

import "testing"

func TestIsPalindromPermutation(t *testing.T) {
	testCases := []struct {
		Name     string
		S        string
		Expected bool
	}{
		{"palindrom permutation", "ababc", true},
		{"not a palindrom permutation", "abdabc", false},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := IsPalindromPermutation(tc.S)
			if actual != tc.Expected {
				t.Errorf("Expect %t, actual %t", tc.Expected, actual)
			}
		})
	}
}
