package profiling

import (
	"go-boilerplate/pkg/logger"
	"log"
	"os"
	"runtime/trace"
)

// StartTrace starts the tracing functionality to profile the program's execution.
func StartTrace() *os.File {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}

	// Start tracing
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

// StopTrace stops tracing.
func StopTrace(f *os.File) {
	trace.Stop()
	if err := f.Close(); err != nil {
		logger.Warn("failed to close trace file: " + err.Error())
	}

}
