// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeGuernseyLocation  sync.Once
	cacheEuropeGuernseyLocation *time.Location
)

type EuropeGuernsey struct{}

func (EuropeGuernsey) Location() *time.Location {
	onceEuropeGuernseyLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Guernsey")
		if err != nil {
			panic(err)
		}
		cacheEuropeGuernseyLocation = loc
	})
	return cacheEuropeGuernseyLocation
}
