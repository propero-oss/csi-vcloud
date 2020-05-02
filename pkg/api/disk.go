package api

import (
	"fmt"
	"github.com/propero-oss/csi-vcloud/pkg/common"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

func AttachDiskToVM(vm *govcd.VM, diskRef *types.Reference) {
	var diskAttachParams *types.DiskAttachOrDetachParams

	diskAttachParams = &types.DiskAttachOrDetachParams{
		Disk:       diskRef,
	}

	task, err := vm.AttachDisk(diskAttachParams)
	if err != nil {
		fmt.Print(err)
	}

	err = task.WaitTaskCompletion()
	if err != nil {
		fmt.Print(fmt.Errorf("error: %s", err))
	}
}

type DiskInfo struct {
	BusNumber, UnitNumber int
}


func contains(arr []DiskInfo, busNumber int, unitNumber int) bool {
	for _, a := range arr {
		if (a.UnitNumber == unitNumber) && (a.BusNumber == busNumber) {
			return true
		}
	}
	return false
}

func FindNextBusAndUnitNumber(vm *govcd.VM) (*DiskInfo, error) {
	var reserved []DiskInfo

	diskSettings := vm.VM.VmSpecSection.DiskSection.DiskSettings
	for _, disk := range diskSettings {
		reserved = append(reserved, DiskInfo{
			BusNumber:  disk.BusNumber,
			UnitNumber: disk.UnitNumber,
		})
	}

	for busNumber := 0; busNumber <= common.MAX_BUS_NUMBER; busNumber++ {
		for unitNumber := 0; unitNumber <= common.MAX_UNIT_NUMBER; unitNumber++ {
			if !contains(reserved, busNumber, unitNumber) {
				return &DiskInfo{
					BusNumber:  busNumber,
					UnitNumber: unitNumber,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("all controllers are fully utilized")
}
