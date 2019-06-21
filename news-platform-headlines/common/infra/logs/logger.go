package logs

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func InitLogger() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/myapp/news-platform-headlines.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	})
}
