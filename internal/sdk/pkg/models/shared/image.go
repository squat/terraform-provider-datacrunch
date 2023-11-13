// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Image - the image type (operating system)
type Image string

const (
	ImageUbuntu1804        Image = "ubuntu-18.04"
	ImageUbuntu2004        Image = "ubuntu-20.04"
	ImageFastai            Image = "fastai"
	ImageUbuntu2004Cuda112 Image = "ubuntu-20.04-cuda-11.2"
)

func (e Image) ToPointer() *Image {
	return &e
}

func (e *Image) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ubuntu-18.04":
		fallthrough
	case "ubuntu-20.04":
		fallthrough
	case "fastai":
		fallthrough
	case "ubuntu-20.04-cuda-11.2":
		*e = Image(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Image: %v", v)
	}
}