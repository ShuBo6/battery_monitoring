package internal

import (
	"os"
	"path"
	"time"

	logs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := logs.New(
		path.Join(global.C.Zap.Director, "%Y-%m-%d.log"),
		logs.WithMaxAge(7*24*time.Hour),
		logs.WithRotationTime(24*time.Hour),
	)
	if global.C.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
