package main

import (
	"fmt"

	goltime "github.com/abhishekkr/gol/goltime"
)

func main() {
	timestamp01 := goltime.Timestamp{
		Year:  2014,
		Month: 3,
		Day:   8,
		Hour:  9,
		Min:   44,
		Sec:   55,
	}

	timestamp02 := goltime.CreateTimestamp([]string{
		"2014", "3", "8", "9", "45", "56",
	})

	fmt.Println(timestamp01.Time())
	fmt.Println(timestamp02.Time())
	fmt.Println(goltime.TimestampNow())

	timestamp03 := goltime.CreateTimestamp([]string{
		"2014", "3", "8", "9", "45", "56.0.0.100",
	})
	fmt.Println(timestamp03.Time())

	timestamp04 := goltime.CreateTimestamp([]string{
		"2014", "3", "8", "9", "45", "56.0.0.200",
	})
	fmt.Println(timestamp04)
}
