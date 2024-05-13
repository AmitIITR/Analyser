package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Event represents a customer activity event.
type Event struct {
	Type       string    `json:"type"`
	CustomerID string    `json:"customer_id"`
	Timestamp  time.Time `json:"timestamp"`
	ItemID     string    `json:"item_id"`
	Price      float64   `json:"price,omitempty"`
}

// UnmarshalJSON customizes the unmarshalling of Event to parse timestamp correctly.
func (e *Event) UnmarshalJSON(data []byte) error {
	var aux struct {
		Type       string  `json:"event_type"`
		CustomerID string  `json:"customer_id"`
		Timestamp  string  `json:"timestamp"`
		ItemID     string  `json:"item_id"`
		Price      float64 `json:"price,omitempty"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Attempt to parse the timestamp
	parsedTime, err := time.Parse("2006-01-02T15:04:05.000000", aux.Timestamp)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02T15:04:05.000000Z", aux.Timestamp)
		if err != nil {
			log.Printf("invalid timestamp: %v", aux.Timestamp)
			return fmt.Errorf("invalid timestamp: %v", aux.Timestamp)
		}
	}

	e.Type = aux.Type
	e.CustomerID = aux.CustomerID
	e.Timestamp = parsedTime
	e.ItemID = aux.ItemID
	e.Price = aux.Price
	return nil
}
