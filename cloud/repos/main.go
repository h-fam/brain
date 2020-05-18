package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sourcerepo/v1"
)

const (
	projectPath = "projects/hines-alloc"
)

func main() {
	ctx := context.Background()
	s, err := sourcerepo.NewService(ctx)
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}
	cfg, err := s.Projects.GetConfig("projects/hines-alloc").Do()
	cfgs := cfg.PubsubConfigs
	if cfgs == nil {
		cfgs = map[string]sourcerepo.PubsubConfig{}
	}
	cfgs["build-brain"] = sourcerepo.PubsubConfig{
		Topic:         "build-brain",
		MessageFormat: "PROTOBUF",
	}
	cfg.PubsubConfigs = cfgs
	fmt.Println(cfg, err)
	s.Projects.UpdateConfig("projects/hines-alloc")
}
