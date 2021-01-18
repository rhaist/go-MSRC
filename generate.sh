#!/bin/bash

set -e

export GO_POST_PROCESS_FILE="$(which gofmt) -w"
export GIT_USER_ID="rhaist"
export GIT_REPO_ID="go-MSRC/openapi"

# Check for openapi-generator
# Install for Mac: brew install openapi-generator
if ! [ -x "$(command -v openapi-generator)" ]; then
  echo 'Error: openapi-generator is not installed.' >&2
  exit 1
fi

openapi-generator generate \
  --git-user-id $GIT_USER_ID \
  --git-repo-id $GIT_REPO_ID \
  -g go \
  --enable-post-process-file \
  -o openapi \
  -i https://portal.msrc.microsoft.com/developer/open-api-spec/v1
