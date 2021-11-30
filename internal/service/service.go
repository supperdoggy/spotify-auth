package service

import "go.uber.org/zap"

type IService interface {

}

type Service struct {
	logger *zap.Logger
}

func NewService(l *zap.Logger) IService {
	return &Service{logger: l}
}


