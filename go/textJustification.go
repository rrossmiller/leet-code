package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	in := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	fullJustify(in, maxWidth)
}

func fullJustify(words []string, maxWidth int) []string {
	rtn := []string{words[0]}
	row := 0
	for _, v := range words[1:] {
		if utf8.RuneCountInString(rtn[row])+utf8.RuneCountInString(v) <= maxWidth {
			rtn[row] += "@" + v
		} else {
			row++
			rtn = append(rtn, v)
		}
	}
	justifyRows(rtn, maxWidth)
	os.Exit(0)
	for _, v := range rtn {
		// fmt.Println(utf8.RuneCountInString(v))
		fmt.Println(v + ",")
	}
	return rtn
}

func justifyRows(rows []string, maxWidth int) {
	for _, row := range rows {
		rowSplit := strings.Split(row, "@")
		rowLen := utf8.RuneCountInString(row)
		numSpace := maxWidth - rowLen
		idk := numSpace / len(rowSplit)
		fmt.Println(rowSplit, rowLen, numSpace, idk)

	}
}
