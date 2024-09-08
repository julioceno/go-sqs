#!/usr/bin/env bash

set -euo pipefail

# enable debug
# set -x

echo "configuring sqs"
echo "==================="

awslocal sqs create-queue --queue-name queue1

echo "finished"