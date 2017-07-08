package doc

import (
	"testing"
)

func TestExecuteLocal(t *testing.T) {
	tests := []struct {
		content string
		ops     []Op
		expect  string
	}{
		// Insert tests
		{
			content: "ab",
			ops:     []Op{&InsertOp{at: 1, insert: "x"}},
			expect:  "xab",
		},
		{
			content: "ab",
			ops:     []Op{&InsertOp{at: 2, insert: "x"}},
			expect:  "axb",
		},
		{
			content: "ab",
			ops:     []Op{&InsertOp{at: 3, insert: "x"}},
			expect:  "abx",
		},

		// Delete tests
		{
			content: "ab",
			ops:     []Op{&DeleteOp{at: 1, length: 0}},
			expect:  "ab",
		},
		{
			content: "ab",
			ops:     []Op{&DeleteOp{at: 1, length: 1}},
			expect:  "b",
		},
		{
			content: "ab",
			ops:     []Op{&DeleteOp{at: 1, length: 2}},
			expect:  "",
		},
		{
			content: "ab",
			ops:     []Op{&DeleteOp{at: 2, length: 1}},
			expect:  "a",
		},
	}

	for _, test := range tests {
		doc := New(test.content)
		for _, op := range test.ops {
			op.ExecuteLocal(doc)
		}
		if doc.String() != test.expect {
			t.Errorf("%#v.ExecuteLocal(%q) got %s; expected %s", test.ops, test.content, doc.String(), test.expect)
		}
	}
}
