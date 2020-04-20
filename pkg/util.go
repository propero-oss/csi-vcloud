package pkg

import (
	"fmt"
	"log"
	"os"
	)


func GetHostname() (string, error){
	name, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("could not retrieve Hostname: %s", err)
	}

	return name, nil
}