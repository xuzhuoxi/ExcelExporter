module github.com/xuzhuoxi/ExcelExporter

go 1.16

require (
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/richardlehane/mscfb v1.0.4
	github.com/spf13/viper v1.10.1
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/sjson v1.2.5
	github.com/xuri/efp v0.0.0-20230422071738-01f4e37c47e9
	github.com/xuzhuoxi/infra-go v1.0.4
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5
	golang.org/x/text v0.3.7
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/xuzhuoxi/infra-go => ../infra-go
)
