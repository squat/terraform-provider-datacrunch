// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-datacrunch/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSSHKeysResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns an array of ssh key objects
	SSHKeys []shared.SSHKey
}

func (o *GetSSHKeysResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSSHKeysResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSSHKeysResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSSHKeysResponse) GetSSHKeys() []shared.SSHKey {
	if o == nil {
		return nil
	}
	return o.SSHKeys
}
