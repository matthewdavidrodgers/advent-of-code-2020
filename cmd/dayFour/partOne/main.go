package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func naive(validate bool) {
	os.Chdir("cmd/dayFour")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validations := make(map[string]func(val string) bool)
	validations["byr"] = func(val string) bool {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num >= 1920 && num <= 2002
	}
	validations["iyr"] = func(val string) bool {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num >= 2010 && num <= 2020
	}
	validations["eyr"] = func(val string) bool {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num >= 2020 && num <= 2030
	}
	validations["hgt"] = func(val string) bool {
		cmRegex := regexp.MustCompile("^([0-9]+)cm$")
		inRegex := regexp.MustCompile("^([0-9]+)in$")

		if cmRegex.Match([]byte(val)) {
			subMatch := cmRegex.FindStringSubmatch(val)[1]
			num, err := strconv.Atoi(subMatch)
			if err != nil {
				return false
			}
			return num >= 150 && num <= 193
		} else if inRegex.Match([]byte(val)) {
			subMatch := inRegex.FindStringSubmatch(val)[1]
			num, err := strconv.Atoi(subMatch)
			if err != nil {
				return false
			}
			return num >= 59 && num <= 76
		}
		return false
	}
	validations["hcl"] = func(val string) bool {
		regex := regexp.MustCompile("^#((?:[a-f]|[0-9]){6})$")
		return regex.Match([]byte(val))
	}
	validations["ecl"] = func(val string) bool {
		validOpts := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, opt := range validOpts {
			if opt == val {
				return true
			}
		}
		return false
	}
	validations["pid"] = func(val string) bool {
		regex := regexp.MustCompile("^[0-9]{9}$")
		return regex.Match([]byte(val))
	}
	reqFields := make([]string, 0, len(validations))
	for k := range validations {
		reqFields = append(reqFields, k)
	}

	byteEntries := bytes.Split(data, []byte("\n\n"))
	splitRegex := regexp.MustCompile(" |\n")

	valid := 0
	for _, byteEntry := range byteEntries {
		if len(byteEntry) > 0 {
			fields := splitRegex.Split(string(byteEntry), -1)

			remainingReqFields := make([]string, len(reqFields))
			copy(remainingReqFields, reqFields)

			for _, field := range fields {
				if len(remainingReqFields) == 0 {
					break
				}
				fieldParts := strings.Split(field, ":")
				if len(fieldParts) == 2 && len(fieldParts[1]) > 0 {
					currField := fieldParts[0]
					currVal := fieldParts[1]
					for i, reqField := range remainingReqFields {
						if reqField == currField && (!validate || validations[currField](currVal)) {
							copy(remainingReqFields[i:], remainingReqFields[i+1:])
							remainingReqFields = remainingReqFields[:len(remainingReqFields)-1]
							break
						}
					}
				}
			}
			if len(remainingReqFields) == 0 {
				valid++
			}
		}
	}

	fmt.Println(valid, "entries are valid")
}

func main() {
	naive(true)
}
