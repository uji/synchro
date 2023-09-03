// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAntarcticaMacquarieLocation  sync.Once
	cacheAntarcticaMacquarieLocation *time.Location
)

type AntarcticaMacquarie struct{}

func (AntarcticaMacquarie) Location() *time.Location {
	onceAntarcticaMacquarieLocation.Do(func() {
		loc, err := time.LoadLocation("Antarctica/Macquarie")
		if err != nil {
			panic(err)
		}
		cacheAntarcticaMacquarieLocation = loc
	})
	return cacheAntarcticaMacquarieLocation
}
