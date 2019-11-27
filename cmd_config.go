package omp

import "encoding/xml"

// Config .
type Config struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,omitempty"`
}

// CreateConfig .
type CreateConfig struct {
	requestBase
	XMLName xml.Name `xml:"create_config"`
	Text    string   `xml:",chardata"`
	Copy    string   `xml:"copy,omitempty"`
	Name    string   `xml:"name,omitempty"`
	// GetConfigsResponse GetConfigsResponse `xml:"get_configs_response"`
}

// CreateConfigResponse .
type CreateConfigResponse struct {
	responseBase
	XMLName xml.Name `xml:"create_config_response"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr,omitempty"`
}

// GetConfigs .
type GetConfigs struct {
	requestBase
	XMLName xml.Name `xml:"get_configs"`
}

// GetConfigsResponse .
type GetConfigsResponse struct {
	responseBase
	XMLName xml.Name `xml:"get_configs_response"`
	Text    string   `xml:",chardata"`
	Config  []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id,attr,omitempty"`
		Name             string `xml:"name,omitempty"`
		Comment          string `xml:"comment,omitempty"`
		CreationTime     string `xml:"creation_time,omitempty"`
		ModificationTime string `xml:"modification_time,omitempty"`
		FamilyCount      *struct {
			Text    string `xml:",chardata"`
			Growing string `xml:"growing,omitempty"`
		} `xml:"family_count,omitempty"`
		NvtCount *struct {
			Text    string `xml:",chardata"`
			Growing string `xml:"growing,omitempty"`
		} `xml:"nvt_count,omitempty"`
		InUse    string    `xml:"in_use,omitempty"`
		Writable string    `xml:"writable,omitempty"`
		UserTags *UserTags `xml:"user_tags,omitempty"`
	} `xml:"config,omitempty"`
}

// CreateConfig .
func (conn *Connector) CreateConfig(id string, name string) (string, error) {
	req := CreateConfig{
		Copy: id,
		Name: name,
	}
	resp := &CreateConfigResponse{}
	err := conn.doRequest(req, resp)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

// GetConfigs .
func (conn *Connector) GetConfigs() ([]Config, error) {
	req := GetConfigs{}
	res := []Config{}

	resp := &GetConfigsResponse{}
	err := conn.doRequest(req, resp)
	if err != nil {
		return res, err
	}

	configs := resp.Config
	for i := range configs {
		config := Config{
			ID:   configs[i].ID,
			Name: configs[i].Name,
		}
		res = append(res, config)
	}
	return res, nil
}
