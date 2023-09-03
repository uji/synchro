// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaAnguillaLocation  sync.Once
	cacheAmericaAnguillaLocation *time.Location
)

type AmericaAnguilla struct{}

func (AmericaAnguilla) Location() *time.Location {
	onceAmericaAnguillaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Anguilla")
		if err != nil {
			panic(err)
		}
		cacheAmericaAnguillaLocation = loc
	})
	return cacheAmericaAnguillaLocation
}
