package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Log         *log.Logger
	TraceLogger *log.Logger
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
)

func Init(test bool, prefix string) {
	var test_prefix = ""
	if test {
		test_prefix = "test"
	}

	// set location of log file
	var logpath = fmt.Sprintf("../log/%s_%s_log_%v.txt", test_prefix, prefix, time.Now().Format("2006-01-02_15-04-05"))

	// flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	TraceLogger = log.New(file, "TRACE: ", log.LstdFlags|log.Lshortfile)
	DebugLogger = log.New(file, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarnLogger = log.New(file, "WARN: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
	FatalLogger = log.New(file, "FATAL: ", log.LstdFlags|log.Lshortfile)

	Log.Printf("LogFile : %s", logpath)
}
