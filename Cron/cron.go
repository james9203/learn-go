package main

import (
"github.com/robfig/cron"
"log"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0/1 * * * * * "
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	//c.AddFunc("@every 1h1m", func() {
	//	i++
	//	log.Println("cron running:", i)
	//})
	c.Start()
	select {

	}
}