package main

import (
	"flag"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/propero-oss/csi-vcloud/pkg/api"
	mycommon "github.com/propero-oss/csi-vcloud/pkg/common"
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

	test()

	// Setup Sentry
	//setupSentry()

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	//defer sentry.Flush(2 * time.Second)

	//gocsi.Run(context.Background(), service.Name, "A CSI Plugin for VMware vCloud Storage", "", provider.New())
}

func timeit(a func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}


func test() {
	var config = mycommon.Config{VCloud: struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		ORG      string `yaml:"org"`
		API      string `yaml:"api"`
		VDC      string `yaml:"vdc"`
		Insecure bool `yaml:"insecure"`
	}{Username: os.Getenv("USERNAME"), Password: os.Getenv("PASSWORD"), ORG: os.Getenv("ORG"), API: os.Getenv("API") , VDC: os.Getenv("VDC") , Insecure: false}}
	
	manager := api.Manager{
		Client: nil,
		Config: nil,
		Vdc:    nil,
	}


	client, err := api.Client(&config)
	if err != nil {
		panic(err)
	}
	manager.Client = client
	manager.Config = &config

	vdc, err := manager.GetVDC(config.VCloud.ORG, config.VCloud.VDC)
	if err != nil {
		panic(err)
	}

	manager.Vdc = vdc


	vm, err := manager.GetVMByVAppName("worker-nodes-nschad", "worker-node-0")

	data, err := api.FindNextBusAndUnitNumber(vm)

	fmt.Println(data)


}