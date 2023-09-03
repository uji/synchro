// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaAsuncionLocation  sync.Once
	cacheAmericaAsuncionLocation *time.Location
)

type AmericaAsuncion struct{}

func (AmericaAsuncion) Location() *time.Location {
	onceAmericaAsuncionLocation.Do(func() {
		loc, err := time.LoadLocation("America/Asuncion")
		if err != nil {
			panic(err)
		}
		cacheAmericaAsuncionLocation = loc
	})
	return cacheAmericaAsuncionLocation
}
