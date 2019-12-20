package omp

import (
	"encoding/xml"
	"errors"

	"github.com/jinzhu/copier"
)

// Scanner .
type Scanner struct {
	Text             string    `xml:",chardata"`
	ID               string    `xml:"id,attr,omitempty"`
	Name             string    `xml:"name,omitempty"`
	Comment          string    `xml:"comment,omitempty"`
	CreationTime     string    `xml:"creation_time,omitempty"`
	ModificationTime string    `xml:"modification_time,omitempty"`
	Writable         string    `xml:"writable,omitempty"`
	InUse            string    `xml:"in_use,omitempty"`
	UserTags         *UserTags `xml:"user_tags,omitempty"`
	Host             string    `xml:"host,omitempty"`
	Port             string    `xml:"port,omitempty"`
	Type             string    `xml:"type,omitempty"`
	CaPub            string    `xml:"ca_pub,omitempty"`
	KeyPub           string    `xml:"key_pub,omitempty"`
	Tasks            *struct {
		Text string `xml:",chardata"`
		Task *struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr,omitempty"`
			Name string `xml:"name,omitempty"`
		} `xml:"task,omitempty"`
	} `xml:"tasks,omitempty"`
}

// CreateScanner .
type CreateScanner struct {
	requestBase
	XMLName    xml.Name    `xml:"create_scanner"`
	Text       string      `xml:",chardata"`
	Name       string      `xml:"name,omitempty"`
	Host       string      `xml:"host,omitempty"`
	Port       string      `xml:"port,omitempty"`
	Type       string      `xml:"type,omitempty"`
	CaPub      string      `xml:"ca_pub,omitempty"`
	Credential *Credential `xml:"credential,omitempty"`
}

// CreateScannerResponse .
type CreateScannerResponse struct {
	responseBase
	XMLName xml.Name `xml:"create_scanner_response"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr,omitempty"`
}

// CreateScanner .
func (c *Connector) CreateScanner(s *Scanner, credID string) (string, error) {
	req := CreateScanner{}
	req.Credential = &Credential{}
	copier.Copy(&req, &s)

	req.Credential.ID = credID
	resp := &CreateScannerResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

// GetScanners .
type GetScanners struct {
	requestBase
	XMLName   xml.Name `xml:"get_scanners"`
	Text      string   `xml:",chardata"`
	Details   string   `xml:"details,attr,omitempty"`
	ScannerID string   `xml:"scanner_id,attr,omitempty"`
}

// GetScannersResponse .
type GetScannersResponse struct {
	responseBase
	XMLName  xml.Name  `xml:"get_scanners_response"`
	Text     string    `xml:",chardata"`
	Scanner  []Scanner `xml:"scanner,omitempty"`
	Truncate string    `xml:"truncate,omitempty"`
}

// GetScanners .
func (c *Connector) GetScanners() ([]Scanner, error) {
	res := []Scanner{}
	req := GetScanners{}
	resp := &GetScannersResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return res, err
	}
	for i := range resp.Scanner {
		res = append(res, resp.Scanner[i])
	}
	return res, nil
}

// GetDefaultScanner return the ID of default openvas scanner
func (c *Connector) GetDefaultScanner() (string, error) {
	return c.GetScannerByName("OpenVAS Default")
}

func (c *Connector) GetScannerByName(scannerName string) (string, error) {
	scannerLst, err := c.GetScanners()
	if err != nil {
		return "", nil
	}
	for i := range scannerLst {
		if scannerLst[i].Name == scannerName {
			return scannerLst[i].ID, nil
		}
	}
	return "", errors.New("no such scanner with this name", scannerName)
}
