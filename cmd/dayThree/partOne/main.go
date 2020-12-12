package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func naive() {
	os.Chdir("cmd/dayThree")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteRows := bytes.Split(data, []byte("\n"))

	entries := make([][]bool, 0, len(byteRows))
	for _, byteRow := range byteRows {
		if len(byteRow) > 0 {
			entryRow := make([]bool, 0, len(byteRow))
			for _, byteChar := range byteRow {
				if string(byteChar) == "#" {
					entryRow = append(entryRow, true)
				} else {
					entryRow = append(entryRow, false)
				}
			}
			entries = append(entries, entryRow)
		}
	}

	encountered := 0
	i := 0
	for _, row := range entries {
		if row[i] {
			encountered++
		}
		i += 3
		if i >= len(row) {
			i = i - len(row)
		}
	}

	fmt.Printf("Encountered %d trees\n", encountered)
}

type bits int

func addOne(val bits) bits {
	shifted := val << 1
	shifted = shifted | 1
	return shifted
}

func addZero(val bits) bits {
	shifted := val << 1
	return shifted
}

func createMask(oneIndexes []int) bits {
	var mask bits = 0
	for _, index := range oneIndexes {
		mask = mask >> index
		mask = mask | 1
		mask = mask << index
	}
	return mask
}

func bitwise() {
	os.Chdir("cmd/dayThree")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteRows := bytes.Split(data, []byte("\n"))

	encountered := 0
	rowLen := 0
	i := 0
	for _, byteRow := range byteRows {
		if len(byteRow) > 0 {
			var entryRow bits = 0
			for _, byteChar := range byteRow {
				if string(byteChar) == "#" {
					entryRow = addOne(entryRow)
				} else {
					entryRow = addZero(entryRow)
				}
			}
			if rowLen == 0 {
				rowLen = len(byteRow)
			}
			mask := createMask([]int{rowLen - i - 1})
			if (entryRow & mask) > 0 {
				encountered++
			}
			i += 3
			if i >= rowLen {
				i = i - rowLen
			}
		}
	}

	fmt.Printf("Encountered %d trees\n", encountered)
}

func main() {
	// naive()
	bitwise()
}
