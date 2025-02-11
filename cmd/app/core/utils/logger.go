package utils

import (
	"context"
	"log"
	"os"
)

const RequestIDKey = "RequestID"

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
	WarnLogger  *log.Logger
)

func InitLogger() {

	// Open or create log files
	infoFile, err := os.OpenFile("logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open info log file:", err)
	}
	errorFile, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open error log file:", err)
	}
	debugFile, err := os.OpenFile("logs/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open debug log file:", err)
	}
	warnFile, err := os.OpenFile("logs/warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open warn log file:", err)
	}

	// Initialize loggers
	InfoLogger = log.New(infoFile, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorFile, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(debugFile, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(warnFile, "WARN ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs informational messages with requestID if available
func LogInfo(message string, ctx context.Context) {
	logWithRequestID(InfoLogger, "INFO", message, ctx)
}

// LogError logs error messages with requestID if available
func LogError(message string, ctx context.Context) {
	logWithRequestID(ErrorLogger, "ERROR", message, ctx)
}

// LogDebug logs debug messages with requestID if available
func LogDebug(message string, ctx context.Context) {
	logWithRequestID(DebugLogger, "DEBUG", message, ctx)
}

// LogWarn logs warning messages with requestID if available
func LogWarn(message string, ctx context.Context) {
	logWithRequestID(WarnLogger, "WARN", message, ctx)
}

// logWithRequestID logs the message with a specified logger, log level, and request ID if available
func logWithRequestID(logger *log.Logger, level, message string, ctx context.Context) {
	requestID, ok := ctx.Value(RequestIDKey).(string)
	if ok {
		logger.Printf("[%s] [RequestID: %s] %s", level, requestID, message)
	} else {
		logger.Printf("[%s] %s", level, message)
	}
}
