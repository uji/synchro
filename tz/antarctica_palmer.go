// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAntarcticaPalmerLocation  sync.Once
	cacheAntarcticaPalmerLocation *time.Location
)

type AntarcticaPalmer struct{}

func (AntarcticaPalmer) Location() *time.Location {
	onceAntarcticaPalmerLocation.Do(func() {
		loc, err := time.LoadLocation("Antarctica/Palmer")
		if err != nil {
			panic(err)
		}
		cacheAntarcticaPalmerLocation = loc
	})
	return cacheAntarcticaPalmerLocation
}
