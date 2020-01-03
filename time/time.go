package main

import (
	"fmt"
	"time"
)

func month() {
	_, month, day := time.Now().Date()
	if month == time.January && day == 2 {
		fmt.Println("Happy Go day!")
	}
}

func newticker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time:", t)
		}
	}

}

func parse() {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2020 at 7:54pm (PST)")
	fmt.Println(t)
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2020-Feb-03")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2020-01-02T16:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err)

}

func parseDuration() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	micro2, _ := time.ParseDuration("1us")
	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro2)
}

func parseInLocation() {
	loc, _ := time.LoadLocation("Europe/Berlin")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2020 at 5:02am (CEST)", loc)
	fmt.Println(t)
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2020-Jul-09", loc)
	fmt.Println(t)
}

func sleep() {
	time.Sleep(100 * time.Millisecond)
}

func statusUpdate() string {
	return ""
}

func tick() {
	c := time.Tick(5 * time.Second)
	var a int = 0
	for now := range c {
		a++
		if (a == 5) {
			break
		}
		fmt.Printf("%v %s\n", now, statusUpdate())
	}
}

func timeAdd() {
	start := time.Date(2020, 01, 01, 12, 12, 12, 12, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)
	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second *10)=%v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute *10)=%v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour *10)=%v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour*24*10)=%v\n", afterTenDays)
}

func timeAddDate() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)
	fmt.Printf("oneDayLater:start.AddDate(0,0,1)=%v\n", oneDayLater)
	fmt.Printf("oneMonthLater:start.AddDate(0,1,0)=%v\n", oneMonthLater)
	fmt.Printf("oneYearLater:start.AddDate(1,0,0)=%v\n", oneYearLater)
}

func timeAfter() {
	year2000 := time.Date(2000, 1, 1, 0, 0,0, 0, time.UTC)
	year3000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	isYear3000AfterYear2000 := year3000.After(year2000)
	isYear2000AfterYear3000 := year2000.After(year3000)
	fmt.Printf("year3000.After(year2000)=%v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000)=%v\n", isYear2000AfterYear3000)
}

func main() {
	fmt.Println("-------------------------")
	month()
	fmt.Println("-------------------------")
	newticker()
	fmt.Println("-------------------------")
	parse()
	fmt.Println("-------------------------")
	parseDuration()
	fmt.Println("-------------------------")
	parseInLocation()
	fmt.Println("-------------------------")
	sleep()
	fmt.Println("-------------------------")
	tick()
	fmt.Println("-------------------------")
	timeAdd()
	fmt.Println("-------------------------")
	timeAddDate()
	fmt.Println("-------------------------")
	timeAfter()
	fmt.Println("-------------------------")
}
