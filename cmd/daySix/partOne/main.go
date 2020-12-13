package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type set struct {
	mp map[string]bool
}

func newSet() *set {
	return &set{mp: make(map[string]bool)}
}

func (s *set) size() int {
	return len(s.mp)
}

func (s *set) contains(item string) bool {
	_, found := s.mp[item]
	return found
}

func (s *set) insert(item string) bool {
	if s.contains(item) {
		return false
	}
	s.mp[item] = true
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
			for _, member := range members {
				if len(member) > 0 {
					for _, charByte := range member {
						groupAnswers.insert(string(charByte))
					}
				}
			}
			total += groupAnswers.size()
		}
	}

	fmt.Println("total counts", total)
}

func main() {
	naive()
}
