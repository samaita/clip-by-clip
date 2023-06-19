package log

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func InitLogger() {
	Sugar = zap.NewExample().Sugar()
}
