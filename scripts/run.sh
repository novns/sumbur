#! /bin/sh

set -e


[ -x "$GOPATH/bin/gin" ]  ||  go install github.com/codegangsta/gin@latest
[ -x "$GOPATH/bin/qtc" ]  ||  go install github.com/valyala/quicktemplate/qtc@latest


cd "$(dirname "$0")/../views"

qtc -skipLineComments


cd ../public

export SUMBUR_CONFIG=sumbur-debug.yaml
gin  -a 8000  -b sumbur-debug  -t ..  run
