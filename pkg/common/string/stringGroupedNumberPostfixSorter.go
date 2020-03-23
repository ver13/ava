package string

import (
	"strconv"
)

type StringGroupedNumberPostfixSorter []string

// Len is the number of elements in the collection.
func (s StringGroupedNumberPostfixSorter) Len() int {
	return len(s)
}

// Less reports whether the element with index i should sort before the element with index j.
func (s StringGroupedNumberPostfixSorter) Less(i, j int) bool {
	bi, ni := stringSplitNumberPostfix(s[i])
	bj, nj := stringSplitNumberPostfix(s[j])

	if bi == bj {
		if len(ni) == len(nj) {
			inti, _ := strconv.Atoi(ni)
			intj, _ := strconv.Atoi(nj)
			return inti < intj
		} else {
			return len(ni) < len(nj)
		}
	}

	return bi < bj
}

// Swap swaps the elements with indexes i and j.
func (s StringGroupedNumberPostfixSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func stringSplitNumberPostfix(str string) (base, number string) {
	return StringSplitNumberPostfix(str)
}
