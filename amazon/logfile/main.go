package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type LogEntry struct {
	User  string
	Value int
	Raw   string
}

func main() {
	f, _ := os.Open("logfile.txt")
	data, _ := ioutil.ReadAll(f)

	lines := strings.Split(string(data), "\n")

	dedup := make(map[string][]string)
	for _, line := range lines {

		dedup[line] = strings.Split(line, "=")
	}

	var log []LogEntry

	for k, v := range dedup {
		logEntry := LogEntry{
			User:  v[0],
			Value: v[1],
			Raw:   k,
		}
		log = append(log, logEntry)
	}

	fmt.Println(log)
}
