package main

import (
	"fmt"
	"strings"
)

func findErrors(logs []string) []string {
	var errors []string

	for _, log := range logs {
		if strings.Contains(log, "ERROR") {
			errors = append(errors, log)
		}
	}

	return errors
}

func main() {
	logs := []string{
		"INFO Server started",
		"ERROR Database connection failed",
	}

	fmt.Println(findErrors(logs))
}
