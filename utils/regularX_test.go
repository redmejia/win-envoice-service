package utils

import "testing"

func TestRegularX(t *testing.T) {
	src := "1111222233334444"
	want := "****-****-****-4444"

	ReplChars(&src, `\d{12}`, "****-****-****-")

	if src != want {
		t.Fatalf("want %s but source output is %s", want, src)
	}

}
