package main

import "testing"

func TestReverseString(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"short", "short"},
		{"this is a test", "test a is this"},
		{"this is a really long sentence that should hopefully not break anything and is going to be 120 words long but everything should still be fine and will repeat again this is a really long sentence that should hopefully not break anything and is going to be 120 words long but everything should still be fine and will repeat again this is a really long sentence that should hopefully not break anything and is going to be 120 words long but everything should still be fine and will repeat again this is a really long sentence that should hopefully not break anything and is going to be 120 words long but everything should still be fine and will repeat again",
			"again repeat will and fine be still should everything but long words 120 be to going is and anything break not hopefully should that sentence long really a is this again repeat will and fine be still should everything but long words 120 be to going is and anything break not hopefully should that sentence long really a is this again repeat will and fine be still should everything but long words 120 be to going is and anything break not hopefully should that sentence long really a is this again repeat will and fine be still should everything but long words 120 be to going is and anything break not hopefully should that sentence long really a is this"},
	}
	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			reverse := reverseString(test.in)
			if reverse != test.out {
				t.Errorf("expected %s but got %s", test.out, reverse)
			}
		})
	}
}
