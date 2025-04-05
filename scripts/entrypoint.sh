#!/bin/sh

echo "===== Starting Elderwise Application ====="
echo "Running network connectivity tests..."

trap 'echo "Network test script encountered an error, continuing startup anyway"' ERR

if [ -f /app/network_check ]; then
  /app/network_check -internal=true -gmail=true -google=true || echo "Network checks failed, but continuing startup"
else
  echo "Network check utility not found, skipping tests"
fi

trap - ERR

echo "Starting main application..."
exec /app/elderwise
