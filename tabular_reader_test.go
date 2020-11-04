package tabular

import "testing"

func TestRead(t *testing.T) {
	data := `
a        b         c
12.43             4.94
12.43    10.20   4.94
`
	rows := ReadTabular(data)
	if len(rows) != 3 {
		t.Error("bad")
	}
	for _, row := range rows {
		if len(row) != 3 {
			t.Error("bad")
		}
	}
}
