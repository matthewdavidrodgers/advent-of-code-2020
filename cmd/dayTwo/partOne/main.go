package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func naive() {
	os.Chdir("./cmd/dayTwo")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteRows := bytes.Split(data, []byte("\n"))

	validCount := 0
	invalidCount := 0
	for _, byteRow := range byteRows {
		if len(byteRow) != 0 {
			bySpace := bytes.Split(byteRow, []byte(" "))
			reqRangeBytes, reqCharByte, password := bytes.Split(bySpace[0], []byte("-")), bySpace[1][0], bySpace[2]
			min, err := strconv.Atoi(string(reqRangeBytes[0]))
			if err != nil {
				log.Fatal(err)
			}
			max, err := strconv.Atoi(string(reqRangeBytes[1]))
			if err != nil {
				log.Fatal(err)
			}

			reqCharCount := 0
			for _, char := range password {
				if char == reqCharByte {
					reqCharCount++
				}
			}
			if reqCharCount >= min && reqCharCount <= max {
				validCount++
			} else {
				invalidCount++
			}

		}
	}

	fmt.Printf("%d passwords are valid out of %d\n", validCount, invalidCount)
}

func main() {
	naive()
}
