package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var canBeContainedBy = make(map[string][]string)

func listPossibleContainers(contained string) []string {
	if _, found := canBeContainedBy[contained]; !found {
		return []string{contained}
	}
	results := make([]string, 0)
	for _, container := range canBeContainedBy[contained] {
		results = append(results, listPossibleContainers(container)...)
	}
	return append(results, canBeContainedBy[contained]...)
}

func naive() {
	os.Chdir("cmd/daySeven")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	containerRule := regexp.MustCompile("((?:[a-z]+ )+)bags contain")
	containedRule := regexp.MustCompile("([0-9]) ((?:[a-z]+ )+)bags?(?:,|\\.)")
	emptyRule := regexp.MustCompile("no other bags")

	rules := bytes.Split(data, []byte("\n"))
	for _, rule := range rules {
		if len(rule) > 0 {
			if emptyRule.Match(rule) {
				continue
			}
			containerMatch := containerRule.FindAllStringSubmatch(string(rule), -1)
			containerColor := containerMatch[0][1]

			containedMatch := containedRule.FindAllStringSubmatch(string(rule), -1)
			for _, containedEntry := range containedMatch {
				if len(containedEntry) > 2 {
					// containedNum := containedEntry[1]
					containedColor := containedEntry[2]

					if _, found := canBeContainedBy[containedColor]; !found {
						canBeContainedBy[containedColor] = []string{containerColor}
					} else {
						canBeContainedBy[containedColor] = append(canBeContainedBy[containedColor], containerColor)
					}
				}
			}
		}
	}

	possibleContainers := listPossibleContainers("shiny gold ")
	uniquePossibleContainers := []string{}
	for _, container := range possibleContainers {
		alreadyIncluded := false
		for _, uniqueContainer := range uniquePossibleContainers {
			if uniqueContainer == container {
				alreadyIncluded = true
				break
			}
		}
		if !alreadyIncluded {
			uniquePossibleContainers = append(uniquePossibleContainers, container)
		}
	}
	fmt.Println("containers for shiny gold:", uniquePossibleContainers, len(uniquePossibleContainers))
}

func main() {
	naive()
}
