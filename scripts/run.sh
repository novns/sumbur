#! /bin/sh

set -e

cd "$(dirname "$0")/../public"

[ -x "$GOPATH/bin/gin" ]  ||  go install github.com/codegangsta/gin@latest
[ -x "$GOPATH/bin/qtc" ]  ||  go install github.com/valyala/quicktemplate/qtc@latest

export SUMBUR_CONFIG=sumbur-debug.yaml
gin  -a 8000  -b sumbur-debug  -t ..  run
