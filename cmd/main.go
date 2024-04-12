package main

import (
	"context"
	"time"

	_ "go.uber.org/automaxprocs"

	"github.com/nikitads9/banner-service-api/internal/pkg/banner"

	"flag"
	"log"
)

var configType, pathConfig string

func init() {
	flag.StringVar(&configType, "configtype", "file", "type of configuration: environment variables (env) or env/yaml file (file)")
	flag.StringVar(&pathConfig, "config", "C:\\Users\\swnik\\Desktop\\projects\\banner-service-api\\configs\\banners_config.yml", "path to config file")
	time.Local = time.UTC
}

func main() {
	flag.Parse()

	ctx := context.Background()

	app, err := banner.NewApp(ctx, configType, pathConfig)
	if err != nil {
		log.Fatalf("failed to create banners-api app object:%s\n", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run banners-api app: %s", err.Error())
	}
}
