package model

import "time"

type DnsData struct {
	Class   string
	Type    string
	Payload string
	Tag     bool
	Time    time.Time
}

type LabeledDomain struct {
	Payload string `json:"payload"`
	Tag     bool   `json:"tag"`
	Time    string `json:"time"`
}
