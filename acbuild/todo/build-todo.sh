#!/usr/bin/env bash
set -e

if [ "$EUID" -ne 0 ]; then
    echo "This script uses functionality which requires root privileges"
    exit 1
fi

LATEST_VERSION=$(curl https://api.github.com/repos/mix3/go-todo-webapp/releases | python -c 'import sys; import json; print(json.loads(sys.stdin.read()))[0]["tag_name"]');
curl -L -o todo.zip https://github.com/mix3/go-todo-webapp/releases/download/${LATEST_VERSION}/todo.zip
unzip todo.zip

acbuild --debug begin

trap "{ export EXT=$?; acbuild --debug end && rm -rf todo/ todo.zip  exit $EXT; }" EXIT

acbuild --debug set-name mix3.github.io/aci/todo

acbuild --debug copy todo/static /static
acbuild --debug copy todo/todo /bin/todo
acbuild --debug copy run.sh /bin/run.sh

acbuild --debug label add version latest
acbuild --debug label add arch amd64
acbuild --debug label add os linux

acbuild --debug dep add quay.io/coreos/alpine-sh

acbuild --debug run -- apk update
acbuild --debug run -- apk add bash mysql-client

acbuild --debug set-exec -- /bin/run.sh

acbuild --debug write --overwrite todo-latest-linux-amd64.aci
