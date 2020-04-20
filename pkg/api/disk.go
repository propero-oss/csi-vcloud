package api

import (
	"fmt"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

func AttachDiskToVM(vm *govcd.VM, diskRef *types.Reference) {
	var diskAttachParams *types.DiskAttachOrDetachParams
	busNumber := 1
	unitNumber := 0
	diskAttachParams = &types.DiskAttachOrDetachParams{
		Disk: diskRef,
		BusNumber: &busNumber,
		UnitNumber: &unitNumber,
	}

	task, err := vm.AttachDisk(diskAttachParams)
	if err != nil {
		fmt.Print(err)
	}


	err = task.WaitTaskCompletion()
	if err != nil {
		fmt.Print(fmt.Errorf("Error: %s", err))
	}
}
