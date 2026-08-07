package cleaner

import "testing"

func TestClean_PySBDBasics(t *testing.T) {
	c := NewCleaner()
	tests := []struct {
		in, want string
	}{
		{"It was a cold \nnight in the city.", "It was a cold night in the city."},
		{"This is the U.S. Senate my friends. <em>Yes.</em> <em>It is</em>!", "This is the U.S. Senate my friends. Yes. It is!"},
		{"\nW\nA\nRN\nI\nNG\n", "WARNING\r"},
		{"Hello world.Today is Tuesday.", "Hello world. Today is Tuesday."},
		{"", ""},
	}
	for _, tt := range tests {
		if got := c.Clean(tt.in); got != tt.want {
			t.Errorf("Clean(%q)=%q want %q", tt.in, got, tt.want)
		}
	}
}
