// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaArgentinaTucumanLocation  sync.Once
	cacheAmericaArgentinaTucumanLocation *time.Location
)

type AmericaArgentinaTucuman struct{}

func (AmericaArgentinaTucuman) Location() *time.Location {
	onceAmericaArgentinaTucumanLocation.Do(func() {
		loc, err := time.LoadLocation("America/Argentina/Tucuman")
		if err != nil {
			panic(err)
		}
		cacheAmericaArgentinaTucumanLocation = loc
	})
	return cacheAmericaArgentinaTucumanLocation
}
