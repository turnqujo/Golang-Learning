package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Goal: Unhash a 9 character string from 910897038977002
func main() {
	set := "acdegilmnoprstuw"

	runs := []string{"w", "ww", "www", "wwww", "wwwww", "wwwwww"}

	for i, target := range runs {
		fmt.Printf("Run: %d\tTarget: %s\n", i + 1, target)
		fmt.Printf("From Hash: %s\n\n", fromHash(toHash(target, 7, 37, set), len(target), 7, 37, set))
	}
}

func toHash(source string, seed int, multi int, set string) int {
	hash := seed
	for i := 0; i < len(source); i++ {
		hash = hash*multi + strings.Index(set, string(source[i]))
	}
	return hash
}

func fromHash(hash int, outputLen int, seed int, multi int, set string) string {
	start := time.Now()
	possibleCombinations := math.Pow(float64(len(set)), float64(outputLen))
	fmt.Printf("Source: %d\nPossible Combinations: %d\nStarting at: %s\n", hash, int(possibleCombinations), start.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))

	foundChan := make(chan string)
	stopChan := make(chan struct{})

	for i := 0; i < int(possibleCombinations); i++ {
		select {
		case <-stopChan:
			break
		default:
			go asyncGuess(foundChan, stopChan, hash, i, seed, multi, outputLen, set)
		}
	}

	fmt.Printf("Done in: %s\n", time.Since(start))
	return <-foundChan
}

func asyncGuess(foundChan chan string, stopChan chan struct{}, hash int, attempt int, seed int, multi int, outputLen int, set string) {
	for {
		select {
		case <-stopChan:
			return
		default:
			guess := generateGuess(attempt, outputLen, set)
			guessHash := toHash(guess, seed, multi, set)

			if guessHash != hash {
				return
			}

			close(stopChan)

			foundChan <- guess
			close(foundChan)
		}
	}
}

func generateGuess(attempt int, outputLen int, set string) string {
	rawGuess := make([]int, outputLen)

	for i := attempt; i > 0; i-- {
		rawGuess = allocatePoint(rawGuess, len(set), 0)
	}

	guess := ""
	for i := 0; i < len(rawGuess); i++ {
		guess += string(set[rawGuess[i]])
	}

	return guess
}

// NOTE: Recursive; will increment the character number in the guess sequence from left to right
func allocatePoint(rawGuess []int, maxAmount int, offset int) []int {
	rawGuess[offset]++
	if rawGuess[offset] >= maxAmount {
		rawGuess[offset] = 0
		newOffset := offset + 1

		if newOffset > len(rawGuess) {
			return rawGuess
		}

		rawGuess = allocatePoint(rawGuess, maxAmount, newOffset)
	}

	return rawGuess
}
