package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

func naive() {
	os.Chdir("./cmd/dayOne")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteEntries := bytes.Split(data, []byte("\n"))
	entries := make([]int, 0, len(byteEntries))

	for _, byteEntry := range byteEntries {
		if len(byteEntry) != 0 {
			entry, err := strconv.Atoi(string(byteEntry))
			if err != nil {
				log.Fatal(err)
			}
			entries = append(entries, entry)
		}
	}

	for x := 0; x < len(entries)-2; x++ {
		a := entries[x]
		for y := x + 1; y < len(entries)-1; y++ {
			b := entries[y]
			for z := y + 1; z < len(entries); z++ {
				c := entries[z]
				if a+b+c == 2020 {
					fmt.Printf("%d + %d + %d = 2020\n%d * %d * %d = %d\n", a, b, c, a, b, c, a*b*c)
					return
				}
			}
		}
	}

	fmt.Println("No three entries sum to 2020")
}

func optimized() {
	os.Chdir("./cmd/dayOne")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteEntries := bytes.Split(data, []byte("\n"))
	entries := make([]int, 0, len(byteEntries))

	for _, byteEntry := range byteEntries {
		if len(byteEntry) != 0 {
			entry, err := strconv.Atoi(string(byteEntry))
			if err != nil {
				log.Fatal(err)
			}
			insertAt := sort.Search(len(entries), func(i int) bool { return entries[i] >= entry })
			entries = append(entries, 0)
			copy(entries[insertAt+1:], entries[insertAt:])
			entries[insertAt] = entry
		}
	}

	aTarget := 2020
	aBound := aTarget / 3
	for x := 0; x < len(entries); x++ {
		a := entries[x]
		if a > aBound {
			break
		}
		bTarget := aTarget - a
		bBound := bTarget / 2
		for y := x + 1; y < len(entries); y++ {
			b := entries[y]
			if b > bBound {
				break
			}
			cTarget := bTarget - b
			remaining := entries[y+1:]
			cTargetIndex := sort.Search(len(remaining), func(i int) bool { return remaining[i] >= cTarget })
			if cTargetIndex != len(remaining) && remaining[cTargetIndex] == cTarget {
				fmt.Printf("%d + %d + %d = 2020\n%d * %d * %d = %d\n", a, b, cTarget, a, b, cTarget, a*b*cTarget)
				return
			}
		}
	}

	fmt.Println("No three entries sum to 2020")
}

func findSummingItems(items []int, sum int, n int) ([]int, bool) {
	if n == 1 {
		index := sort.Search(len(items), func(i int) bool { return items[i] >= sum })
		if index != len(items) && items[index] == sum {
			return []int{sum}, true
		}
		return nil, false
	}
	bound := sum / n
	for i := 0; i < len(items); i++ {
		curr := items[i]
		if curr > bound {
			return nil, false
		}
		if next, ok := findSummingItems(items[i+1:], sum-curr, n-1); ok {
			return append(next, curr), ok
		}
	}
	return nil, false
}

func optimizedRecursive() {
	os.Chdir("./cmd/dayOne")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteEntries := bytes.Split(data, []byte("\n"))
	entries := make([]int, 0, len(byteEntries))

	for _, byteEntry := range byteEntries {
		if len(byteEntry) != 0 {
			entry, err := strconv.Atoi(string(byteEntry))
			if err != nil {
				log.Fatal(err)
			}
			insertAt := sort.Search(len(entries), func(i int) bool { return entries[i] >= entry })
			entries = append(entries, 0)
			copy(entries[insertAt+1:], entries[insertAt:])
			entries[insertAt] = entry
		}
	}

	if found, ok := findSummingItems(entries, 2020, 3); ok {
		fmt.Printf("%d + %d + %d = 2020\n%d * %d * %d = %d\n", found[0], found[1], found[2], found[0], found[1], found[2], found[0]*found[1]*found[2])
		return
	}
	fmt.Println("No three entries sum to 2020")
}

func main() {
	// naive()
	// optimized()
	optimizedRecursive()
}
