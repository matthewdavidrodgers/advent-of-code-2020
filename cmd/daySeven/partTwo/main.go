package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

type bagReq struct {
	color string
	num   int
}

var containingReqs = make(map[string][]bagReq)

func countContainedBags(container string) int {
	if _, found := containingReqs[container]; !found {
		return 0
	}
	count := 0
	for _, rule := range containingReqs[container] {
		count += rule.num + (rule.num * countContainedBags(rule.color))
	}
	return count
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
			containerRules := make([]bagReq, 0, len(containedMatch))
			for _, containedEntry := range containedMatch {
				if len(containedEntry) > 2 {
					containedNum, _ := strconv.Atoi(containedEntry[1])
					containedColor := containedEntry[2]
					containerRules = append(containerRules, bagReq{
						color: containedColor,
						num:   containedNum,
					})
				}
			}
			containingReqs[containerColor] = containerRules
		}
	}

	fmt.Println("Count of contained bags for shiny gold", countContainedBags("shiny gold "))
}

func main() {
	naive()
}
