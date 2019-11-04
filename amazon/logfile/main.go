package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
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
	//fmt.Println(lines)
	dedup := make(map[string]LogEntry)

	for _, line := range lines {
		s := strings.Split(line, "=")
		//fmt.Println(s)
		user := s[0]
		value, _ := strconv.Atoi(s[1])

		logEntry := LogEntry{
			User:  user,
			Value: value,
			Raw:   line,
		}
		dedup[line] = logEntry
	}

	var log []LogEntry

	for _, v := range dedup {
		log = append(log, v)
	}

	sort.Slice(log, func(i, j int) bool {
		return log[i].Value < log[j].Value
	})

	for _, entry := range log {
		fmt.Println(entry.Raw)
	}
}
