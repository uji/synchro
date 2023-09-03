// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeVolgogradLocation  sync.Once
	cacheEuropeVolgogradLocation *time.Location
)

type EuropeVolgograd struct{}

func (EuropeVolgograd) Location() *time.Location {
	onceEuropeVolgogradLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Volgograd")
		if err != nil {
			panic(err)
		}
		cacheEuropeVolgogradLocation = loc
	})
	return cacheEuropeVolgogradLocation
}
