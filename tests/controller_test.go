package tests

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

var _ = Describe("Controller", func() {
	var (
		err      error
		stopMock func()
		ctx      context.Context
		gclient  *grpc.ClientConn
		client   csi.ControllerClient
	)
	BeforeEach(func() {
		ctx = context.Background()
	})
	JustBeforeEach(func() {
		gclient, stopMock, err = startMockServer(ctx)
		Ω(err).ShouldNot(HaveOccurred())
		client = csi.NewControllerClient(gclient)
	})
	AfterEach(func() {
		ctx = nil
		gclient.Close()
		gclient = nil
		client = nil
		stopMock()
	})
	Describe("GetCapacity", func() {
		It("GetCapacity should not be zero", func() {
			rep, err := client.GetCapacity(ctx, &csi.GetCapacityRequest{})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(rep).ShouldNot(BeNil())
			Ω(rep.AvailableCapacity).ShouldNot(BeZero())
		})
	})
})