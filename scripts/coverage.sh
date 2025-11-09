#!/bin/bash

CACHE_DIR=".cache"
if [ ! -d ${CACHE_DIR} ]; then
    mkdir -p ${CACHE_DIR}
fi

COVERAGE_PROFILE=${CACHE_DIR}/coverage.out
LCOV_FILE=lcov.info
go test -coverprofile=${COVERAGE_PROFILE} ./...
go run github.com/jandelgado/gcov2lcov -infile=${COVERAGE_PROFILE} -outfile=${LCOV_FILE}
