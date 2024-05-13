package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("sample_events.jsonl")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	metrics := NewMetrics()
	bucketDuration := time.Hour // Use hourly buckets

	for scanner.Scan() {
		var event Event
		err := json.Unmarshal(scanner.Bytes(), &event)
		if err != nil {
			log.Printf("Error decoding JSON: %v, line: %v", err, scanner.Text())
			continue
		}
		log.Printf("Processing event: %+v", event)
		metrics.ProcessEvent(event, bucketDuration)
	}

	metrics.GenerateInsights()
}
