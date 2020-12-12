package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type slope struct {
	slopeX      int
	slopeY      int
	encountered int
}

func naive() {
	os.Chdir("cmd/dayThree")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteRows := bytes.Split(data, []byte("\n"))

	trackers := []*slope{
		{slopeX: 1, slopeY: 1},
		{slopeX: 3, slopeY: 1},
		{slopeX: 5, slopeY: 1},
		{slopeX: 7, slopeY: 1},
		{slopeX: 1, slopeY: 2},
	}
	mapY := 0
	for _, byteRow := range byteRows {
		if len(byteRow) > 0 {
			for _, tracker := range trackers {
				if (mapY % tracker.slopeY) == 0 {
					x := ((mapY / tracker.slopeY) * tracker.slopeX) % len(byteRow)
					if string(byteRow[x]) == "#" {
						tracker.encountered++
					}
				}
			}
			mapY++
		}
	}

	sum := 1
	for _, tracker := range trackers {
		sum *= tracker.encountered
		fmt.Printf("for slope x=%d and y=%d : encountered %d\n", tracker.slopeX, tracker.slopeY, tracker.encountered)
	}
	fmt.Println("Sum of encounters : ", sum)

}

func main() {
	naive()
}
