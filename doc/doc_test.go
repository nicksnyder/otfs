package doc

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		content string
		insert  string
		at      int
		expect  string
	}{
		{
			content: "ab",
			insert:  "x",
			at:      1,
			expect:  "xab",
		},
		{
			content: "ab",
			insert:  "x",
			at:      2,
			expect:  "axb",
		},
		{
			content: "ab",
			insert:  "x",
			at:      3,
			expect:  "abx",
		},
	}

	for _, test := range tests {
		d := New(test.content)
		d.Ins(test.at, test.insert)
		if d.String() != test.expect {
			t.Errorf("%q.Ins(%d, %s) got %s; expected %s", test.content, test.at, test.insert, d.String(), test.expect)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		content string
		at      int
		length  int
		expect  string
	}{
		{
			content: "ab",
			at:      1,
			length:  0,
			expect:  "ab",
		},
		{
			content: "ab",
			at:      1,
			length:  1,
			expect:  "b",
		},
		{
			content: "ab",
			at:      1,
			length:  2,
			expect:  "",
		},
		{
			content: "ab",
			at:      2,
			length:  1,
			expect:  "a",
		},
	}

	for _, test := range tests {
		d := New(test.content)
		d.Del(test.at, test.length)
		if d.String() != test.expect {
			t.Errorf("%q.Del(%d, %d) got %s; expected %s", test.content, test.at, test.length, d.String(), test.expect)
		}
	}
}
