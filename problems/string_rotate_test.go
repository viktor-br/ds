package problems

import "testing"

func TestIsStringRotate(t *testing.T) {
	testCases := []struct {
		Name     string
		Initial  string
		Rotated  string
		Expected bool
	}{
		{"rotated", "waterbottle", "erbottlewat", true},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := IsStringRotate(tc.Initial, tc.Rotated)
			if actual != tc.Expected {
				t.Errorf("Expect %t, actual %t", tc.Expected, actual)
			}
		})
	}
}
