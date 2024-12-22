#!/usr/bin/env bash

set -euo pipefail

REPO_DIR="$(git rev-parse --show-toplevel)"
readonly REPO_DIR

jar_file=antlr-4.13.2-complete.jar
jar_path="${REPO_DIR}/tools/${jar_file}"
readonly jar_path jar_file

if [ ! -f "${jar_path}" ]; then
  curl "https://www.antlr.org/download/${jar_file}" > "${jar_path}"
fi

docker run                                                                     \
  -v "${REPO_DIR}:/workspace"                                                  \
  -w /workspace/data                                                           \
  --entrypoint java                                                            \
  openjdk:11 -jar "../tools/${jar_file}"                                       \
    -Dlanguage=Go                                                              \
    -o ../internal/parser                                                      \
    -visitor                                                                   \
    -no-listener                                                               \
    yamlpath.g4

