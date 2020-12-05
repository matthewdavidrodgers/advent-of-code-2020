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
			posOne, err := strconv.Atoi(string(reqRangeBytes[0]))
			if err != nil {
				log.Fatal(err)
			}
			posTwo, err := strconv.Atoi(string(reqRangeBytes[1]))
			if err != nil {
				log.Fatal(err)
			}

			posOneMatches := posOne <= len(password) && password[posOne-1] == reqCharByte
			posTwoMatches := posTwo <= len(password) && password[posTwo-1] == reqCharByte

			if posOneMatches != posTwoMatches {
				validCount++
			} else {
				invalidCount++
			}
		}
	}

	fmt.Printf("%d passwords are valid out of %d\n", validCount, validCount+invalidCount)
}

func main() {
	naive()
}
