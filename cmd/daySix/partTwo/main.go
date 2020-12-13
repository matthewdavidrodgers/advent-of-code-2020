package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type set struct {
	mp map[string]int
}

func newSet() *set {
	return &set{mp: make(map[string]int)}
}

func (s *set) sizeFor(target int) int {
	size := 0
	for _, v := range s.mp {
		if v == target {
			size++
		}
	}
	return size
}

func (s *set) contains(item string) bool {
	_, found := s.mp[item]
	return found
}

func (s *set) insert(item string) bool {
	if s.contains(item) {
		s.mp[item] = s.mp[item] + 1
		return false
	}
	s.mp[item] = 1
	return true
}

func naive() {
	os.Chdir("cmd/daySix")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	groups := bytes.Split(data, []byte("\n\n"))

	for _, group := range groups {
		if len(group) > 0 {
			groupAnswers := newSet()
			members := bytes.Split(group, []byte("\n"))
			groupLen := 0
			for _, member := range members {
				if len(member) > 0 {
					groupLen++
					for _, charByte := range member {
						groupAnswers.insert(string(charByte))
					}
				}
			}
			total += groupAnswers.sizeFor(groupLen)
		}
	}

	fmt.Println("total counts", total)
}

func main() {
	naive()
}
