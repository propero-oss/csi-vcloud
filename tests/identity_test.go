package tests

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

var _ = Describe("Identity", func() {
	var (
		err      error
		stopMock func()
		ctx      context.Context
		gclient  *grpc.ClientConn
		client   csi.IdentityClient
	)
	BeforeEach(func() {
		ctx = context.Background()
	})
	JustBeforeEach(func() {
		gclient, stopMock, err = startMockServer(ctx)
		Ω(err).ShouldNot(HaveOccurred())
		client = csi.NewIdentityClient(gclient)
	})
	AfterEach(func() {
		ctx = nil
		gclient.Close()
		gclient = nil
		client = nil
		stopMock()
	})

	Describe("Probe", func() {
		It("Should Be Ready", func() {
			rep, err := client.Probe(ctx, &csi.ProbeRequest{})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(rep).ShouldNot(BeNil())
			Ω(rep.GetReady().GetValue()).To(Equal(true))
		})
	})
})