package omp

import (
	"encoding/xml"

	"github.com/jinzhu/copier"
)

// Credential .
type Credential struct {
	Text             string    `xml:",chardata"`
	Password         string    `xml:"password,omitempty"`
	ID               string    `xml:"id,attr,omitempty"`
	Name             string    `xml:"name,omitempty"`
	Login            string    `xml:"login,omitempty"`
	Comment          string    `xml:"comment,omitempty"`
	CreationTime     string    `xml:"creation_time,omitempty"`
	ModificationTime string    `xml:"modification_time,omitempty"`
	Writable         string    `xml:"writable,omitempty"`
	InUse            string    `xml:"in_use,omitempty"`
	UserTags         *UserTags `xml:"user_tags,omitempty"`
	Type             string    `xml:"type,omitempty"`
	FullType         string    `xml:"full_type,omitempty"`
}

// CreateCredential .
type CreateCredential struct {
	requestBase
	XMLName  xml.Name `xml:"create_credential"`
	Text     string   `xml:",chardata"`
	Name     string   `xml:"name,omitempty"`
	Login    string   `xml:"login,omitempty"`
	Password string   `xml:"password,omitempty"`
	Comment  string   `xml:"comment,omitempty"`
}

// CreateCredentialResponse .
type CreateCredentialResponse struct {
	responseBase
	XMLName xml.Name `xml:"create_credential_response"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr,omitempty"`
}

// GetCredentials .
type GetCredentials struct {
	requestBase
	XMLName      xml.Name `xml:"get_credentials"`
	Text         string   `xml:",chardata"`
	CredentialID string   `xml:"credential_id,attr,omitempty"`
	Format       string   `xml:"format,attr,omitempty"`
	Targets      string   `xml:"targets,attr,omitempty"`
}

// GetCredentialsResponse .
type GetCredentialsResponse struct {
	responseBase
	XMLName    xml.Name     `xml:"get_credentials_response"`
	Text       string       `xml:",chardata"`
	Credential []Credential `xml:"credential,omitempty"`
}

// CreateCredential .
func (conn *Connector) CreateCredential(cred Credential) (string, error) {
	req := CreateCredential{}
	copier.Copy(&req, &cred)

	resp := &CreateCredentialResponse{}
	err := conn.doRequest(req, resp)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

// GetCredentials .
func (conn *Connector) GetCredentials( /*cred Credential*/ ) ([]Credential, error) {
	res := []Credential{}
	req := GetCredentials{}
	// copier.Copy(&req, &cred)

	resp := &GetCredentialsResponse{}
	err := conn.doRequest(req, resp)
	if err != nil {
		return res, err
	}

	for i := range resp.Credential {
		res = append(res, resp.Credential[i])
	}
	return res, nil
}
