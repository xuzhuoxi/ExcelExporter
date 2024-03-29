#!/usr/bin/env bash

CRTDIR=$(cd `dirname $0`; pwd)

goxc -os="linux darwin windows freebsd openbsd" -arch="amd64 arm" -n=ExcelExporter -pv=v2.2 -wd=${CRTDIR}/../src -d=${CRTDIR}/./release -include=*.go