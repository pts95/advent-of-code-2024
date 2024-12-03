package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	flagLoc := flag.String("f", "", "file location (absolute path)")
	flag.Parse()

	if flagLoc == nil || *flagLoc == "" {
		flag.Usage()
		return
	}

	a, b := parseInputs(*flagLoc)
	distance := compareDistances(a, b)
	fmt.Printf("current distance: %d (size %d)\n", distance, len(a))
	similarityScore := computeSimilarityScore(a, b)
	fmt.Printf("similarity score: %d (size %d)\n", similarityScore, len(a))
}

func parseInputs(path string) (a []int64, b []int64) {
	f, err := os.Open(path)
	if err != nil {
		log.Panicf("couldn't open file %s: %s", path, err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		var aint, bint int64
		if _, err := fmt.Sscanf(text, "%d %d", &aint, &bint); err != nil {
			log.Printf("couldn't parse %s: %s", text, err)
			continue
		}
		a = append(a, aint)
		b = append(b, bint)
	}
	return a, b
}

func compareDistances(a []int64, b []int64) int64 {
	sortInts(a)
	sortInts(b)
	var (
		distance int64
		curval   int64
	)
	for i := 0; i < len(a); i++ {
		curval = distance + abs(a[i]-b[i])
		if curval < distance {
			log.Panicf("overflow detected at index %d: current distance %d, left: %d, right: %d",
				i, distance, a[i], b[i])
		}
		distance = curval
	}
	return distance
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func numOccurences(b []int64) map[int64]int64 {
	occurences := make(map[int64]int64)
	for i := 0; i < len(b); i++ {
		occurences[b[i]]++
	}
	return occurences
}

func sortInts(a []int64) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}

func computeSimilarityScore(a []int64, b []int64) int64 {
	occurrences := numOccurences(b)
	var similarityScore int64
	var curval int64
	for i := 0; i < len(a); i++ {
		curval = similarityScore + a[i]*occurrences[a[i]]
		if curval < similarityScore {
			log.Panicf("potential overflow at index %d: current similarityScore: %d, left: %d, right: %d",
				i, similarityScore, a[i], b[i])
		}
		similarityScore = curval
	}
	return similarityScore
}
