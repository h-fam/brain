package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/h-fam/brain/base/go/log"
	"github.com/h-fam/brain/windows/display/resolv"
)

func loadResolvConf(fName string) ([]byte, error) {
	b, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	b, err := loadResolvConf("/etc/resolv.conf")
	if err != nil {
		log.Errorf("Failed to load resolv.conf: %v", err)
		os.Exit(1)
	}
	p := resolv.NewParser(bytes.NewBuffer(b))
	r, err := p.Parse()
	if err != nil {
		log.Errorf("failed to parse file: %v", err)
	}
	if len(r.NameServers) == 0 {
		fmt.Println("0.0.0.0")
	}
	fmt.Printf("%s:0", r.NameServers[0])
}
