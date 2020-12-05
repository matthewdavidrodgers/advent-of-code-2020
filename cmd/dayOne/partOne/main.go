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

	for i := 0; i < len(entries); i++ {
		target := entries[i]
		for j := i + 1; j < len(entries); j++ {
			attempt := entries[j]
			if target+attempt == 2020 {
				fmt.Printf("%d + %d = 2020\n%d * %d = %d\n", target, attempt, target, attempt, target*attempt)
				return
			}
		}
	}

	fmt.Println("No two entries sum to 2020")
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

	for i := 0; i < len(entries); i++ {
		target := entries[i]
		if target > 1010 {
			break
		}
		remaining := entries[i+1:]
		match := 2020 - target
		matchIndex := sort.Search(len(remaining), func(i int) bool { return remaining[i] >= match })
		if matchIndex != len(remaining) && remaining[matchIndex] == match {
			fmt.Printf("%d + %d = 2020\n%d * %d = %d\n", target, match, target, match, target*match)
			return
		}
	}

	fmt.Println("No two entries sum to 2020")
}

func main() {
	// naive()
	optimized()
}
