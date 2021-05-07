package tabular

import (
	"fmt"
	"strings"
)

func ReadTabular(s string) [][]string {
	lines := strings.Split(s, "\n")
	width := 0
	for _, c := range lines {
		if len(c) > width {
			width = len(c)
		}
	}
	spaces := make([]bool, width)
	for i := range spaces {
		spaces[i] = true
	}
	for _, l := range lines {
		for i, c := range l {
			if c != ' ' {
				spaces[i] = false
			}
		}
	}
	startSearch := true
	cols := []int{}
	for i, x := range spaces {
		if startSearch {
			if !x {
				cols = append(cols, i)
				startSearch = false
			}
		} else {
			if x {
				startSearch = true
			}
		}
	}

	rows := [][]string{}
	for _, l := range lines {
		row := []string{}
		for i, start := range cols {
			if start >= len(l) {
				if len(row) > 0 {
					row = append(row, "")
				}
				continue
			}
			end := len(l)
			if (i + 1) < len(cols) {
				end = cols[i+1]
			}
			if end > len(l) {
				end = len(l)
			}
			s := l[start:end]
			s = strings.TrimSpace(s)
			row = append(row, s)
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	return rows
}

func FormatTSV(s string) string {
	lines := strings.Split(s, "\n")
	widths := []int{}
	for _, l := range lines {
		splits := strings.Split(l, "\t")
		for i, col := range splits {
			n := len(col)
			if len(widths) <= i {
				widths = append(widths, 0)
			}
			if widths[i] < n {
				widths[i] = n
			}
		}
	}

	out := []string{}
	for _, l := range lines {
		splits := strings.Split(l, "\t")
		for i, col := range splits {
			if i != 0 {
				out = append(out, " ")
			}
			out = append(out, fmt.Sprintf("% *s", widths[i], col))
		}
		out = append(out, "\n")
	}
	return strings.Join(out, "")
}
