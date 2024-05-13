package main

import "log"

// Bucket represents a collection of events within a specific time frame.
type Bucket struct {
	Name   string
	Events []Event
}

// NewBucket initializes a new bucket with a given name.
func NewBucket(name string) *Bucket {
	return &Bucket{
		Name:   name,
		Events: []Event{},
	}
}

// AddEvent adds an event to the bucket.
func (b *Bucket) AddEvent(event Event) {
	log.Printf("Adding event %v to bucket %v", event, b.Name)
	b.Events = append(b.Events, event)
}
