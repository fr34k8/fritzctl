package console

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTableGeneration asserts that printing of table table data matches the expected format.
func TestTableGeneration(t *testing.T) {
	tcs := []struct {
		name     string
		table    *Table
		expected string
	}{
		{
			name:  "without_body",
			table: NewTable(Headers("a", "b", "c")),
			expected: `+---+---+---+
| a | b | c |
+---+---+---+
+---+---+---+
`,
		},
		{
			name:  "one_line_body",
			table: NewTable(Headers("a", "b", "c"), body([][]string{{"x", "y", "z"}})),
			expected: `+---+---+---+
| a | b | c |
+---+---+---+
| x | y | z |
+---+---+---+
`,
		},
		{
			name: "with_padding",
			table: NewTable(Headers("NAME", "IP", "MAC"),
				body([][]string{
					{"host1", "192.168.12.23", "AA:BB:CC:DD:EE:FF"},
					{"host2", "192.168.12.102", "99:99:99:99:99:99"},
					{"fritz.box", "192.168.12.1", "11:22:33:44:55:66"},
				})),
			expected: `+-----------+----------------+-------------------+
|   NAME    |       IP       |        MAC        |
+-----------+----------------+-------------------+
| host1     | 192.168.12.23  | AA:BB:CC:DD:EE:FF |
| host2     | 192.168.12.102 | 99:99:99:99:99:99 |
| fritz.box | 192.168.12.1   | 11:22:33:44:55:66 |
+-----------+----------------+-------------------+
`,
		},
		{
			name: "with_alignment",
			table: NewTable(Headers("NAME", "IP", "MAC", "ACT/ONL", "SPEED [MBIT/S]"),
				body([][]string{
					{"host1", "192.168.12.23", "AA:BB:CC:DD:EE:FF", "✔/✘", "104"},
					{"host2", "192.168.12.102", "99:99:99:99:99:99", "✘/✘", "0"},
					{"fritz.box", "192.168.12.1", "11:22:33:44:55:66", "✔/✘", "0"},
				})),
			expected: `+-----------+----------------+-------------------+---------+----------------+
|   NAME    |       IP       |        MAC        | ACT/ONL | SPEED [MBIT/S] |
+-----------+----------------+-------------------+---------+----------------+
| host1     | 192.168.12.23  | AA:BB:CC:DD:EE:FF | ✔/✘     |            104 |
| host2     | 192.168.12.102 | 99:99:99:99:99:99 | ✘/✘     |              0 |
| fritz.box | 192.168.12.1   | 11:22:33:44:55:66 | ✔/✘     |              0 |
+-----------+----------------+-------------------+---------+----------------+
`,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assertions := assert.New(t)
			assertions.NotNil(tc.table)
			buf := new(bytes.Buffer)
			tc.table.Print(buf)
			assertions.Equal(tc.expected, buf.String())
		})
	}
}

func body(rows [][]string) Option {
	return func(t *Table) {
		for _, r := range rows {
			t.Append(r)
		}
	}
}
