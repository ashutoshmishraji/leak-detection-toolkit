#!/bin/bash

# LeakFinder - Production-Optimized Leak Detection Toolkit
# Generic toolkit for any Go application with pprof enabled

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PPROF_URL="http://localhost:6060"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m'

log() { echo -e "${BLUE}$1${NC}"; }
success() { echo -e "${GREEN}✅ $1${NC}"; }
warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }
error() { echo -e "${RED}❌ $1${NC}"; }

# Check if target application is running with pprof
check_pprof() {
    if curl -s "$PPROF_URL/debug/pprof/" > /dev/null 2>&1; then
        return 0
    else
        return 1
    fi
}

# Show help
show_help() {
    echo -e "${BLUE}🎯 Optimized Leak Detection Toolkit${NC}"
    echo ""
    echo "Usage: $0 [options] <command>"
    echo ""
    echo "Commands:"
    echo "  check     - Quick goroutine and memory check"
    echo "  monitor   - Real-time monitoring with configurable duration"
    echo "  analyze   - Detailed analysis with leak detection"
    echo "  compare   - Before/after comparison test"
    echo "  filter    - Filtered analysis (use env vars for filters)"
    echo "  patterns  - Stack pattern analysis (advanced leak detection)"
    echo "  channels  - Channel leak detection (production-grade)"
    echo "  help      - Show this help"
    echo ""
    echo "Common Options:"
    echo "  -url=URL                    pprof server URL (default: http://localhost:6060)"
    echo "  -verbose                    enable verbose output"
    echo "  -timeout=DURATION           HTTP request timeout (default: 30s)"
    echo "  -monitor-duration=DURATION  monitoring duration (default: 5m)"
    echo "  -monitor-interval=DURATION  monitoring interval (default: 3s)"
    echo "  -goroutine-critical=N       critical goroutine threshold (default: 500)"
    echo "  -memory-critical=N          critical memory threshold MB (default: 500)"
    echo ""
        echo "Examples:"
        echo "  $0 check"
        echo "  $0 -verbose monitor"
        echo "  $0 -url=http://localhost:8080 analyze"
        echo "  $0 -goroutine-critical=1000 -monitor-duration=10m monitor"
        echo "  $0 -timeout=60s patterns"
        echo ""
        echo "Focus Options (available for check, monitor, analyze, compare):"
        echo "  $0 check                           # Both goroutines and memory (default)"
        echo "  $0 check -goroutines-only          # Only goroutine analysis"
        echo "  $0 check -memory-only              # Only memory analysis"
        echo "  $0 monitor -goroutines-only        # Monitor only goroutines"
        echo "  $0 analyze -memory-only            # Analyze only memory"
        echo "  $0 compare -goroutines-only        # Compare only goroutines"
        echo ""
        echo "GC Analysis Examples (add -force-gc to any command):"
        echo "  $0 check -force-gc                 # Quick check with GC analysis"
        echo "  $0 analyze -force-gc -memory-only  # Analyze memory with GC test"
        echo "  $0 compare -force-gc               # Compare with GC analysis"
        echo "  $0 -gc-cycles=3 -gc-wait=60s check -force-gc  # Custom GC parameters"
        echo "  $0 -gc-threshold=70 analyze -force-gc         # Custom leak threshold"
        echo ""
        echo "Pattern Examples:"
        echo "  $0 -suspicious-patterns=\"worker.stuck,processor.timeout\" analyze"
        echo "  $0 -safe-patterns=\"app.healthCheck,app.monitor\" monitor"
        echo "  SUSPICIOUS_PATTERNS=\"custom.pattern\" $0 analyze"
        echo "  SAFE_PATTERNS=\"app.background\" $0 monitor"
    echo ""
    echo "Filter Examples:"
    echo "  FILTER_FUNC=myFunction $0 filter"
    echo "  FILTER_PKG=github.com/myapp $0 filter"
    echo "  EXCLUDE_PATTERNS=time.Sleep,chan $0 filter"
    echo ""
    echo "Prerequisites:"
    echo "  • Go application running with pprof enabled"
    echo "  • pprof endpoint accessible (default: http://localhost:6060)"
    echo ""
    echo "Features:"
    echo "  ✅ Configurable thresholds and timeouts"
    echo "  ✅ Advanced command-line options"
    echo "  ✅ Production-ready leak detection algorithms"
    echo "  ✅ Memory pattern analysis (Datadog-inspired)"
    echo "  ✅ Channel leak detection (Uber LeakProf)"
    echo "  ✅ Comprehensive reporting with file output"
    echo ""
}

# Parse command line arguments for URL extraction
parse_url_from_args() {
    for arg in "$@"; do
        if [[ "$arg" == -url=* ]]; then
            PPROF_URL="${arg#-url=}"
            break
        fi
    done
}

# Main execution
main() {
    cd "$SCRIPT_DIR"
    
    # Parse URL from arguments if provided
    parse_url_from_args "$@"
    
    # Get the command (last non-flag argument)
    COMMAND=""
    for arg in "$@"; do
        if [[ "$arg" != -* ]]; then
            COMMAND="$arg"
        fi
    done
    
    # Default to help if no command provided
    if [[ -z "$COMMAND" ]]; then
        COMMAND="help"
    fi
    
    # Check if target application is running (except for help)
    if [[ "$COMMAND" != "help" ]]; then
        if ! check_pprof; then
            error "No pprof endpoint found at $PPROF_URL"
            echo ""
            echo "💡 Make sure your Go application is running with pprof enabled:"
            echo "   import _ \"net/http/pprof\""
            echo "   go http.ListenAndServe(\":6060\", nil)"
            echo ""
            echo "Or specify a different URL:"
            echo "   $0 -url=http://your-app:port $COMMAND"
            exit 1
        fi
        success "Connected to pprof at $PPROF_URL"
    fi
    
    case "$COMMAND" in
        "check")
            log "Running optimized quick check..."
            go run leakfinder.go "$@"
            ;;
        "monitor")
            log "Starting optimized real-time monitor..."
            go run leakfinder.go "$@"
            ;;
        "analyze")
            log "Running enhanced detailed analysis..."
            go run leakfinder.go "$@"
            ;;
        "compare")
            log "Running before/after comparison..."
            go run leakfinder.go "$@"
            ;;
        "filter")
            log "Running filtered analysis..."
            if [[ -z "$FILTER_FUNC" && -z "$FILTER_PKG" ]]; then
                warning "No filters specified. Set FILTER_FUNC or FILTER_PKG environment variables."
                echo "Example: FILTER_FUNC=myFunction $0 filter"
            fi
            go run leakfinder.go "$@"
            ;;
        "patterns")
            log "Running stack pattern analysis..."
            go run leakfinder.go "$@"
            ;;
        "channels")
            log "Running channel leak detection..."
            go run leakfinder.go "$@"
            ;;
        "help"|"")
            show_help
            ;;
        *)
            error "Unknown command: $COMMAND"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

main "$@"
