package api

import (
	"fmt"
	"github.com/propero-oss/csi-vcloud/pkg/common"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	"net/url"
)


func Client(c *common.Config) (*govcd.VCDClient, error) {
	u, err := url.ParseRequestURI(c.VCloud.API)
	if err != nil {
		return nil, fmt.Errorf("unable to pass url: %s", err)
	}

	vcdclient := govcd.NewVCDClient(*u, c.VCloud.Insecure)
	err = vcdclient.Authenticate(c.VCloud.Username, c.VCloud.Password, c.VCloud.ORG)
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