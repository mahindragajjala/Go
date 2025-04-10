package common

import (
	"fmt"
    "time"
	"runtime"
	"path/filepath"
	"gopkg.in/natefinch/lumberjack.v2"
)

//var myByteArray []byte = []byte("Hello, MarketSplash!")

var (
	dbLogger * lumberjack.Logger;
	apLogger * lumberjack.Logger;
	pfLogger * lumberjack.Logger;
)

func InitLogger( logdb bool, logapp bool, logper bool) {
	
	if logdb == true {
		dbLogger = &lumberjack.Logger{
			// Log file abbsolute path, os agnostic
			Filename:   "./dblogs/dblog.log", 
			MaxSize:    1, // MB
			MaxBackups: 10,
			MaxAge:     30,   // days
			Compress:   false, // disabled by default
		}
	}
	
	if logapp == true {
		apLogger = &lumberjack.Logger{
			// Log file abbsolute path, os agnostic
			Filename:   "./applogs/app.log", 
			MaxSize:    1, // MB
			MaxBackups: 10,
			MaxAge:     30,   // days
			Compress:   false, // disabled by default
		}
	}
	
	if logper == true {
		pfLogger = &lumberjack.Logger{
			// Log file abbsolute path, os agnostic
			Filename:   "./applogs/perf.log", 
			MaxSize:    1, // MB
			MaxBackups: 10,
			MaxAge:     30,   // days
			Compress:   false, // disabled by default
		}
	}	
}


func WriteLine( logtype string , str string, logger * lumberjack.Logger) {
	currentTime := time.Now()
	_, filename, line, _ := runtime.Caller(2);
	msg := fmt.Sprintf( "%s|%s|%s %s,%d\n", logtype,  currentTime.Format("01-02-2006 15:04:05"), str, filepath.Base(filename), line)
	logger.Write( []byte(msg));
}


func AppLogDebug( str string) {
	if apLogger != nil {
		if Config.DBLog.LogDebug == true {
			WriteLine( "D", str, apLogger)
		}
	}
}

func AppLogError( str string) {
	if apLogger != nil {
		WriteLine( "E", str, apLogger)
	}
}

func AppLogCritical( str string) {
	if apLogger != nil {
		WriteLine( "C", str, apLogger)
	}
}


func DbLogDebug( str string) {
	if dbLogger != nil {
		if Config.DBLog.LogDebug == true {
			WriteLine( "D", str, dbLogger)
		}
	}
}

func DbLogError( str string) {
	if dbLogger != nil {
		WriteLine( "E", str, dbLogger)
	}
}

func DbLogCritical( str string) {
	if dbLogger != nil {
		WriteLine( "C", str, dbLogger)
	}
}

func PerfLog( str string) {
	if pfLogger != nil {
		pfLogger.Write( []byte(str));
	}
}

