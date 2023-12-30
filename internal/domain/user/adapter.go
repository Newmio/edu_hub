package user

import (
	"ed"
	"ed/internal/domain/logger"
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	s      IUserService
	logger logger.ILoggerService
}

func NewHandler(s IUserService, logger logger.ILoggerService) *handler {
	return &handler{s: s, logger: logger}
}

func (h *handler) InitUserRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/createAccount", h.CreateAccountRout)
	return r
}

func (h *handler) RegisterRout(c *gin.Context) {
	var acc Account

	log := h.logger.InitLog(c)

	err := json.Unmarshal([]byte(log.Body_req), &acc)
	if err != nil {
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	token, refresh, err := h.s.Register(acc.Login, acc.Pass)
	if err != nil {
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	h.logger.HttpTokenResponse(c, log, token, refresh)
}

func (h *handler) CreateAccountRout(c *gin.Context) {
	var acc Account

	log := h.logger.InitLog(c)

	err := json.Unmarshal([]byte(log.Body_req), &acc)
	if err != nil {
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	id, err := h.s.CreateAccount(&acc)
	if err != nil {
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	c.JSON(201, gin.H{"status": true, "id": id})
}

func (h *handler) UserIdentity(c *gin.Context) {
	log := h.logger.InitLog(c)

	header := c.GetHeader("Authorization")
	if header == "" {
		h.logger.HttpErrorResponse(c, log, errors.New("empty auth token"))
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2{
		h.logger.HttpErrorResponse(c, log, errors.New("invalid auth header"))
		return
	}

	id, err := h.s.ParseToken(parts[1])
	if err != nil{
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	c.Set("account_id", id)
}
