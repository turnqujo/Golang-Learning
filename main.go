package main

import (
	"fmt"
	"strings"
)

func main() {
	set := "acdegilmnoprstuw"
	seed := 7
	multi := 37 // NOTE: Needs to be at least the length of the acceptable characters

	runs := []string{"a", "deg", "www", "padw", "lmnop", "awcudt", "leepadg", "asparagus"}

	for i, target := range runs {
		fmt.Printf("Run: %d\tTarget: %s\n", i+1, target)
		fmt.Printf("From Hash: %s\n\n", fromHash(toHash(target, seed, multi, set), len(target), multi, set))
	}
}

func toHash(source string, seed int, multi int, set string) int {
	hash := seed
	for i := 0; i < len(source); i++ {
		hash = hash*multi + strings.Index(set, string(source[i]))
	}
	return hash
}

// TODO: This requires the multiplier, but what if we didn't know it?
func fromHash(hash int, outputLen int, multi int, set string) string {
	charIndexes, _ := findCharIndexes(hash, len(set), multi, make([]int, outputLen), outputLen-1)

	output := ""
	for i := 0; i < len(charIndexes); i++ {
		output += string(set[charIndexes[i]])
	}

	return output
}

func findCharIndexes(hash int, setLength int, multi int, charIndexes []int, currentIndex int) ([]int, bool) {
	if currentIndex < 0 {
		return charIndexes, true
	}

	for i := 0; i < setLength; i++ {
		if (hash-i)%multi == 0 {
			charIndexes[currentIndex] = i

			if currentIndex == 0 {
				return charIndexes, true
			}

			nextHash := (hash - i) / multi
			nextIndex := currentIndex - 1
			withNext, foundNext := findCharIndexes(nextHash, setLength, multi, charIndexes, nextIndex)

			if !foundNext {
				continue
			}

			return withNext, true
		}
	}

	return charIndexes, false
}
