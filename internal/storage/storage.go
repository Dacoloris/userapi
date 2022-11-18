package storage

import (
	"go.uber.org/zap"
)

type storage struct {
	logger   *zap.Logger
	fileName string
}

func New(logger *zap.Logger, fileName string) *storage {
	return &storage{
		logger:   logger,
		fileName: fileName,
	}
}
