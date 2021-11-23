package db

import (
	"time"
)

type Image struct {
	url    string
	width  int
	height int
}

type Ad struct {
	id          string
	url         string
	path        string
	referer     string
	host        string
	href        string
	sourceId    int
	on          string
	destination string
	nonce       string
	pcmVersion  int
	image       Image
	time        time.Time
}

type PublickToken struct {
	id      string
	url     string
	path    string
	referer string
	host    string
	token   string
	time    time.Time
}

type SourceToken struct {
	id             string
	url            string
	path           string
	referer        string
	host           string
	engagementType string
	nonce          string
	token          string
	pcmVersion     int
	time           time.Time
}

type Trigger struct {
	id       string
	url      string
	path     string
	referer  string
	host     string
	cv       int
	priority int
	time     time.Time
}

type Report struct {
	id                   string
	url                  string
	path                 string
	referer              string
	host                 string
	engagementType       string
	sourceSite           int
	sourceId             string
	triggerData          int
	pcmVersion           int
	secretToken          string
	secretTokenSignature string
	time                 time.Time
}
