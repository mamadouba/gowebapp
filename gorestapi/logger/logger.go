package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type logWriter struct {
	io.Writer
	timeFormat string
}

func (w logWriter) Write(bytes []byte) (int, error) {
	return w.Writer.Write(append([]byte(time.Now().UTC().Format(w.timeFormat)), bytes...))
}

type Logger struct {
	name  string
	log   *log.Logger
	debug bool
}

func Info(message string, v ...interface{}) {
	Log.log.Printf("INF "+message, v...)
}

func Error(message string, v ...interface{}) {
	Log.log.Printf("ERR "+message, v...)
}

func Debug(message string, v ...interface{}) {
	if Log.debug {
		Log.log.Printf("DEB "+message, v...)
	}
}

func Warn(message string, v ...interface{}) {
	Log.log.Printf("WAR "+message, v...)
}

func Fatal(message string, v ...interface{}) {
	Log.log.Printf("FAT "+message, v...)
	os.Exit(-1)
}

func Panic(message string, v ...interface{}) {
	Log.log.Printf("FATAL "+message, v...)
	os.Exit(-1)
}

var Log *Logger

func Init(name string, debug bool) {
	fname := fmt.Sprintf("%s.log", name)
	fflag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(fname, fflag, 0666)
	if err != nil {
		panic(err.Error())
	}
	writer := io.MultiWriter(os.Stdout, file)
	lg := log.New(&logWriter{writer, "02-01-2006 15:04:05.000 "}, "", 0)
	Log = &Logger{name: name, log: lg, debug: debug}
}
