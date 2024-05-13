package main

import (
	"fmt"
	"log"
	"time"
)

// Metrics stores time bucketed metrics for different time frames.
type Metrics struct {
	TimeBuckets map[time.Time]*Bucket
}

// NewMetrics creates a new Metrics struct.
func NewMetrics() *Metrics {
	return &Metrics{
		TimeBuckets: make(map[time.Time]*Bucket),
	}
}

// ProcessEvent adds an event to the appropriate bucket based on the given bucket duration.
func (m *Metrics) ProcessEvent(event Event, bucketDuration time.Duration) {
	bucketTime := event.Timestamp.Truncate(bucketDuration)
	if _, exists := m.TimeBuckets[bucketTime]; !exists {
		m.TimeBuckets[bucketTime] = NewBucket(bucketTime.String())
		log.Printf("Creating new bucket %v", bucketTime.String())
	}
	m.TimeBuckets[bucketTime].AddEvent(event)
}

// GenerateInsights generates insights from the metrics.
func (m *Metrics) GenerateInsights() {
	log.Println("Generating insights from the collected metrics...")
	for bucketTime, bucket := range m.TimeBuckets {
		fmt.Printf("Time Bucket: %s\n", bucketTime)

		views := make(map[string]int)
		addToCarts := make(map[string]int)
		purchases := make(map[string]float64)

		for _, event := range bucket.Events {
			switch event.Type {
			case "page_visit":
				views[event.ItemID]++
			case "add_to_cart":
				addToCarts[event.ItemID]++
			case "purchase":
				purchases[event.ItemID] += event.Price
			}
		}
		// Most viewed item
		var maxViews int
		var mostViewedItem string
		for item, count := range views {
			if count > maxViews {
				maxViews = count
				mostViewedItem = item
			}
		}

		// Highest-grossing item
		var maxRevenue float64
		var highestGrossingItem string
		for item, revenue := range purchases {
			if revenue > maxRevenue {
				maxRevenue = revenue
				highestGrossingItem = item
			}
		}

		// Conversion rates
		fmt.Println("Conversion rates:")
		for item, views := range views {
			addToCart := addToCarts[item]
			var rate float64
			if views > 0 {
				rate = float64(addToCart) / float64(views)
			}
			fmt.Printf("Item %s: %.2f%%\n", item, rate*100)
		}

		fmt.Printf("Most viewed item: %s with %d views\n", mostViewedItem, maxViews)
		fmt.Printf("Highest-grossing item: %s with %.2f revenue\n", highestGrossingItem, maxRevenue)
		fmt.Println()
	}
}
