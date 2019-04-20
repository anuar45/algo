package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const photosStr string = `photo1.jpg, Warsaw, 2013-05-03 14:20:54
photo2.png, London, 2013-05-04 14:20:00
photo3.jpg, Dublin, 2013-07-03 18:50:00
photo4.png, Dublin, 2013-04-03 12:35:00
photo5.jpg, Warsaw, 2013-05-03 13:10:00
photo6.png, Dublin, 2013-02-23 15:40:00
photo7.jpg, Dublin, 2013-09-16 15:40:00
photo8.jpg, Dublin, 2013-02-25 15:40:00
photo9.png, Warsaw, 2013-06-04 15:40:00
photo10.jpg, Dublin, 2013-09-01 15:40:00
photo11.png, Dublin, 2013-02-09 15:40:00
photo12.jpg, Dublin, 2013-09-01 15:40:00
photo13.png, London, 2013-09-02 15:40:00
photo14.jpg, Dublin, 2013-08-07 15:40:00
photo15.png, Dublin, 2013-09-03 15:40:00
photo16.png, Dublin, 2013-09-05 15:40:00`

type Photo struct {
	Name  string
	City  string
	Date  time.Time
	Order int
}

func main() {
	log.SetOutput(os.Stdout)
	fmt.Println(photosStr)
	photos := parsePhotos(photosStr)
	fmt.Println(photos)

	sortPhotosByDate(photos)
	fmt.Println(photos)

	cityCount := countPhotosInCity(photos)
	fmt.Println(cityCount)

	renamePhotos(photos, cityCount)
	fmt.Println(photos)

}

func renamePhotos(photos []Photo, cityCount map[string]int) {
	photoNum := 1
	var zerosPrefix string
	for city, count := range cityCount {
		fmt.Println(city, count)
		zerosPrefix = getZerosStr(count)
		fmt.Println("zerosPrefix", zerosPrefix)
		for i := 0; i < len(photos); i++ {
			if photos[i].City == city {
				fmt.Println("photo:", photos[i])
				photoNumStr := strconv.Itoa(photoNum)
				extName := photos[i].Name[len(photos[i].Name)-4:]
				photos[i].Name = photos[i].City + zerosPrefix + photoNumStr + extName
			}
			photoNum++
		}
	}
}

func getZerosStr(num int) string {
	numStr := strconv.Itoa(num)

	var s string
	for i := 0; i < len(numStr)-1; i++ {
		s += "0"
	}
	return s
}

func countPhotosInCity(photos []Photo) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(photos); i++ {
		m[photos[i].City]++
	}
	return m
}

func sortPhotosByDate(photos []Photo) {
	for i := 0; i < len(photos); i++ {
		for j, k := 0, 1; k < len(photos); {
			if photos[k].Date.Before(photos[j].Date) {
				photos[j], photos[k] = photos[k], photos[j]
			}
			j++
			k++
		}
	}
}

func parsePhotos(s string) []Photo {
	var photos []Photo
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		fields := strings.Split(line, ", ")

		var photo Photo
		photo.Name = fields[0]
		photo.City = fields[1]
		date, _ := time.Parse("2006-01-02 15:04:05", fields[2])
		photo.Date = date
		photo.Order = i

		photos = append(photos, photo)
	}
	return photos
}
