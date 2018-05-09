// Package diff implements a simple diff algorithm ported from
// https://github.com/paulgb/simplediff (Python version)
//
// References:
//   https://www.google.com/search?q=line+diff+algorithm
//   https://stackoverflow.com/questions/805626/diff-algorithm#comment19356828_805626
//   http://paulbutler.org/archives/a-simple-diff-algorithm-in-php/
package diff

type DiffPair struct {
	// "+" (insertion) or "-" (deletion) or "=" (no change)
	typ string
	// items in *before* or *after* slice
	items []string
}

// Find the differences between two string slices/arrays. Returns a slice of
// pairs. The first value of the pair indicates "+" (insertion), "-" (deletion),
// or "=" (no change). The second value of pair is the items of input
// slices/arrays.
func Diff(before, after []string) (pairs []DiffPair) {
	beforeIndexMap := make(map[string][]int)
	for i, str := range before {
		if _, ok := beforeIndexMap[str]; ok {
			beforeIndexMap[str] = append(beforeIndexMap[str], i)
		} else {
			beforeIndexMap[str] = []int{i}
		}
	}

	// Fine longest common subsequence (LCS)
	var overlap map[int]int
	subStartBefore := 0
	subStartAfter := 0
	subLength := 0

	for iafter, str := range after {
		overlap_ := make(map[int]int)
		if ids, ok := beforeIndexMap[str]; ok {
			for _, ibefore := range ids {
				if val, ok := overlap[ibefore-1]; ok {
					if ibefore > val {
						overlap_[ibefore] = ibefore + 1
					} else {
						overlap_[ibefore] = val + 1
					}
				} else {
					overlap_[ibefore] = ibefore + 1
				}

				if overlap_[ibefore] > subLength {
					// this is the largest substring seen so
					// far, so store its indices
					subLength = overlap_[ibefore]
					subStartBefore = ibefore - subLength + 1
					subStartAfter = iafter - subLength + 1
				}
			}
			overlap = overlap_
		}
	}

	if subLength == 0 {
		if len(before) > 0 {
			pairs = append(pairs, DiffPair{typ: "-", items: before})
		}
		if len(after) > 0 {
			pairs = append(pairs, DiffPair{typ: "+", items: after})
		}
	} else {
	}

	return
}
