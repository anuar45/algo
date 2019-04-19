package main

import (
	"fmt"
	"strings"
	"time"
)

const photosStr string = "photo1, Warsaw, 2013-05-03 14:20:54\nphoto2, London, 2013-05-04 14:20:00\nphoto3, Dublin, 2013-07-03 18:50:00\nphoto4, Dublin, 2013-04-03 12:35:00\nphoto5, Warsaw, 2013-05-03 13:10:00\nphoto6, Dublin, 2013-09-01 15:40:00"

type Photo struct {
	Name  string
	City  string
	Date  time.Time
	Order int
}

func main() {
	photos := parsePhotos(photosStr)
}

func parsePhotos(s string) []Photo {
	var photos []Photo

	lines := strings.Split("\n")

	for i, line := range lines {
		fields := strings.Split(line, ", ")

		var photo Photo
		photo.Name = fields[0]
		photo.City = fields[1]
		photo.Date = time.Parse("2006-01-02 15:04:05", fields[2])

		photos = append(photos, photo)
	}
}
