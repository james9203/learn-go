package main

type IpLocationCmd struct {
	Host   string `json:"host"`
	Format string `json:"fmt"`
}

type location struct {
	Host string `json:"host"`
	Country string
}
