package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	set := "acdegilmnoprstuw"

	result := fromHash(toHash("gato", 7, 37, set), 4, 7, 37, set)
	// result := fromHash(910897038977002, 9, 7, 37, set)
	fmt.Printf("From Hash: %s\n", result)
}

func toHash(source string, seed int, multi int, set string) int {
	hash := seed
	for i := 0; i < len(source); i++ {
		hash = hash*multi + strings.Index(set, string(source[i]))
	}
	return hash
}

/**
 * TODO:
 *  - This approach takes quite a lot of time when the outputLen is greater than 5.
 */
func fromHash(hash int, outputLen int, seed int, multi int, set string) string {
	possibleCombinations := math.Pow(float64(len(set)), float64(outputLen))
	fmt.Printf("Source: %d\tPossible Combinations: %d\n", hash, int(possibleCombinations))

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
