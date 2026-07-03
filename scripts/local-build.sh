#!/bin/bash

OLD_PWD=$(pwd)
SHELL_FOLDER=$(
    cd "$(dirname "$0")" || exit
    pwd
)
PROJECT_FOLDER=$SHELL_FOLDER/..

cd "$SHELL_FOLDER" || exit >/dev/null 2>&1

process() {
    COMMAND="\"$PROJECT_FOLDER/scripts/build/clean-build-cache.sh\""
    echo exec: "$COMMAND"
    if ! eval "$COMMAND"; then
        echo ""
        echo ""
        echo "[ERROR]: failed to clean build cache"
        echo ""
        echo ""

        exit 1
    fi

    COMMAND="\"$PROJECT_FOLDER/scripts/build/build.sh\""
    echo exec: "$COMMAND"
    if ! eval "$COMMAND"; then
        echo ""
        echo ""
        echo "[ERROR]: failed to build export"
        echo ""
        echo ""

        exit 1
    fi

    if [ "$IS_SUPPORT_COMPRESS_RELEASE" = "true" ]; then
        COMMAND="\"$PROJECT_FOLDER/scripts/build/compress-release.sh\""
        echo exec: "$COMMAND"
        if ! eval "$COMMAND"; then
            echo ""
            echo ""
            echo "[ERROR]: failed to compress release"
            echo ""
            echo ""

            exit 1
        fi
    fi
}

usage() {
    echo "Usage:"
    echo "  $(basename "$0") [--flag-compress-release] [-h]"
    echo ""
    echo "Description:"
    echo "  --flag-compress-release: whether to compress the release binaries, default is false"
    echo ""
    echo "Example:"
    echo "  $(basename "$0") --flag-compress-release"
    echo ""
    echo "Note: publishing to GitHub Releases is handled by .github/workflows/release.yml"
    echo "on tag push, not by this script."

    exit 1
}

IS_SUPPORT_COMPRESS_RELEASE="false"
while true; do
    if [ -z "$1" ]; then
        break
    fi

    case "$1" in
    -h | --h | h | -help | --help | help | -H | --H | HELP)
        usage
        ;;
    --flag-compress-release)
        IS_SUPPORT_COMPRESS_RELEASE="true"
        shift 1
        ;;
    *)
        echo ""
        echo "unknown option: $1"
        echo ""

        usage
        ;;
    esac
done

process

cd "$OLD_PWD" || exit >/dev/null 2>&1
