#!/bin/bash

# LeakFinder Basic Usage Examples
# This script demonstrates common LeakFinder usage patterns

set -e

echo "🎯 LeakFinder Basic Usage Examples"
echo "=================================="

# Ensure leakfinder is available
if ! command -v ./leakfinder &> /dev/null; then
    echo "Building leakfinder..."
    go build -o leakfinder leakfinder.go
fi

echo ""
echo "1. 🔍 Quick Health Check"
echo "------------------------"
./leakfinder check

echo ""
echo "2. 💾 Memory-Only Analysis"
echo "--------------------------"
./leakfinder check -memory-only

echo ""
echo "3. 🧵 Goroutines-Only Analysis"
echo "------------------------------"
./leakfinder check -goroutines-only

echo ""
echo "4. 🧠 GC Analysis (Distinguish real leaks from GC delays)"
echo "--------------------------------------------------------"
./leakfinder -force-gc -gc-cycles=1 -gc-wait=5s check -memory-only

echo ""
echo "5. 📊 Detailed Analysis with Verbose Output"
echo "-------------------------------------------"
timeout 10s ./leakfinder -verbose analyze -memory-only || echo "Analysis completed"

echo ""
echo "6. 🎯 Pattern-Based Leak Detection"
echo "----------------------------------"
timeout 5s ./leakfinder patterns || echo "Pattern analysis completed"

echo ""
echo "7. 🔧 Custom Configuration Example"
echo "----------------------------------"
./leakfinder -goroutine-critical=200 -memory-critical=100 check

echo ""
echo "✅ All examples completed successfully!"
echo ""
echo "💡 Tips:"
echo "  - Use -verbose for detailed output"
echo "  - Use -force-gc to verify real leaks"
echo "  - Use -goroutines-only or -memory-only for focused analysis"
echo "  - Check ./reports/ directory for saved analysis files"
