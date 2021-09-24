#! /bin/sh

set -e

cd "$(dirname "$0")/../public"

[ -x "$GOPATH/bin/gin" ]  ||  go install github.com/codegangsta/gin@latest

gin  -a 8000  -b sumbur-debug  -t ..  run
