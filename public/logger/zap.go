package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	module ModuleFlag
}

var s_logger *Logger
var s_moduleFlag ModuleFlag = SharedInstance

func newZap(level zapcore.Level) *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger
}

func init() {
	logger := newZap(zapcore.DebugLevel)
	s_logger = &Logger{logger, SharedInstance}
	// writer1 := &bytes.Buffer{}
	// writer2 := os.Stdout
	// writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatalf("create file log.txt failed: %v", err)
	// }

	// s_logger = log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
	const msg = "logger init success"
	s_logger.Info(msg)
	s_logger.Warn(msg)
	s_logger.Error(msg)
}

func SetModuleEnabled(module ModuleFlag, enabled bool) {
	if enabled {
		s_moduleFlag |= module
	} else {
		s_moduleFlag ^= module
	}
}

func New(module ModuleFlag) *Logger {
	logger := newZap()
	SetModuleEnabled(module, true)
	return &Logger{logger, module}
}

func (l *Logger) Debug(args ...interface{}) {
	if s_moduleFlag&l.module > 0 {
		defer s_logger.Sync()
		msg := fmt.Sprint(args...)
		l.Logger.Debug(fmt.Sprintf("[%s] %s", modules[l.module], msg))
	}
}
