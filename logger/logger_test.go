package logger

import (
	"testing"
)

func TestWarn(t *testing.T) {
	LWarn(true, "log test...")
}

func TestWarnNew(t *testing.T) {
	LWarn_("log test...")
}

func BenchmarkWarnFilenameLineno(b *testing.B) {
    for i := 0; i < b.N; i++ {
        LWarn(true, "benchmark log test...")
    }
}

func BenchmarkWarnNoFilenameLineno(b *testing.B) {
    for i := 0; i < b.N; i++ {
        LWarn(false, "benchmark log test...")
    }
}
