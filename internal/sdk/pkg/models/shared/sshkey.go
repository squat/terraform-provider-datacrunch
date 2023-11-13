// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SSHKey - ssh key object
type SSHKey struct {
	// ssh key id
	ID string `json:"id"`
	// the actual public ssh key value
	Key string `json:"key"`
	// ssh key name
	Name string `json:"name"`
}

func (o *SSHKey) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SSHKey) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *SSHKey) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
