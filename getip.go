package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func uniqueIP(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	content, err := ioutil.ReadFile("log")

	if err != nil {
		log.Fatal(err)
	}

	str1 := string(content)
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	submatchall := re.FindAllString(str1, -1)

	reverseSubmatchall := reverse(submatchall)

	countOfIP := 0
	ipScope := (uniqueIP(reverseSubmatchall))
	for _, element := range ipScope {
		fmt.Println(element)
		countOfIP += 1
		if countOfIP == 100 {
			break
		}
	}
}
