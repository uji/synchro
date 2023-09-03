// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeWarsawLocation  sync.Once
	cacheEuropeWarsawLocation *time.Location
)

type EuropeWarsaw struct{}

func (EuropeWarsaw) Location() *time.Location {
	onceEuropeWarsawLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Warsaw")
		if err != nil {
			panic(err)
		}
		cacheEuropeWarsawLocation = loc
	})
	return cacheEuropeWarsawLocation
}
