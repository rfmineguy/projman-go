package main

import (
	"fmt"
	"strings"
)

func center(s string, n int, fill string) string {
	div := n / 2;
	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
}

func abs(v int) int {
	if (v < 0) {
		return -v;
	}
	return v;
}

// Pad the shorter string with spaces
func padString(s string, length int) string {
    if len(s) >= length {
        return s
    }
    return fmt.Sprintf("%s%s", s, strings.Repeat(" ", length-len(s)))
}

// Calculate Hamming distance between two strings
func hammingDistance(s1 string, s2 string) int {
    paddedS1 := padString(s1, max(len(s1), len(s2)))
    paddedS2 := padString(s2, max(len(s1), len(s2)))

    var distance int
    for i := 0; i < len(paddedS1); i++ {
        if paddedS1[i] != paddedS2[i] {
            distance++
        }
    }
    return distance
}

// check if a contains all of b
func isSubset(subset []string, superset []string) bool {
   checkset := make(map[string]bool)
   for _, element := range subset {
      checkset[element] = true
   }
   for _, value := range superset {
      if checkset[value] {
         delete(checkset, value)
      }
   }
   return len(checkset) == 0 //this implies that set is subset of superset}
}
