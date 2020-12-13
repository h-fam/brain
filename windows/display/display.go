package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/h-fam/brain/base/go/log"
)

func loadResolvConf(fName string) (string, error) {
	b, err := ioutil.ReadFile(fName)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	s, err := loadResolvConf("/etc/resolv.conf")
	if err != nil {
		log.Errorf("failed to load resolv.conf: %v", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
