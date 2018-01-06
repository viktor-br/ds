package problems

import "testing"

func TestStringCompressionUtf8(t *testing.T) {
	s := "ЛЛЛллЦЦЇЇЇ"
	actual := StringCompressionUtf8(s)
	expected := "Л3л2Ц2Ї3"

	if actual != expected {
		t.Errorf("Expect %s, actual %s", expected, actual)
	}

	s = "абвгд"
	actual = StringCompressionUtf8(s)
	expected = "абвгд"

	if actual != expected {
		t.Errorf("Expect %s, actual %s", expected, actual)
	}

	s = ""
	actual = StringCompressionUtf8(s)
	expected = ""

	if actual != expected {
		t.Errorf("Expect %s, actual %s", expected, actual)
	}
}
