package api

import (
	"fmt"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)


func (vdc *Vdc) GetVMByVAppName(vAppName string, vmName string) (*govcd.VM, error) {

	vapp, err := vdc.GetVAppByName(vAppName, true)
	if err != nil {
		return nil, fmt.Errorf("GetVAppByName error")
	}
	vm, err := vapp.GetVMByName(vmName, true)
	if err != nil {
		return nil, fmt.Errorf("could not find VM")
	}

	return vm, nil
}

// CreateIndependentDisk
func (vdc *Vdc) CreateIndependentDisk(diskName string, storageProfile string, size int64) *types.Reference{
	var diskCreateParams *types.DiskCreateParams
	diskCreateParams = &types.DiskCreateParams{Disk: &types.Disk{
		Name: diskName,
		// size is expected to be in MiB therefore we convert it to bytes
		Size: size * int64(MiB),
	}}

	profileRef, err := vdc.FindStorageProfileReference(storageProfile)
	if err != nil {
		fmt.Print("")
	}

	diskCreateParams.Disk.StorageProfile = &profileRef
	diskCreateParams.Disk.BusSubType = "lsilogic"
	diskCreateParams.Disk.BusType = SCSI
	diskCreateParams.Disk.Description = "GOVCD Test"

	task, err := vdc.CreateDisk(diskCreateParams)
	if err != nil {
		fmt.Print(err)
	}

	err = task.WaitTaskCompletion()
	if err != nil {
		fmt.Print(fmt.Errorf("error: %s", err))
	}

	return task.Task.Owner
}


func (m *Manager) GetVDC(orgName string, vdcName string) (*govcd.Vdc, error) {
	if m.Client == nil {
		return nil, fmt.Errorf("m.Client is nil")
	}

	org, err := m.Client.GetOrgByName(orgName)

	if err != nil {
		return nil, fmt.Errorf("unable to GetOrgByName: %s", err)
	}
	vdc, err := org.GetVDCByName(vdcName, true)
	if err != nil {
		return nil, fmt.Errorf("unable to GetVDCByName: %s", err)
	}

	return vdc, nil
}
