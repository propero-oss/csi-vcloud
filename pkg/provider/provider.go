package provider

import (
	"github.com/propero-oss/csi-vcloud/pkg/service"
	"github.com/rexray/gocsi"
)

func New() gocsi.StoragePluginProvider {
	svc := service.NewService()
	ctrl := svc.GetController()
	return &gocsi.StoragePlugin{
		Controller:  ctrl,
		Identity:    svc,
		Node:        svc,
		BeforeServe: svc.BeforeServe,
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
