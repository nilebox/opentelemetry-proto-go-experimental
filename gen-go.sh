#!/usr/bin/env bash

set -e

cp -r ./proto/opentelemetry ./opentelemetry
make gen-go-ci
rm -rf ./opentelemetry
