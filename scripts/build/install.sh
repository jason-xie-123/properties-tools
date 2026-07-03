#!/bin/bash

OLD_PWD=$(pwd)
SHELL_FOLDER=$(
    cd "$(dirname "$0")" || exit
    pwd
)
PROJECT_FOLDER=$SHELL_FOLDER/../..

cd "$PROJECT_FOLDER/properties-cli" || exit >/dev/null 2>&1

go install ./cmd/properties-cli

cd "$OLD_PWD" || exit >/dev/null 2>&1
