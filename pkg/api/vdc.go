package api

import (
	"fmt"
	"github.com/propero-oss/csi-vcloud/pkg/common"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)


func (m *Manager) GetVMByVAppName(vAppName string, vmName string) (*govcd.VM, error) {

	vapp, err := m.Vdc.GetVAppByName(vAppName, true)
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
func (m *Manager) CreateIndependentDisk(diskName string, storageProfile string, size int64) (*types.Reference, error){
	var diskCreateParams *types.DiskCreateParams
	diskCreateParams = &types.DiskCreateParams{Disk: &types.Disk{
		Name: diskName,
		// size is expected to be in MiB therefore we convert it to bytes
		Size: size * int64(common.MBinBytes),
	}}
	
	profileRef, err := m.Vdc.FindStorageProfileReference(storageProfile)
	if err != nil {
		return nil, err
	}

	diskCreateParams.Disk.StorageProfile = &profileRef
	diskCreateParams.Disk.BusSubType = "lsilogic"
	diskCreateParams.Disk.BusType = SCSI
	diskCreateParams.Disk.Description = "GOVCD Test"

	task, err := m.Vdc.CreateDisk(diskCreateParams)
	if err != nil {
		return nil, err
	}

	err = task.WaitTaskCompletion()
	if err != nil {
		return nil, err
	}

	return task.Task.Owner, nil
}

func (m *Manager) DeleteDisk(volId string) error {
	disk, err := m.Vdc.GetDiskById(volId, true)
	if err != nil {
		return err
	}

	task, err := disk.Delete()
	if err != nil {
		return err
	}

	err = task.WaitTaskCompletion()
	if err != nil {
		return err
	}

	return nil
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
