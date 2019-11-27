package omp

import "encoding/xml"

// Result .
type Result struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id,attr,omitempty"`
	Owner *struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,omitempty"`
	} `xml:"owner,omitempty"`
	Comment          string    `xml:"comment,omitempty"`
	CreationTime     string    `xml:"creation_time,omitempty"`
	ModificationTime string    `xml:"modification_time,omitempty"`
	UserTags         *UserTags `xml:"user_tags,omitempty"`
	Host             string    `xml:"host,omitempty"`
	Port             string    `xml:"port,omitempty"`
	Nvt              *struct {
		Text     string `xml:",chardata"`
		Oid      string `xml:"oid,attr,omitempty"`
		Name     string `xml:"name,omitempty"`
		CvssBase string `xml:"cvss_base,omitempty"`
		Cve      string `xml:"cve,omitempty"`
		Bid      string `xml:"bid,omitempty"`
		Tags     string `xml:"tags,omitempty"`
		Cert     string `xml:"cert,omitempty"`
		Xref     string `xml:"xref,omitempty"`
	} `xml:"nvt,omitempty"`
	Threat         string `xml:"threat,omitempty"`
	Description    string `xml:"description,omitempty"`
	OriginalThreat string `xml:"original_threat,omitempty"`
	Notes          string `xml:"notes,omitempty"`
	Overrides      *struct {
		Text     string `xml:",chardata"`
		Override *struct {
			Chardata string `xml:",chardata"`
			ID       string `xml:"id,attr,omitempty"`
			Nvt      *struct {
				Text string `xml:",chardata"`
				Oid  string `xml:"oid,attr,omitempty"`
				Name string `xml:"name,omitempty"`
			} `xml:"nvt,omitempty"`
			Text *struct {
				Text    string `xml:",chardata"`
				Excerpt string `xml:"excerpt,attr,omitempty"`
			} `xml:"text,omitempty"`
			NewThreat string `xml:"new_threat,omitempty"`
			Orphan    string `xml:"orphan,omitempty"`
		} `xml:"override,omitempty"`
	} `xml:"overrides,omitempty"`
}

// GetResults .
type GetResults struct {
	responseBase
	XMLName  xml.Name `xml:"get_results"`
	Text     string   `xml:",chardata"`
	ResultID string   `xml:"result_id,attr,omitempty"`
	TaskID   string   `xml:"task_id,attr,omitempty"`
	Filter   string   `xml:"filter,attr,omitempty"`
}

// GetResultsResponse .
type GetResultsResponse struct {
	responseBase
	XMLName xml.Name `xml:"get_results_response"`
	Text    string   `xml:",chardata"`
	Result  []Result `xml:"result,omitempty"`
	Filters *struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"id,attr,omitempty"`
		Term     string `xml:"term,omitempty"`
		Keywords *struct {
			Text    string `xml:",chardata"`
			Keyword *struct {
				Text     string `xml:",chardata"`
				Column   string `xml:"column,omitempty"`
				Relation string `xml:"relation,omitempty"`
				Value    string `xml:"value,omitempty"`
			} `xml:"keyword,omitempty"`
		} `xml:"keywords,omitempty"`
	} `xml:"filters,omitempty"`
	Sort *struct {
		Text  string `xml:",chardata"`
		Field *struct {
			Text  string `xml:",chardata"`
			Order string `xml:"order,omitempty"`
		} `xml:"field,omitempty"`
	} `xml:"sort,omitempty"`
	Results *struct {
		Text  string `xml:",chardata"`
		Max   string `xml:"max,attr,omitempty"`
		Start string `xml:"start,attr,omitempty"`
	} `xml:"results,omitempty"`
	ResultCount *struct {
		Text     string `xml:",chardata"`
		Filtered string `xml:"filtered,omitempty"`
		Page     string `xml:"page,omitempty"`
	} `xml:"result_count,omitempty"`
}

// GetResults .
func (c *Connector) GetResults(r *Result, taskID string, filter ...string) ([]Result, error) {
	res := []Result{}
	req := GetResults{}
	if filter != nil && len(filter) != 0 {
		req.Filter = filter[0]
	}
	if r != nil {
		if len(r.ID) != 0 {
			req.ResultID = r.ID
		}
	}
	if len(taskID) != 0 {
		req.TaskID = taskID
	}
	resp := &GetResultsResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return []Result{}, err
	}
	if len(resp.Result) != 0 {
		res = append(res, resp.Result...)
	}
	return res, nil
}
