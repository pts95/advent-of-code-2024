package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flagLoc := flag.String("f", "", "file location (absolute path)")
	flag.Parse()

	if flagLoc == nil || *flagLoc == "" {
		flag.Usage()
		return
	}

	safeLevels := countSafeLevels(*flagLoc)
	fmt.Printf("number of safe levels: %d\n", safeLevels)
	safeLevelsWithTolerances := countSafeLevelsWithTolerances(*flagLoc)
	fmt.Printf("number of safe levels with tolerances: %d\n", safeLevelsWithTolerances)
}

func countSafeLevels(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var numSafeLevels int
	for scanner.Scan() {
		text := scanner.Text()
		allElements := parseInts(strings.Split(text, " "))
		if len(allElements) == 1 {
			numSafeLevels++
			continue
		}
		if isLevelSafe, _, _ := isLevelCompletelySafe(allElements); isLevelSafe {
			numSafeLevels++
		}
	}
	return numSafeLevels
}

func countSafeLevelsWithTolerances(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var numSafeLevels int
	var rowindex int
	for scanner.Scan() {
		rowindex++
		text := scanner.Text()
		allElements := parseInts(strings.Split(text, " "))
		if len(allElements) <= 2 {
			numSafeLevels++
			continue
		}
		if isLevelCompletelySafeWithTolerance(allElements) {
			numSafeLevels++
		}
	}
	return numSafeLevels
}

func isLevelCompletelySafe(allElements []int) (bool, int, error) {
	for i := 1; i < len(allElements); i++ {
		if safe, err := isPairSafe(allElements[i-1], allElements[i], allElements[1]-allElements[0] > 0); !safe {
			return false, i, err
		}
	}
	return true, -1, nil
}

func isLevelCompletelySafeWithTolerance(allElements []int) bool {
	isLevelSafe, idxWhereUnsafeFirst, _ := isLevelCompletelySafe(allElements)
	if isLevelSafe {
		return true
	}
	removingSecondIdx := arrayExcludingIndex(allElements, idxWhereUnsafeFirst)
	isSafeByRemovingSecondIdx, _, _ := isLevelCompletelySafe(removingSecondIdx)
	removingFirstIdx := arrayExcludingIndex(allElements, idxWhereUnsafeFirst-1)
	isSafeByRemovingFirstIdx, _, _ := isLevelCompletelySafe(removingFirstIdx)
	if idxWhereUnsafeFirst == 2 {
		// The base case might be wrong to begin with. Adding a base case for 0 to cover for that.
		removingBaseIdx := arrayExcludingIndex(allElements, 0)
		isSafeRemovingBaseIdx, _, _ := isLevelCompletelySafe(removingBaseIdx)
		return isSafeRemovingBaseIdx || isSafeByRemovingFirstIdx || isSafeByRemovingSecondIdx
	}
	return isSafeByRemovingFirstIdx || isSafeByRemovingSecondIdx
}

func isPairSafe(previousNumber int, currentNumber int, isPositive bool) (bool, error) {
	difference := currentNumber - previousNumber
	if difference == 0 {
		return false, errors.New("equal elements detected")
	}
	if isPositive {
		return difference >= 1 && difference <= 3, fmt.Errorf("difference outside positive range: %d", difference)
	}
	return difference >= -3 && difference <= -1, fmt.Errorf("difference outside negative range: %d", difference)
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func arrayExcludingIndex(a []int, idx int) []int {
	var toReturn []int
	if idx > 0 {
		toReturn = append(toReturn, a[:idx]...)
	}
	if idx < len(a)-1 {
		toReturn = append(toReturn, a[idx+1:]...)
	}
	return toReturn
}

func parseInts(s []string) []int {
	var toReturn []int
	for _, str := range s {
		toReturn = append(toReturn, parseInt(str))
	}
	return toReturn
}
