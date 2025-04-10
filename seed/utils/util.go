package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strconv"
	"strings"
	"time"
)

type TimeRes struct {
	Day         string
	OpeningTime time.Time
	ClosingTime time.Time
}

func parseTime(timeStr, amOrPm string) (bool, int, int) {
	timeParts := strings.Split(timeStr, ":")
	var hour, minute int
	var err error

	if len(timeParts) == 1 { // Only hour is provided
		hour, err = strconv.Atoi(timeParts[0])
		minute = 0
	} else if len(timeParts) == 2 { // Hour and minute provided
		hour, err = strconv.Atoi(timeParts[0])
		if err == nil {
			minute, err = strconv.Atoi(timeParts[1])
		}
	} else {
		return false, 0, 0
	}

	if err != nil || hour > 12 || minute < 0 || minute > 59 {
		return false, 0, 0
	}
	if amOrPm == "pm" && hour < 12 {
		hour += 12
	} else if amOrPm == "am" && hour >= 12 {
		hour -= 12
	}
	return true, hour, minute
}

func parseDay(day string) (string, error) {
	var dayMap = map[string]string{
		"fri":   "friday",
		"frid":  "friday",
		"fr":    "friday",
		"f":     "friday",
		"sa":    "saturday",
		"sat":   "saturday",
		"satur": "saturday",
		"sun":   "sunday",
		"su":    "sunday",
		"mon":   "monday",
		"mo":    "monday",
		"m":     "monday",
		"tu":    "tuesday",
		"tues":  "tuesday",
		"tue":   "tuesday",
		"wed":   "wednesday",
		"weds":  "wednesday",
		"we":    "wednesday",
		"w":     "wednesday",
		"th":    "thursday",
		"thu":   "thursday",
		"thur":  "thursday",
		"thurs": "thursday",
	}
	if fullDay, exists := dayMap[day]; exists {
		return fullDay, nil
	}
	return "", errors.New("invalid day format")
}

func CheckValidity(day string, closingTime string, openingTime string, openingTimeAmPM string, closingTimeAmPM string) (bool, string, int, int, int, int) {
	if !(openingTimeAmPM == "am" || openingTimeAmPM == "pm") {
		return false, "", 0, 0, 0, 0
	}
	if !(closingTimeAmPM == "am" || closingTimeAmPM == "pm") {
		return false, "", 0, 0, 0, 0
	}

	day, err := parseDay(day)
	if err != nil {
		return false, "", 0, 0, 0, 0
	}

	flag, hour, min := parseTime(openingTime, openingTimeAmPM)
	if !flag {
		return false, "", 0, 0, 0, 0
	}

	flag, hour1, min1 := parseTime(closingTime, closingTimeAmPM)
	if !flag {
		return false, "", 0, 0, 0, 0
	}

	// opening hour is smaller than closing hour
	if hour*60+min >= hour1*60+min1 && closingTimeAmPM == "am" {
		hour1 = 23 // Todo take accurate data
		min1 = 59
	} else if hour*60+min >= hour1*60+min1 {
		return false, "", 0, 0, 0, 0
	}

	return true, day, hour, min, hour1, min1
}

func CreateIndex(esClient *elasticsearch.Client, indexName string) {
	mapping := `{
		"mappings": {
			"properties": {
				"name": { "type": "text" },
				"type": { "type": "keyword" }
			}
		}
	}`
	res, err := esClient.Indices.Create(indexName, esClient.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		log.Fatalf("Error creating index: %v", err)
	}
	defer res.Body.Close()

	fmt.Println("Index created successfully")
}

func InsertDataEs(esClient *elasticsearch.Client, indexName string, data []map[string]interface{}) {
	var buf bytes.Buffer
	for _, item := range data {
		meta := []byte(`{ "index": {} }` + "\n")
		jsonData, err := json.Marshal(item)
		if err != nil {
			log.Fatalf("Error marshalling data: %v", err)
		}
		jsonData = append(jsonData, '\n')

		buf.Write(meta)
		buf.Write(jsonData)
	}

	res, err := esClient.Bulk(bytes.NewReader(buf.Bytes()), esClient.Bulk.WithIndex(indexName))
	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}
	defer res.Body.Close()

	fmt.Println("Data inserted successfully")
}
