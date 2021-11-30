package service

import (
	"errors"
	"github.com/supperdoggy/spotify-web-project/spotify-auth/internal/db"
	"github.com/supperdoggy/spotify-web-project/spotify-auth/shared/structs"
	globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/night-codes/types.v1"
	"time"
)

type IService interface {
	NewToken(req structs.NewTokenReq) (resp structs.NewTokenResp, err error)
	CheckToken(req structs.CheckTokenReq) (resp structs.CheckTokenResp, err error)
	Login(req structs.LoginReq) (resp structs.LoginResp, err error)
	Register(req structs.RegisterReq) (resp structs.RegisterResp, err error)
}

type Service struct {
	logger *zap.Logger
	db db.IDB
}

func NewService(l *zap.Logger, db db.IDB) IService {
	return &Service{logger: l, db: db}
}

func (s *Service) NewToken(req structs.NewTokenReq) (resp structs.NewTokenResp, err error) {
	if req.UserID == "" {
		resp.Error = "userID cannot be empty"
		return resp, errors.New(resp.Error)
	}

	token, err := s.db.NewToken(req.UserID)
	if err != nil {
		s.logger.Error("error creating new token", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	resp.Token = token
	return
}

func (s *Service) CheckToken(req structs.CheckTokenReq) (resp structs.CheckTokenResp, err error) {
	if req.Token == "" {
		resp.Error = "Token cannot be empty"
		return resp, errors.New(resp.Error)
	}

	ok, userID := s.db.CheckToken(req.Token)
	if !ok {
		return
	}

	resp.OK = true
	resp.UserID = userID
	return
}

func (s *Service) Register(req structs.RegisterReq) (resp structs.RegisterResp, err error) {
	if req.Email == "" || req.Password == "" {
		resp.Error = "fill al the fields"
		return resp, errors.New(resp.Error)
	}

	if len(req.Password) < 7 {
		resp.Error = "password should be more than 7 chars"
		return resp, errors.New(resp.Error)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("error hashing password", zap.Error(err), zap.Any("password", req.Password))
		resp.Error = err.Error()
		return resp, err
	}

	creds := globalStructs.Creds{
		UserID: types.String(time.Now().UnixNano()),
		Email: req.Email,
		Username: req.Username,
		Password: hashed,
	}

	err = s.db.NewCreds(creds)
	if err != nil {
		s.logger.Error("error creating new creds", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	token, err := s.db.NewToken(creds.UserID)
	if err != nil {
		s.logger.Error("error generating new token", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	resp.UserID = creds.UserID
	resp.Token = token

	return
}

func (s *Service) Login(req structs.LoginReq) (resp structs.LoginResp, err error) {
	if req.Email == "" || req.Password == "" {
		resp.Error = "you must fill all the fields"
		return resp, errors.New(resp.Error)
	}

	creds, err := s.db.GetCredsByEmail(req.Email)
	if err != nil {
		s.logger.Error("error getting creds by email", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	if err := bcrypt.CompareHashAndPassword(creds.Password, []byte(req.Password)); err == nil {
		s.logger.Error("error comparing hash and password", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	token, err := s.db.NewToken(creds.UserID)
	if err != nil {
		s.logger.Error("error creating new token", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	resp.Token = token
	resp.UserID = creds.UserID
	return resp, err
}
