package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/media/nik/7ec12544-2887-498c-85e9-bb82dfe7ffcc/nik/Setups/news-platform/news.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Info("There is a problem")

	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"attempt", 3,
		"backoff", time.Second,
	)

	return sugar
}
