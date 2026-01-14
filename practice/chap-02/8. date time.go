package main

import (
	"fmt"
	"time"
)

func calculateDuration() {
	start := time.Now()

	// Mimic the processing...
	// Heavy work
	time.Sleep(10 * time.Second)

	// i := 0
	// for {
	// 	i += 1
	// 	if i == 10000000000 {
	// 		break
	// 	}
	// }

	duration := time.Since(start)
	fmt.Println("duration: ", duration)
}
func main() {
	// t := time.Now()
	// fmt.Println("time:", t)

	// unixTime := time.Now().Unix()
	// fmt.Println("Unix time:", unixTime)

	// normalTime := time.Unix(int64(unixTime), 0)
	// fmt.Println("Normal time:", normalTime)

	// // Year, month, day
	// fmt.Println("Year:", t.Year())
	// fmt.Println("Month:", t.Month())
	// fmt.Println("Day:", t.Day())

	// // Timezone/location
	// fmt.Println("Location:", t.Location())

	fmt.Println("-----------------------")
	dateStr := "2026-01-14 19:44:42"
	d, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		fmt.Println("Parse err:", err)
	}

	fmt.Println("dateString:", dateStr)
	fmt.Println("date:", d)
	fmt.Println("Year:", d.Year())
	fmt.Println("Month:", d.Month())
	fmt.Println("Day:", d.Day())
	fmt.Println("Hour:", d.Hour())
	fmt.Println("Minute:", d.Minute())
	fmt.Println("Second:", d.Second())

	dateFromate := d.Format("2006-01-02")
	fmt.Println("Format date:", dateFromate)
	fmt.Printf("Format type: %T\n", dateFromate)

	// Timezones
	dt := time.Now()
	local := dt.Location()
	fmt.Println("local:", local)

	timezone := "Asia/Kolkata"
	fmt.Println("timezone:", timezone)

	// timezone := "America/New_York"
	// fmt.Println("timezone:", timezone)

	// timezone := "Local"
	// fmt.Println("timezone:", timezone)

	loc, _ := time.LoadLocation(timezone) // If loc is invalid it return: UTC
	fmt.Println("loc:", loc)

	fmt.Println("--------------------")
	calculateDuration()
}
