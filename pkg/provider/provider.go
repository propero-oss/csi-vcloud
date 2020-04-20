package provider

import (
	"context"
	"github.com/rexray/gocsi"
	"github.com/rexray/gocsi/mock/service"
	"github.com/sirupsen/logrus"
	"net"
)

func New() gocsi.StoragePluginProvider {
	svc := service.New()
	return &gocsi.StoragePlugin{
		Controller:  svc,
		Identity:    svc,
		Node:        svc,
		BeforeServe: func(ctx context.Context, plugin *gocsi.StoragePlugin, listener net.Listener) error {
			logrus.WithField("service", service.Name).Debug("BeforeServe")
			return nil
		},
		EnvVars:     []string{
			// Enable serial volume access.
			gocsi.EnvVarSerialVolAccess + "=true",

			// Enable request and response validation.
			gocsi.EnvVarSpecValidation + "=true",

			// Treat the following fields as required:
			//   * ControllerPublishVolumeResponse.PublishContext
			//   * NodeStageVolumeRequest.PublishContext
			//   * NodePublishVolumeRequest.PublishContext
			gocsi.EnvVarRequirePubContext + "=true",
		},
	}
}
