package cleaner

import "testing"

func TestClean_PySBDBasics(t *testing.T) {
	c := NewCleaner()
	tests := []struct {
		in, want string
	}{
		{"It was a cold \nnight in the city.", "It was a cold night in the city."},
		{"This is the U.S. Senate my friends. <em>Yes.</em> <em>It is</em>!", "This is the U.S. Senate my friends. Yes. It is!"},
		{"This is &lt;strong&gt;bold&lt;/strong&gt;.", "This is bold."},
		{
			"&lt;TITLE&gt;Heading&lt;/TITLE&gt; &lt;img src=&quot;logo.png&quot;/&gt;Next.",
			"Heading Next.",
		},
		{
			"&lt;span data-id=&quot;1&quot;&gt;A &amp; B&lt;/span&gt;&lt;em&gt;C&lt;/em&gt;",
			"A &amp; BC",
		},
		{"Keep &lt; 5 &amp; 6 &gt; 2.", "Keep &lt; 5 &amp; 6 &gt; 2."},
		{"\nW\nA\nRN\nI\nNG\n", "WARNING\r"},
		{"Hello world.Today is Tuesday.", "Hello world. Today is Tuesday."},
		{
			"Contact Jane.Doe@example.com about Jane.Doe today.",
			"Contact Jane.Doe@example.com about Jane. Doe today.",
		},
		{
			"Visit https://example.com/foo.Bar or type foo.Bar now.",
			"Visit https://example.com/foo.Bar or type foo. Bar now.",
		},
		{
			"  Type 1.About  or visit https://example.com/1.About  now.  ",
			"  Type 1. About  or visit https://example.com/1.About  now.  ",
		},
		{"", ""},
	}
	for _, tt := range tests {
		if got := c.Clean(tt.in); got != tt.want {
			t.Errorf("Clean(%q)=%q want %q", tt.in, got, tt.want)
		}
	}
}
