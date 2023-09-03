// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAfricaJohannesburgLocation  sync.Once
	cacheAfricaJohannesburgLocation *time.Location
)

type AfricaJohannesburg struct{}

func (AfricaJohannesburg) Location() *time.Location {
	onceAfricaJohannesburgLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Johannesburg")
		if err != nil {
			panic(err)
		}
		cacheAfricaJohannesburgLocation = loc
	})
	return cacheAfricaJohannesburgLocation
}
