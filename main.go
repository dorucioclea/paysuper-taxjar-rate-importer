package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/paysuper/paysuper-taxjar-rate-importer/pkg"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
)

// Config define application config object
type Config struct {
	TaxJarToken string `envconfig:"TAX_JAR_TOKEN" required:"true"`
	ZipCodeFile string `envconfig:"ZIP_CODE_FILE" required:"false"`
	MaxRPS      int    `envconfig:"MAX_RPS" required:"false" default:"10"`
}

func init() {
	resty.SetHostURL("https://api.taxjar.com/v2/rates/")
	resty.SetHeader("Accept", "application/json")
}

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		logger.Fatal("Config init failed with error: %s\n", zap.Error(err))
	}

	resty.SetAuthToken(config.TaxJarToken)

	client := taxjar.NewClient(config.MaxRPS)
	if err := client.Run(config.ZipCodeFile); err != nil {
		logger.Fatal("Update rates failed", zap.Error(err))
	}

	logger.Info("Update complete")
}
