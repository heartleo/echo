#!/bin/bash

set -ex

CUR_DIR=$(cd "$(dirname "$0")";pwd)
ROOT_DIR=${CUR_DIR}/../

docker build --no-cache --rm --progress plain \
--file ${ROOT_DIR}/Dockerfile \
--tag heartleo/echo:latest \
${ROOT_DIR}