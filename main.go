package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"github.com/paysuper/paysuper-taxjar-rate-importer/pkg"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
)

// Config define application config object
type Config struct {
	KubernetesHost string `envconfig:"KUBERNETES_SERVICE_HOST" required:"false"`
	TaxJarToken    string `envconfig:"TAX_JAR_TOKEN" required:"true"`
	ZipCodeFile    string `envconfig:"ZIP_CODE_FILE" required:"false"`
	CachePath      string `envconfig:"CACHE_PATH" required:"false" default:"./cache"`
	MaxRPS         int    `envconfig:"MAX_RPS" required:"false" default:"250"`
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
		logger.Fatal("Config init failed with error", zap.Error(err))
	}

	resty.SetAuthToken(config.TaxJarToken)

	db, err := leveldb.OpenFile(config.CachePath, nil)
	if err != nil {
		logger.Fatal("Failed to load cache db", zap.Error(err))
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("Failed to properly close cache db", zap.Error(err))
		}
	}()

	var clientService micro.Service
	if config.KubernetesHost == "" {
		clientService = micro.NewService()
		logger.Info("Initialize micro service")
	} else {
		clientService = k8s.NewService()
		logger.Info("Initialize k8s service")
	}

	clientService.Init()

	taxService := taxjar.NewClient(db, clientService, config.MaxRPS)
	if err := taxService.Run(config.ZipCodeFile); err != nil {
		logger.Fatal("Update rates failed", zap.Error(err))
	}

	logger.Info("Update complete")
}
