package main

import (
	"os"
	"log"
	"fmt"
	"encoding/csv"
	"strings"
	"io"
)

func csv_reader() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}

func csv_readerall() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`
	r := csv.NewReader(strings.NewReader(in))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}

func csv_writer() {
	records := [][] string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func csv_writerall() {
	records := [][] string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func main() {
	fmt.Println("---------------------------------------")
	csv_writer()
	fmt.Println("---------------------------------------")
	csv_writerall()
	fmt.Println("---------------------------------------")
	csv_reader()
	fmt.Println("---------------------------------------")
	csv_readerall()
	fmt.Println("---------------------------------------")
}
