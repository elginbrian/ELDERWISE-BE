#!/bin/sh
set -e

echo "===== Starting Elderwise Application ====="
echo "Running network connectivity tests..."

go build -o /tmp/network_check /app/scripts/network_check.go
/tmp/network_check -internal=true -gmail=true -google=true

TEST_EXIT_CODE=$?
if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo "WARNING: Network tests failed with exit code $TEST_EXIT_CODE"
  echo "The application will continue to start, but there might be connectivity issues."
else
  echo "Network tests completed successfully!"
fi

echo "Starting main application..."
exec /app/elderwise
