// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/squat/terraform-provider-datacrunch/internal/sdk/pkg/utils"
)

type CPU struct {
	Description   *string  `json:"description,omitempty"`
	NumberOfCores *float64 `json:"number_of_cores,omitempty"`
}

func (o *CPU) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CPU) GetNumberOfCores() *float64 {
	if o == nil {
		return nil
	}
	return o.NumberOfCores
}

type Gpu struct {
	Description  *string  `json:"description,omitempty"`
	NumberOfGpus *float64 `json:"number_of_gpus,omitempty"`
}

func (o *Gpu) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Gpu) GetNumberOfGpus() *float64 {
	if o == nil {
		return nil
	}
	return o.NumberOfGpus
}

type GpuMemory struct {
	Description     *string  `json:"description,omitempty"`
	SizeInGigabytes *float64 `json:"size_in_gigabytes,omitempty"`
}

func (o *GpuMemory) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GpuMemory) GetSizeInGigabytes() *float64 {
	if o == nil {
		return nil
	}
	return o.SizeInGigabytes
}

type Memory struct {
	Description     *string  `json:"description,omitempty"`
	SizeInGigabytes *float64 `json:"size_in_gigabytes,omitempty"`
}

func (o *Memory) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Memory) GetSizeInGigabytes() *float64 {
	if o == nil {
		return nil
	}
	return o.SizeInGigabytes
}

type Storage struct {
	Description     *string  `json:"description,omitempty"`
	SizeInGigabytes *float64 `json:"size_in_gigabytes,omitempty"`
}

func (o *Storage) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Storage) GetSizeInGigabytes() *float64 {
	if o == nil {
		return nil
	}
	return o.SizeInGigabytes
}

// Instance details object
type Instance struct {
	CPU         *CPU       `json:"cpu,omitempty"`
	CreatedAt   *string    `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	Gpu         *Gpu       `json:"gpu,omitempty"`
	GpuMemory   *GpuMemory `json:"gpu_memory,omitempty"`
	Hostname    *string    `json:"hostname,omitempty"`
	// id of an instance
	ID *string `json:"id,omitempty"`
	// the image type (operating system)
	Image *Image `json:"image,omitempty"`
	// the instance type
	InstanceType *InstanceType `json:"instance_type,omitempty"`
	IP           *string       `json:"ip,omitempty"`
	// Datacenter Location
	LocationCode *LocationCode `default:"FIN-01" json:"location_code"`
	Memory       *Memory       `json:"memory,omitempty"`
	// id of a storage volume
	OsVolumeID   *string  `json:"os_volume_id,omitempty"`
	PricePerHour *float64 `json:"price_per_hour,omitempty"`
	// array of ssh key ids
	SSHKeyIds []string `json:"ssh_key_ids,omitempty"`
	// startup script id. you need to first add a startup script to datacrunch via the startup script endpoint
	StartupScriptID *string         `json:"startup_script_id,omitempty"`
	Status          *InstanceStatus `json:"status,omitempty"`
	Storage         *Storage        `json:"storage,omitempty"`
}

func (i Instance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Instance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Instance) GetCPU() *CPU {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *Instance) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Instance) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Instance) GetGpu() *Gpu {
	if o == nil {
		return nil
	}
	return o.Gpu
}

func (o *Instance) GetGpuMemory() *GpuMemory {
	if o == nil {
		return nil
	}
	return o.GpuMemory
}

func (o *Instance) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *Instance) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Instance) GetImage() *Image {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *Instance) GetInstanceType() *InstanceType {
	if o == nil {
		return nil
	}
	return o.InstanceType
}

func (o *Instance) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *Instance) GetLocationCode() *LocationCode {
	if o == nil {
		return nil
	}
	return o.LocationCode
}

func (o *Instance) GetMemory() *Memory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *Instance) GetOsVolumeID() *string {
	if o == nil {
		return nil
	}
	return o.OsVolumeID
}

func (o *Instance) GetPricePerHour() *float64 {
	if o == nil {
		return nil
	}
	return o.PricePerHour
}

func (o *Instance) GetSSHKeyIds() []string {
	if o == nil {
		return nil
	}
	return o.SSHKeyIds
}

func (o *Instance) GetStartupScriptID() *string {
	if o == nil {
		return nil
	}
	return o.StartupScriptID
}

func (o *Instance) GetStatus() *InstanceStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Instance) GetStorage() *Storage {
	if o == nil {
		return nil
	}
	return o.Storage
}