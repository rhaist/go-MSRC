#!/bin/bash

set -e

export GO_POST_PROCESS_FILE="$(which gofmt) -w"

# Check for openapi-generator
# Install for Mac: brew install openapi-generator
if ! [ -x "$(command -v openapi-generator)" ]; then
  echo 'Error: openapi-generator is not installed.' >&2
  exit 1
fi

openapi-generator generate -g go --enable-post-process-file -o openapi -i https://portal.msrc.microsoft.com/developer/open-api-spec/v1
