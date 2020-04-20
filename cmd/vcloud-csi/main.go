package main

import (
	"flag"
	"github.com/getsentry/sentry-go"
	"github.com/propero-oss/csi-vcloud/pkg/api"
	"k8s.io/klog"
	"log"
	"os"
	"time"
)

func setupSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// Setup Sentry
	setupSentry()

	api.TestAuth()

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	//gocsi.Run(context.Background(), service.Name, "A CSI Plugin for VMware vCloud Storage", "", testauth)
}