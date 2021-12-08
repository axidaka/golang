#!/bin/bash
set -ue
cd `dirname $0`
protoc -I ./ --go_out=plugins=grpc:./go/ *.proto
