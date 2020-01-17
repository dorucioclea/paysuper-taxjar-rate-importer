module github.com/paysuper/paysuper-taxjar-rate-importer

require (
	dmitri.shuralyov.com/app/changes v0.0.0-20181114035150-5af16e21babb // indirect
	dmitri.shuralyov.com/service/change v0.0.0-20190203025214-430bf650e55a // indirect
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/micro/go-micro v1.8.0
	github.com/micro/go-plugins v1.2.0
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paysuper/paysuper-tax-service v1.0.0
	github.com/russross/blackfriday v2.0.0+incompatible // indirect
	github.com/shurcooL/go v0.0.0-20190121191506-3fef8c783dec // indirect
	github.com/shurcooL/gofontwoff v0.0.0-20181114050219-180f79e6909d // indirect
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20181215221002-9d8641ddf2e1 // indirect
	github.com/shurcooL/htmlg v0.0.0-20190120222857-1e8a37b806f3 // indirect
	github.com/shurcooL/httpfs v0.0.0-20181222201310-74dc9339e414 // indirect
	github.com/shurcooL/issues v0.0.0-20190120000219-08d8dadf8acb // indirect
	github.com/shurcooL/issuesapp v0.0.0-20181229001453-b8198a402c58 // indirect
	github.com/shurcooL/notifications v0.0.0-20181111060504-bcc2b3082a7a // indirect
	github.com/shurcooL/octicon v0.0.0-20181222203144-9ff1a4cf27f4 // indirect
	github.com/shurcooL/reactions v0.0.0-20181222204718-145cd5e7f3d1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/shurcooL/webdavfs v0.0.0-20181215192745-5988b2d638f6 // indirect
	github.com/syndtr/goleveldb v1.0.0
	go.uber.org/ratelimit v0.1.0
	go.uber.org/zap v1.10.0
	go4.org v0.0.0-20181109185143-00e24f1b2599 // indirect
	golang.org/x/perf v0.0.0-20190124201629-844a5f5b46f4 // indirect
	gopkg.in/resty.v1 v1.12.0
	sourcegraph.com/sqs/pbtypes v1.0.0 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
)

go 1.13
