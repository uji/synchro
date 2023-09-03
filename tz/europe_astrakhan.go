// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeAstrakhanLocation  sync.Once
	cacheEuropeAstrakhanLocation *time.Location
)

type EuropeAstrakhan struct{}

func (EuropeAstrakhan) Location() *time.Location {
	onceEuropeAstrakhanLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Astrakhan")
		if err != nil {
			panic(err)
		}
		cacheEuropeAstrakhanLocation = loc
	})
	return cacheEuropeAstrakhanLocation
}
