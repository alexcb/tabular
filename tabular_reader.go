package tabular

import (
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
