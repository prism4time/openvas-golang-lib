package omp

import (
	"encoding/xml"
	"strings"
)

// OMPtarget .
type OMPtarget struct {
	Text             string    `xml:",chardata"`
	ID               string    `xml:"id,attr,omitempty"`
	Name             string    `xml:"name,omitempty"`
	Comment          string    `xml:"comment,omitempty"`
	CreationTime     string    `xml:"creation_time,omitempty"`
	ModificationTime string    `xml:"modification_time,omitempty"`
	Writable         string    `xml:"writable,omitempty"`
	InUse            string    `xml:"in_use,omitempty"`
	UserTags         *UserTags `xml:"user_tags,omitempty"`
	Hosts            string    `xml:"hosts,omitempty"`
	MaxHosts         string    `xml:"max_hosts,omitempty"`
	SSHCredential    *struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr,omitempty"`
		Name string `xml:"name,omitempty"`
	} `xml:"ssh_credential,omitempty"`
	SmbCredential *struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr,omitempty"`
		Name string `xml:"name,omitempty"`
	} `xml:"smb_credential,omitempty"`
	EsxiCredential *struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr,omitempty"`
		Name string `xml:"name,omitempty"`
	} `xml:"esxi_credential,omitempty"`
}

// Target .
type Target struct {
	Text         string `xml:",chardata"`
	ID           string `xml:"id,attr,omitempty"`
	Name         string `xml:"name,omitempty"`
	Hosts        []string
	ExcludeHosts []string
}

// CreateTarget .
type CreateTarget struct {
	requestBase
	XMLName xml.Name `xml:"create_target"`
	Name    string   `xml:"name,omitempty"`
	Hosts   string   `xml:"hosts,omitempty"`
}

type responseCreateTarget struct {
	responseBase
	ID string `xml:"id,attr,omitempty"`
}

// ModifyTarget .
type ModifyTarget struct {
	requestBase
	XMLName      xml.Name `xml:"modify_target"`
	ID           string   `xml:"target_id,attr,omitempty"`
	Name         string   `xml:"name,omitempty"`
	Hosts        string   `xml:"hosts,omitempty"`
	ExcludeHosts string   `xml:"exclude_hosts,omitempty"`
}

type responseModifyTarget struct {
	responseBase
}

// GetTargets .
type GetTargets struct {
	requestBase
	XMLName xml.Name `xml:"get_targets"`
}

// GetTargetsResponse .
type GetTargetsResponse struct {
	responseBase
	XMLName xml.Name    `xml:"get_targets_response"`
	Text    string      `xml:",chardata"`
	Target  []OMPtarget `xml:"target,omitempty"`
}

// DeleteTarget .
type DeleteTarget struct {
	requestBase
	XMLName xml.Name `xml:"delete_target"`
	ID      string   `xml:"target_id,attr,omitempty"`
}

type responseDeleteTarget struct {
	responseBase
}

// CreateTarget create a target and return its ID
func (c *Connector) CreateTarget(t *Target) (string, error) {
	req := CreateTarget{
		Name:  t.Name,
		Hosts: strings.Join(t.Hosts, ","),
	}
	resp := &responseCreateTarget{}
	err := c.doRequest(req, resp)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

// ModifyTarget .
func (c *Connector) ModifyTarget(t *Target) error {
	req := ModifyTarget{
		ID: t.ID,
	}
	if t.Hosts != nil {
		req.Hosts = strings.Join(t.Hosts, ",")
	}
	if t.ExcludeHosts != nil {
		req.ExcludeHosts = strings.Join(t.ExcludeHosts, ",")
	}
	resp := &responseModifyTarget{}
	err := c.doRequest(req, resp)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTarget .
func (c *Connector) DeleteTarget(id string) error {
	req := DeleteTarget{
		ID: id,
	}
	resp := &responseDeleteTarget{}
	err := c.doRequest(req, resp)
	if err != nil {
		return err
	}
	return nil
}

// GetTargets .
func (c *Connector) GetTargets() ([]OMPtarget, error) {
	req := GetTargets{}
	res := []OMPtarget{}

	resp := &GetTargetsResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return res, err
	}

	targets := resp.Target
	for i := range targets {
		config := OMPtarget{
			ID:    targets[i].ID,
			Name:  targets[i].Name,
			Hosts: targets[i].Hosts,
		}
		res = append(res, config)
	}
	return res, nil
}
