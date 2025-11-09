#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Running all examples..."
echo

failed=0
total=0

for dir in */; do
    # Skip if not a directory or if it's a hidden directory
    if [[ ! -d "$dir" || "$dir" == ".*" ]]; then
        continue
    fi
    
    # Remove trailing slash
    example_name="${dir%/}"
    
    # Check if main.go exists
    if [[ -f "$dir/main.go" ]]; then
        total=$((total + 1))
        echo "▶ Running example: $example_name"
        if (cd "$dir" && go run main.go); then
            echo "✓ $example_name passed"
        else
            echo "✗ $example_name failed"
            failed=$((failed + 1))
        fi
        echo
    fi
done

echo "================================"
echo "Results: $((total - failed))/$total passed"

if [[ $failed -gt 0 ]]; then
    echo "❌ $failed example(s) failed"
    exit 1
else
    echo "✅ All examples passed"
    exit 0
fi
