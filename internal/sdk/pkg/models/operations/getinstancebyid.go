// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-datacrunch/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetInstanceByIDRequest struct {
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceId"`
}

func (o *GetInstanceByIDRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

type GetInstanceByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Instance details object
	Instance *shared.Instance
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetInstanceByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInstanceByIDResponse) GetInstance() *shared.Instance {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *GetInstanceByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInstanceByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
