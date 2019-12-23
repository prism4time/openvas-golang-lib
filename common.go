package omp

// Port used in report results
type Port struct {
	Text     string `xml:",chardata"`
	Host     string `xml:"host,omitempty"`
	Severity string `xml:"severity,omitempty"`
	Threat   string `xml:"threat,omitempty"`
}

// Ports wrap port struct
type Ports struct {
	Text  string `xml:",chardata"`
	Start string `xml:"start,attr,omitempty"`
	Max   string `xml:"max,attr,omitempty"`
	Port  []Port `xml:"port,omitempty"`
}

// Cert .
type Cert struct {
	Text    string `xml:",chardata"`
	CertRef *struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr,omitempty"`
		Type string `xml:"type,attr,omitempty"`
	} `xml:"cert_ref,omitempty"`
}

// Nvt .
type Nvt struct {
	Text     string `xml:",chardata"`
	Oid      string `xml:"oid,attr,omitempty"`
	Name     string `xml:"name,omitempty"`
	CvssBase string `xml:"cvss_base,omitempty"`
	Cve      string `xml:"cve,omitempty"`
	Bid      string `xml:"bid,omitempty"`
	Tags     string `xml:"tags,omitempty"`
	Cert     *Cert  `xml:"cert,omitempty"`
	Xref     string `xml:"xref,omitempty"`
}

// GeneralDescription for many entities
type GeneralDescription struct {
	Text  string `xml:",chardata"`
	Count string `xml:"count,omitempty"`
}

// UserTags .
type UserTags GeneralDescription

// Host .
type Host GeneralDescription

// Hosts .
type Hosts GeneralDescription

// Errors .
type Errors GeneralDescription
