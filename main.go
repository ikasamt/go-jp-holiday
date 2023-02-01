package main

import (
	"embed"
	"flag"
	"fmt"
	"strings"
	"time"
)

//go:embed data/*
var data embed.FS

func isWeekend(date time.Time) bool {
	dayOfWeek := date.Weekday()
	if dayOfWeek == time.Sunday {
		return true
	}

	if dayOfWeek == time.Saturday && !saturday {
		return true
	}

	return false
}

func isHoliday(date time.Time) bool {
	files, err := data.ReadDir("data")
	if err != nil {
		return false
	}

	for _, file := range files {
		if file.Name() == fmt.Sprintf("%d.txt", date.Year()) {
			content, err := data.ReadFile(fmt.Sprintf("data/%s", file.Name()))
			if err != nil {
				return false
			}
			for _, row := range strings.Split(string(content), "\n") {
				cols := strings.Fields(row)
				secondCol := cols[1]
				if secondCol == date.Format("1月2日") {
					return true
				}
			}
		}
	}
	return false
}

func gen(length int) string {
	date := time.Now()
	endDate := date.AddDate(0, 0, length)
	lines := []string{}
	for date.Before(endDate) {
		dayOfWeek := []string{"日", "月", "火", "水", "木", "金", "土"}[date.Weekday()]
		if !isHoliday(date) && !isWeekend(date) {
			lines = append(lines, fmt.Sprintf("%s (%s)", date.Format("01月02日"), dayOfWeek))
		} else {
			lines = append(lines, "") // 祝日は空行
		}
		date = date.AddDate(0, 0, 1)
	}
	return string(strings.Join(lines, "\n"))
}

var size int
var saturday bool

func main() {
	flag.IntVar(&size, "l", 30, "length of password")
	flag.BoolVar(&saturday, "s", false, "include saturday")
	flag.Parse()

	fmt.Println(gen(size))
}
