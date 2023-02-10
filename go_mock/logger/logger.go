package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Log *log.Logger
)

func init() {
	// set location of log file
	var logpath = fmt.Sprintf("../log/log_%v.txt", time.Now().Format("2006-01-02_15-04-05"))

	flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
