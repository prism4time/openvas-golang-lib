package omp

import "encoding/xml"

// GetReports represent the get_reports request
type GetReports struct {
	requestBase
	XMLName  xml.Name `xml:"get_reports"`
	Text     string   `xml:",chardata"`
	ReportID string   `xml:"report_id,attr,omitempty"`
}

// GetReportsResponse represent the response from get_reports
type GetReportsResponse struct {
	responseBase
	XMLName xml.Name `xml:"get_reports_response"`
	Text    string   `xml:",chardata"`
	Report  *Report  `xml:"report,omitempty"`
}

// Report for task
type Report struct {
	Text        string          `xml:",chardata"`
	ID          string          `xml:"id,attr,omitempty"`
	FormatID    string          `xml:"format_id,attr,omitempty"`
	Extension   string          `xml:"extension,attr,omitempty"`
	ContentType string          `xml:"content_type,attr,omitempty"`
	Report      *DetailedReport `xml:"report,omitempty"`
}

// DetailedReport struct used in report struct
type DetailedReport struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Omp  *struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version,omitempty"`
	} `xml:"omp,omitempty"`
	Sort *struct {
		Text  string `xml:",chardata"`
		Field struct {
			Text  string `xml:",chardata"`
			Order string `xml:"order,omitempty"`
		} `xml:"field,omitempty"`
	} `xml:"sort,omitempty"`
	Filters *struct {
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
	} `xml:"filters,omitempty"`
	UserTags      *UserTags `xml:"user_tags,omitempty"`
	ScanRunStatus string    `xml:"scan_run_status,omitempty"`
	Hosts         *Hosts    `xml:"hosts,omitempty"`
	ClosedCves    *struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,omitempty"`
	} `xml:"closed_cves,omitempty"`
	Vulns *struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,omitempty"`
	} `xml:"vulns,omitempty"`
	Os *struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,omitempty"`
	} `xml:"os,omitempty"`
	Apps *struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,omitempty"`
	} `xml:"apps,omitempty"`
	SslCerts *struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,omitempty"`
	} `xml:"ssl_certs,omitempty"`
	ResultCount *struct {
		Text     string `xml:",chardata"`
		Full     string `xml:"full,omitempty"`
		Filtered string `xml:"filtered,omitempty"`
		Debug    *struct {
			Text     string `xml:",chardata"`
			Full     string `xml:"full,omitempty"`
			Filtered string `xml:"filtered,omitempty"`
		} `xml:"debug,omitempty"`
		Hole *struct {
			Text     string `xml:",chardata"`
			Full     string `xml:"full,omitempty"`
			Filtered string `xml:"filtered,omitempty"`
		} `xml:"hole,omitempty"`
		Info *struct {
			Text     string `xml:",chardata"`
			Full     string `xml:"full,omitempty"`
			Filtered string `xml:"filtered,omitempty"`
		} `xml:"info,omitempty"`
		Log *struct {
			Text     string `xml:",chardata"`
			Full     string `xml:"full,omitempty"`
			Filtered string `xml:"filtered,omitempty"`
		} `xml:"log,omitempty"`
		Warning *struct {
			Text     string `xml:",chardata"`
			Full     string `xml:"full,omitempty"`
			Filtered string `xml:"filtered,omitempty"`
		} `xml:"warning,omitempty"`
	} `xml:"result_count,omitempty"`
	Task *Task `xml:"task,omitempty"`
	Scan *struct {
		Text string `xml:",chardata"`
		Task string `xml:"task,omitempty"`
	} `xml:"scan,omitempty"`
	Timestamp      string   `xml:"timestamp,omitempty"`
	ScanStart      string   `xml:"scan_start,omitempty"`
	Timezone       string   `xml:"timezone,omitempty"`
	TimezoneAbbrev string   `xml:"timezone_abbrev,omitempty"`
	Ports          *Ports   `xml:"ports,omitempty"`
	Results        *Results `xml:"results,omitempty"`
	HostStart      *Host    `xml:"host_start,omitempty"`
	HostEnd        *Host    `xml:"host_end,omitempty"`
	ScanEnd        string   `xml:"scan_end,omitempty"`
	Errors         *Errors  `xml:"errors,omitempty"`
}

// GetReports get report by its ID
func (conn *Connector) GetReports(id string) (Report, error) {
	req := GetReports{}
	res := Report{}

	resp := &GetReportsResponse{}
	err := conn.doRequest(req, resp)
	if err != nil {
		return res, err
	}

	res = *resp.Report
	return res, nil
}
