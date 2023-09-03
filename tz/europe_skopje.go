// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeSkopjeLocation  sync.Once
	cacheEuropeSkopjeLocation *time.Location
)

type EuropeSkopje struct{}

func (EuropeSkopje) Location() *time.Location {
	onceEuropeSkopjeLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Skopje")
		if err != nil {
			panic(err)
		}
		cacheEuropeSkopjeLocation = loc
	})
	return cacheEuropeSkopjeLocation
}
