#!/bin/bash

# This script is used to run all tests in the project.

set -e

echo "Running unit tests..."
go test ./test/unit/... -v

echo "Running integration tests..."
go test ./test/integration/... -v

echo "Running end-to-end tests..."
go test ./test/e2e/... -v

echo "Running load tests..."
k6 run ./test/load/k6_load_test.js

echo "All tests completed successfully!"