package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	// giving your own time
	// use --> time.Date(year, month, date, hour, minute, second, nanosecond, timeZone or location)
	createdDate := time.Date(2020, time.August, 10, 23, 23, 33, 30000, time.UTC)
	fmt.Println(createdDate)
}
