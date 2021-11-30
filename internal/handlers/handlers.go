package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/spotify-web-project/spotify-auth/internal/service"
	"github.com/supperdoggy/spotify-web-project/spotify-auth/shared/structs"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	logger *zap.Logger
	s service.IService
}

func NewHandlers(l *zap.Logger, s service.IService) Handlers {
	return Handlers{logger: l, s:s}
}

func (h *Handlers) NewToken(c *gin.Context) {
	var req structs.NewTokenReq
	var resp structs.NewTokenResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binging request", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.NewToken(req)
	if err != nil {
		h.logger.Error("error creating new token", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) CheckToken(c *gin.Context) {
	var req structs.CheckTokenReq
	var resp structs.CheckTokenResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding request", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.CheckToken(req)
	if err != nil {
		h.logger.Error("error checking request", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) Register(c *gin.Context) {
	var req structs.RegisterReq
	var resp structs.RegisterResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding request", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.Register(req)
	if err != nil {
		h.logger.Error("error checking request", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) Login(c *gin.Context) {
	var req structs.LoginReq
	var resp structs.LoginResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding request", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.Login(req)
	if err != nil {
		h.logger.Error("error checking request", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
