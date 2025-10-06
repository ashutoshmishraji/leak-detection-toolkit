package main

import (
    "strings"
    "testing"
    "time"
)

func TestDefaultConfig(t *testing.T) {
    config := DefaultConfig()

    if config == nil {
        t.Fatal("DefaultConfig() returned nil")
    }

    if config.PprofURL != "http://localhost:6060" {
        t.Errorf("Expected default pprof URL to be http://localhost:6060, got %s", config.PprofURL)
    }

    if config.Timeout != 30*time.Second {
        t.Errorf("Expected default timeout to be 30s, got %v", config.Timeout)
    }
}

func TestNewLeakDetectionToolkit(t *testing.T) {
    config := DefaultConfig()
    toolkit := NewLeakDetectionToolkit(config)

    if toolkit == nil {
        t.Fatal("NewLeakDetectionToolkit() returned nil")
    }

    if toolkit.config != config {
        t.Error("Toolkit config not set correctly")
    }

    if toolkit.client == nil {
        t.Error("HTTP client not initialized")
    }
}

func TestParseGoroutines(t *testing.T) {
    toolkit := NewLeakDetectionToolkit(DefaultConfig())

    testCases := []struct {
        name     string
        input    string
        expected int
    }{
        {
            name:     "empty input",
            input:    "",
            expected: 0,
        },
        {
            name: "single goroutine",
            input: `goroutine 1 [running]:
main.main()
    /path/to/main.go:10 +0x20`,
            expected: 1,
        },
        {
            name: "multiple goroutines",
            input: `goroutine 1 [running]:
main.main()
    /path/to/main.go:10 +0x20

goroutine 2 [sleep]:
time.Sleep()
    /usr/local/go/src/runtime/time.go:195 +0x135`,
            expected: 2,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            goroutines := toolkit.parseGoroutines(tc.input)
            if len(goroutines) != tc.expected {
                t.Errorf("Expected %d goroutines, got %d", tc.expected, len(goroutines))
            }
        })
    }
}

func TestInternString(t *testing.T) {
    toolkit := NewLeakDetectionToolkit(DefaultConfig())

    // Test string interning
    str1 := toolkit.internString("test")
    str2 := toolkit.internString("test")

    // Should return the same string value
    if str1 != str2 {
        t.Error("String interning not working correctly - values should be equal")
    }

    // Test different strings
    str3 := toolkit.internString("different")
    if str1 == str3 {
        t.Error("Different strings should not be equal")
    }

    // Test that cache is working
    if len(toolkit.stringCache) == 0 {
        t.Error("String cache should not be empty")
    }
}

func TestParsePatterns(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected []string
    }{
        {
            name:     "empty string",
            input:    "",
            expected: nil,
        },
        {
            name:     "single pattern",
            input:    "pattern1",
            expected: []string{"pattern1"},
        },
        {
            name:     "multiple patterns",
            input:    "pattern1,pattern2,pattern3",
            expected: []string{"pattern1", "pattern2", "pattern3"},
        },
        {
            name:     "patterns with spaces",
            input:    " pattern1 , pattern2 , pattern3 ",
            expected: []string{"pattern1", "pattern2", "pattern3"},
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := parsePatterns(tc.input)

            if len(result) != len(tc.expected) {
                t.Errorf("Expected %d patterns, got %d", len(tc.expected), len(result))
                return
            }

            for i, expected := range tc.expected {
                if result[i] != expected {
                    t.Errorf("Expected pattern %d to be %s, got %s", i, expected, result[i])
                }
            }
        })
    }
}

func TestValidateConfig(t *testing.T) {
    // Test valid config
    validConfig := DefaultConfig()
    if err := validateConfig(validConfig); err != nil {
        t.Errorf("Valid config should not return error: %v", err)
    }

    // Test invalid goroutine thresholds
    invalidConfig := DefaultConfig()
    invalidConfig.GoroutineThresholds.Normal = 100
    invalidConfig.GoroutineThresholds.Moderate = 50 // Lower than normal

    if err := validateConfig(invalidConfig); err == nil {
        t.Error("Invalid goroutine thresholds should return error")
    }

    // Test invalid memory thresholds
    invalidConfig2 := DefaultConfig()
    invalidConfig2.MemoryThresholds.High = 50
    invalidConfig2.MemoryThresholds.Critical = 25 // Lower than high

    if err := validateConfig(invalidConfig2); err == nil {
        t.Error("Invalid memory thresholds should return error")
    }
}

func TestGoroutineRegex(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected bool
    }{
        {
            name:     "valid goroutine line",
            input:    "goroutine 1 [running]:",
            expected: true,
        },
        {
            name:     "valid goroutine with different state",
            input:    "goroutine 123 [sleep]:",
            expected: true,
        },
        {
            name:     "invalid line",
            input:    "not a goroutine line",
            expected: false,
        },
        {
            name:     "empty line",
            input:    "",
            expected: false,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            matches := goroutineRegex.MatchString(tc.input)
            if matches != tc.expected {
                t.Errorf("Expected %v for input %q, got %v", tc.expected, tc.input, matches)
            }
        })
    }
}

func TestFunctionRegex(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected bool
    }{
        {
            name:     "valid function call",
            input:    "main.main()",
            expected: true,
        },
        {
            name:     "function with parameters",
            input:    "fmt.Printf(0x1234, 0x5678)",
            expected: true,
        },
        {
            name:     "invalid line",
            input:    " /path/to/file.go:123 +0x456",
            expected: false,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            matches := functionRegex.MatchString(tc.input)
            if matches != tc.expected {
                t.Errorf("Expected %v for input %q, got %v", tc.expected, tc.input, matches)
            }
        })
    }
}

// Benchmark tests
func BenchmarkParseGoroutines(b *testing.B) {
    toolkit := NewLeakDetectionToolkit(DefaultConfig())

    // Create a sample goroutine dump
    var sb strings.Builder
    for i := 0; i < 100; i++ {
        sb.WriteString("goroutine ")
        sb.WriteString(string(rune(i + 1)))
        sb.WriteString(" [running]:\n")
        sb.WriteString("main.worker()\n")
        sb.WriteString("    /path/to/main.go:123 +0x456\n\n")
    }
    dump := sb.String()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        toolkit.parseGoroutines(dump)
    }
}

func BenchmarkInternString(b *testing.B) {
    toolkit := NewLeakDetectionToolkit(DefaultConfig())

    strings := []string{
        "running", "sleep", "chan receive", "chan send",
        "main.main", "runtime.gopark", "time.Sleep",
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, s := range strings {
            toolkit.internString(s)
        }
    }
}