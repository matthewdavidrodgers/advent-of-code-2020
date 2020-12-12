package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
)

type bits int

func readAsBits(chars []byte, positiveBit byte) int {
	b := 0
	for i, char := range chars {
		if char == positiveBit {
			b += int(math.Pow(2, float64(len(chars)-i-1)))
		}
	}
	return b
}

func bitwise() {
	os.Chdir("cmd/dayFive")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(data, []byte("\n"))
	sortedWords := make([]int, 0, len(lines))

	var max int
	for _, line := range lines {
		if len(line) > 0 {
			row := line[:7]
			seat := line[7:]

			rowNum := readAsBits(row, byte('B'))
			seatNum := readAsBits(seat, byte('R'))
			seatID := (rowNum * 8) + seatNum

			word := (rowNum << 3) + seatNum
			insertIndex := sort.Search(len(sortedWords), func(i int) bool { return sortedWords[i] > word })

			sortedWords = append(sortedWords, 0)
			copy(sortedWords[insertIndex+1:], sortedWords[insertIndex:])
			sortedWords[insertIndex] = word

			if max < seatID {
				max = seatID
			}
		}
	}

	for i, word := range sortedWords {
		expectedNext := word + 1
		if i != 0 && i != len(sortedWords)-1 && sortedWords[i+1] != expectedNext {
			rowNum := expectedNext >> 3
			seatNum := expectedNext & 7
			fmt.Println("missing seat id:", (rowNum*8)+seatNum)
		}
	}

	fmt.Println("max seat id:", max)
}

func main() {
	bitwise()
}
