package handlers

import "go.uber.org/zap"

type Handlers struct {
	logger *zap.Logger

}

func NewHandlers(l *zap.Logger) Handlers {
	return Handlers{logger: l}
}
