package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Log colors for different levels
var (
	infoColor    = color.New(color.FgGreen)
	warningColor = color.New(color.FgYellow)
	errorColor   = color.New(color.FgRed)
	debugColor   = color.New(color.FgBlue)
)

// CustomFormatter formats logs with colors, timestamps, and file information
type CustomFormatter struct{}

// Format formats the log entry
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Format the timestamp in USA format
	timestamp := entry.Time.Format("01-02-2006 03:04:05 PM")

	// Select color based on log level
	var levelColor *color.Color
	switch entry.Level {
	case logrus.InfoLevel:
		levelColor = infoColor
	case logrus.WarnLevel:
		levelColor = warningColor
	case logrus.ErrorLevel:
		levelColor = errorColor
	case logrus.DebugLevel:
		levelColor = debugColor
	default:
		levelColor = color.New(color.Reset)
	}

	// Format log with color and additional details
	log := levelColor.Sprintf("[%s] %-7s", timestamp, entry.Level.String())
	log += color.New(color.FgWhite).Sprintf(" %s - %s\n", entry.Data["caller"], entry.Message)

	return []byte(log), nil
}

var (
	loggerInstance *logrus.Logger
	once           sync.Once
)

// GetLogger initializes and returns the logger instance
func GetLogger() *logrus.Logger {
	once.Do(func() {
		loggerInstance = logrus.New()
		loggerInstance.SetFormatter(&CustomFormatter{})

		// Define the log file path
		logFilePath := "/var/log/app/app.log"

		// Ensure the directory exists
		logDir := filepath.Dir(logFilePath)
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			if err := os.MkdirAll(logDir, 0755); err != nil {
				fmt.Printf("Failed to create log directory: %v\n", err)
			}
		}

		// Open or create the log file
		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Failed to open log file, using default stdout: %v\n", err)
			loggerInstance.SetOutput(os.Stdout)
		} else {
			multiOutput := io.MultiWriter(os.Stdout, logFile)
			loggerInstance.SetOutput(multiOutput)
		}

		loggerInstance.SetLevel(logrus.DebugLevel)
	})
	return loggerInstance
}

// Log functions for different log levels, retrieving caller info
func logWithCaller(level logrus.Level, msg string) {
	_, file, line, ok := runtime.Caller(2) // Change this to 2 to skip logWithCaller
	if !ok {
		file = "unknown"
		line = 0
	}
	callerInfo := fmt.Sprintf("%s:%d", filepath.Base(file), line)

	logger := GetLogger()
	logger.WithField("caller", callerInfo).Log(level, msg)
}

func Info(msg string) {
	logWithCaller(logrus.InfoLevel, msg)
}

func Warn(msg string) {
	logWithCaller(logrus.WarnLevel, msg)
}

func Error(msg string) {
	logWithCaller(logrus.ErrorLevel, msg)
}

func Debug(msg string) {
	logWithCaller(logrus.DebugLevel, msg)
}
