#!/bin/bash

set -e

export GO_POST_PROCESS_FILE="$(which gofmt) -w"
export GIT_USER_ID="rhaist"
export GIT_REPO_ID="go-MSRC"

# Check for openapi-generator-cli
if ! [ -x "$(command -v openapi-generator-cli)" ]; then
	echo 'Error: openapi-generator-cli is not installed.' >&2
	exit 1
fi

rm -rf openapi/*

openapi-generator-cli generate \
	--git-user-id $GIT_USER_ID \
	--git-repo-id $GIT_REPO_ID \
	--enable-post-process-file \
	--additional-properties=isGoSubmodule=true \
	--skip-validate-spec \
	-g go \
	-o openapi \
	-i https://api.msrc.microsoft.com/cvrf/v3.0/swagger/v3/swagger.json
