#!/usr/bin/env bash
set -e

if [ "$EUID" -ne 0 ]; then
    echo "This script uses functionality which requires root privileges"
    exit 1
fi

acbuild --debug begin

trap "{ export EXT=$?; acbuild --debug end && exit $EXT; }" EXIT

acbuild --debug set-name mix3.github.io/aci/mysqld

acbuild --debug copy run.sh /bin/run.sh

acbuild --debug label add version latest
acbuild --debug label add arch amd64
acbuild --debug label add os linux

acbuild --debug dep add quay.io/coreos/alpine-sh

acbuild --debug run -- apk update
acbuild --debug run -- apk add mysql mysql-client bash

acbuild --debug set-exec -- /bin/run.sh

acbuild --debug write --overwrite mysqld-latest-linux-amd64.aci
