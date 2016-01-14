#!/bin/bash
set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
OUTPUT_PATH=$DIR/bosh-agent.exe

rm -f $OUTPUT_PATH $CONFIG_PATH $SETTINGS_PATH $SERVICE_CONFIG

GOOS=windows \
  go build \
  -o \
  $OUTPUT_PATH \
  github.com/cloudfoundry/bosh-agent/main

if vagrant status | grep agent | grep running
then
  vagrant provision
else
  vagrant up --provider=aws
fi
