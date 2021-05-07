package tabular

import "testing"
import "fmt"

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

func TestFormatTSV(t *testing.T) {
	data := "a\tb\tc" + "\n" +
		"aaaa" + "\n" +
		"he\tllo\tworld\tbar" + "\n" +
		"a\tb"

	expected := `   a   b     c
aaaa
  he llo world bar
   a   b
`

	out := FormatTSV(data)
	if out != expected {
		fmt.Printf("input:\n%s\n", data)
		fmt.Printf("output:\n%s\n", out)
		fmt.Printf("expected:\n%s\n", expected)
		t.Error("bad")
	}
}

func TestCombo(t *testing.T) {
	data := "a\tb\tc" + "\n" +
		"aaaa" + "\n" +
		"he\tllo\tworld\tbar" + "\n" +
		"a\tb"

	rows := ReadTabular(FormatTSV(data))
	if rows[0][0] != "a" {
		t.Error("bad")
	}
	if rows[0][1] != "b" {
		t.Error("bad")
	}
	if rows[0][2] != "c" {
		t.Error("bad")
	}
	if rows[1][0] != "aaaa" {
		t.Error("bad")
	}
	if rows[1][1] != "" {
		t.Error("bad")
	}
	if rows[1][2] != "" {
		t.Error("bad")
	}
	if rows[2][0] != "he" {
		t.Error("bad")
	}
	if rows[2][1] != "llo" {
		t.Error("bad")
	}
	if rows[2][3] != "bar" {
		t.Error("bad")
	}
}
