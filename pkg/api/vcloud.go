package api

import (
	"fmt"
	"github.com/vmware/go-vcloud-director/v2/govcd"
    "github.com/vmware/go-vcloud-director/v2/types/v56"
	"net/url"
	"os"
)



func (c *Config) Client() (*govcd.VCDClient, error) {
	u, err := url.ParseRequestURI(c.Href)
	if err != nil {
		return nil, fmt.Errorf("unable to pass url: %s", err)
	}

	vcdclient := govcd.NewVCDClient(*u, c.Insecure)
	err = vcdclient.Authenticate(c.User, c.Password, c.Org)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}

	return vcdclient, nil
}

func (m *Manager) GetStorageProfile(storageHref string) (*types.VdcStorageProfile, error) {
	storageProfile, err := govcd.GetStorageProfileByHref(m.Client, "")
	if err != nil {
		return nil, fmt.Errorf("unable to get StorageProfile by HREF: %s", err)
	}
	return storageProfile, nil
}




func TestAuth() {

	config := Config{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Org:      os.Getenv("ORG"),
		Href:     os.Getenv("API"),
		VDC:      os.Getenv("VDC"),
		VApp:     os.Getenv("VAPP"),
		Insecure: false,
	}

	client, err := config.Client() // We now have a client

	obj := Manager{ Client: client, Config: &config}
	fmt.Printf("\n\nDEBUG: %s\n\n", obj.Config.Org)

	vdc, err := obj.GetVDC(os.Getenv("ORG"), os.Getenv("VDC"))
	obj.Vdc.Vdc = vdc
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	myVdc := Vdc{vdc}
	vm, err := myVdc.GetVMByVAppName("worker-nodes-nschad-cluster", "worker-node-0")

	vmwrapper := VM{vm}
	vmwrapper.test()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	blub, err  := vm.GetMetadata()


	/*diskRef := myVdc.CreateIndependentDisk("testdisk-govcd", os.Getenv("STORAGE_PROFILE"), 1024)
	AttachDiskToVM(vm, diskRef)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	org, err := client.GetOrgByName(config.Org)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vdc, err = org.GetVDCByName(config.VDC, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Org URL: %s\n", org.Org.HREF)
	fmt.Printf("VDC URL: %s\n", vdc.Vdc.HREF)*/
}
