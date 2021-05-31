package main

type ICfgParse interface {
	parseJson(proxy IJsonReaderProxy)
	parseBinary(proxy IBinaryReaderProxy)
}
