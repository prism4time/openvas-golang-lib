package omp

import (
	"encoding/xml"
)

// OMP Response Status .
const (
	StatusOK int = 200
)

// request / response interface .
type requestIface interface {
}

type responseIface interface {
	GetStatus() int
	GetStatusText() string
}

// request / response base .
type requestBase struct {
}

type responseBase struct {
	Status     int    `xml:"status,attr,omitempty"`
	StatusText string `xml:"status_text,attr,omitempty"`
}

func (r *responseBase) GetStatus() int {
	return r.Status
}

func (r *responseBase) GetStatusText() string {
	return r.StatusText
}

// Auth .
type requestAuth struct {
	requestBase
	XMLName     xml.Name `xml:"authenticate"`
	Credentials struct {
		Username string `xml:"username,omitempty"`
		Password string `xml:"password,omitempty"`
	} `xml:"credentials,omitempty"`
}

type responseAuth struct {
	responseBase
}

// Version .
type requestVersion struct {
	requestBase
	XMLName xml.Name `xml:"get_version"`
}

type responseVersion struct {
	responseBase
	Version string `xml:"version,omitempty"`
}
