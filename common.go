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

// GeneralDescription for many entities used in reports
type GeneralDescription struct {
	Text  string `xml:",chardata"`
	Count string `xml:"count,omitempty"`
}

// UserTags Description
type UserTags GeneralDescription

// Host Description
type Host GeneralDescription

// Errors Description
type Errors GeneralDescription

// ResultCount struct used in report struct
type ResultCount struct {
	Text     string           `xml:",chardata"`
	Full     string           `xml:"full,omitempty"`
	Filtered string           `xml:"filtered,omitempty"`
	Debug    *ResultCountItem `xml:"debug,omitempty"`
	Hole     *ResultCountItem `xml:"hole,omitempty"`
	Info     *ResultCountItem `xml:"info,omitempty"`
	Log      *ResultCountItem `xml:"log,omitempty"`
	Warning  *ResultCountItem `xml:"warning,omitempty"`
}

// ResultCountItem used in ResultCount struct
type ResultCountItem struct {
	Text     string `xml:",chardata"`
	Full     string `xml:"full,omitempty"`
	Filtered string `xml:"filtered,omitempty"`
}

// Filters for many struct
type Filters struct {
	Text     string   `xml:",chardata"`
	ID       string   `xml:"id,attr,omitempty"`
	Term     string   `xml:"term,omitempty"`
	Filter   []string `xml:"filter,omitempty"`
	Keywords *struct {
		Text    string `xml:",chardata"`
		Keyword []*struct {
			Text     string `xml:",chardata"`
			Column   string `xml:"column,omitempty"`
			Relation string `xml:"relation,omitempty"`
			Value    string `xml:"value,omitempty"`
		} `xml:"keyword,omitempty"`
	} `xml:"keywords,omitempty"`
}
