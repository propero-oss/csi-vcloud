package service

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/propero-oss/csi-vcloud/pkg/common"
	"github.com/rexray/gocsi"
	"k8s.io/klog"
	"net"
	"os"
	"strings"
)

const (
	// Name is the name of the CSI plug-in.
	Name = "csi-vcloud.propero-oss.dev"

	// VendorVersion is the version returned by GetPluginInfo.
	VendorVersion = "0.0.1"

	UnixSocketPrefix = "unix://"
)

var Manifest = map[string]string{
	"url": "https://github.com/rexray/gocsi/tree/master/mock",
}

// Service is the CSI Mock service provider.
type Service interface {
	BeforeServe(context.Context, *gocsi.StoragePlugin, net.Listener) error
	GetController() csi.ControllerServer
	csi.IdentityServer
	csi.NodeServer
}

type service struct {
	mode string
	cs Controller
}


func (s *service) BeforeServe(ctx context.Context, plugin *gocsi.StoragePlugin, listener net.Listener) error {
	defer func() {
		fields := map[string]interface{}{
			"mode": s.mode,
		}
		klog.V(2).Infof("configured: %s with %+v", Name, fields)
	}()

	s.mode = os.Getenv("X_CSI_MODE")

	if !strings.EqualFold(s.mode, "node") {

		cfg, err  := common.ParseConfig()
		if err != nil {
			klog.Errorf("Failed to read config. Error: %v", err)
			return err
		}
		if err := s.cs.Init(cfg); err != nil {
			klog.Errorf("Failed to init controller. Error: %v", err)
			return err
		}
	}

	return nil
}

func (s* service) GetController() csi.ControllerServer {
	s.cs = New()
	return s.cs
}

func NewService() Service {
	return &service{}
}