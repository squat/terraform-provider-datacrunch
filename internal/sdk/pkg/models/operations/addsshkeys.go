// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// AddSSHKeysRequestBody - ssh key details
type AddSSHKeysRequestBody struct {
	// the actual public ssh key value
	Key string `json:"key"`
	// ssh key name
	Name string `json:"name"`
}

func (o *AddSSHKeysRequestBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *AddSSHKeysRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type AddSSHKeysResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Key added, returns the internal ssh key id
	SSHKeyID *string
}

func (o *AddSSHKeysResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddSSHKeysResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddSSHKeysResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddSSHKeysResponse) GetSSHKeyID() *string {
	if o == nil {
		return nil
	}
	return o.SSHKeyID
}