#!/bin/sh

set -e

script_dir=$(cd $(dirname $0) && pwd)

godep_gopath=${script_dir}/Godeps/_workspace
export GOPATH=${godep_gopath}:$GOPATH

cd ${script_dir}/elasticsearch_service
go test
