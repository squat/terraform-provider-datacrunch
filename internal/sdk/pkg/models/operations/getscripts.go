// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-datacrunch/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetScriptsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// an array of startup script objects
	StartupScripts []shared.StartupScript
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetScriptsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetScriptsResponse) GetStartupScripts() []shared.StartupScript {
	if o == nil {
		return nil
	}
	return o.StartupScripts
}

func (o *GetScriptsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetScriptsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
