package service

import (
	"context"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/propero-oss/csi-vcloud/pkg/api"
	"github.com/propero-oss/csi-vcloud/pkg/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
	"os"
)

type Controller interface {
	csi.ControllerServer
	Init(config *common.Config) error
}

type controller struct {
	manager *api.Manager
}


func New() Controller {
	return &controller{}
}

var (
	// controllerCaps represents the capability of controller service
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	}
)

func (c *controller) Init(config *common.Config) error {
	klog.Info("Initializing csi-vcloud controller")

	client, err := api.Client(config)
	if err != nil {
		panic(err)
	}

	c.manager.Client = client

	vdc, err := c.manager.GetVDC(config.VCloud.ORG, config.VCloud.VDC)
	if err != nil {
		panic(err)
	}

	c.manager = &api.Manager{
		Client: client,
		Config: config,
		Vdc: vdc,
	}

	return nil

}

func (c *controller) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.V(4).Infof("CreateVolume: called with args %+v", *req)
	volName := req.GetName()
	if len(volName) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume name not provided")
	}

	volSizeBytes := int64(common.DEFAULT_VOLUME_SIZE * common.GBinBytes)
	if req.GetCapacityRange() != nil && req.GetCapacityRange().RequiredBytes != 0 {
		volSizeBytes = req.GetCapacityRange().GetRequiredBytes()
	}
	disk := c.manager.CreateIndependentDisk(req.Name, os.Getenv("STORAGE_PROFILE"), volSizeBytes)

	attributes := make(map[string]string)
	
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes: volSizeBytes,
			VolumeId:      disk.ID,
			VolumeContext: attributes,
		},
	}, nil
}

func (c *controller) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.V(4).Infof("DeleteVolume: called with args: %+v", *req)

	volId := req.VolumeId

	err := c.manager.DeleteDisk(volId)
	/*
	DeleteVolume:
	This operation MUST be idempotent.
	If a volume corresponding to the specified volume_id does not exist or the artifacts associated with the volume do not exist anymore,
	the Plugin MUST reply 0 OK.
	 */
	if err != nil {
		msg := fmt.Sprintf("Failed to delete volume: %q. Error: %+v", req.VolumeId, err)
		klog.Error(msg)
		//return nil, status.Errorf(codes.Internal, msg)
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (c *controller) ControllerPublishVolume(ctx context.Context, request *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	panic("implement me")
}

func (c *controller) ControllerUnpublishVolume(ctx context.Context, request *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	panic("implement me")
}

func (c *controller) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	klog.V(4).Infof("ControllerGetCapabilities: called with args %+v", *req)
	volCaps := req.GetVolumeCapabilities()
	var confirmed *csi.ValidateVolumeCapabilitiesResponse_Confirmed
	if common.IsValidVolumeCapabilities(volCaps) {
		confirmed = &csi.ValidateVolumeCapabilitiesResponse_Confirmed{VolumeCapabilities: volCaps}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: confirmed,
	}, nil
}

func (c *controller) ListVolumes(ctx context.Context, request *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	//TODO: Get all Worker Nodes in the VDC and get all IndepentDisk attached to the Worker Nodes
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *controller) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	//TODO: Figure out if you can query the storage capacity from the VDC

	klog.V(4).Infof("GetCapacity: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *controller) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	klog.V(4).Infof("ControllerGetCapabilities: called with args %+v", *req)
	var caps []*csi.ControllerServiceCapability
	for _, controllerCap := range controllerCaps {
		c := &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: controllerCap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.ControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (c *controller) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	klog.V(4).Infof("CreateSnapshot: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *controller) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	klog.V(4).Infof("CreateSnapshot: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *controller) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	klog.V(4).Infof("CreateSnapshot: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *controller) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	klog.V(4).Infof("CreateSnapshot: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}