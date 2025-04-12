package time

import (
	"fmt"
	"time"
)

func MoscowLocation() *time.Location {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Printf("could not load location: %v", err)
		return time.Local
	}

	return location
}
