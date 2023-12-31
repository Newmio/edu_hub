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

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.CreateAccountRout)
		auth.POST("/login")
	}

	api := r.Group("/api", h.UserIdentity)
	{
		api.GET("/test", h.test)
	}

	return r
}

func (h *handler) test(c *gin.Context) {
	id, flag := c.Get("id_acc")
	c.JSON(200, gin.H{"val": id, "flag": flag})
}

func (h *handler) CreateAccountRout(c *gin.Context) {
	var acc Account

	log := h.logger.InitLog(c)

	err := json.Unmarshal([]byte(log.Body_req), &acc)
	if err != nil {
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	token, refresh, id, err := h.s.CreateAccount(&acc)
	if err != nil {
		if id == -1{
			h.logger.HttpErrorResponse(c, log, err)
			return
		}
		
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	h.logger.HttpRegisterResponse(c, log, token, refresh, id)
}

func (h *handler) UserIdentity(c *gin.Context) {
	log := h.logger.InitLog(c)

	header := c.GetHeader("Authorization")
	if header == "" {
		h.logger.HttpErrorResponse(c, log, errors.New("empty auth token"))
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		h.logger.HttpErrorResponse(c, log, errors.New("invalid auth header"))
		return
	}

	id, err := h.s.ParseToken(parts[1])
	if err != nil {
		h.logger.HttpErrorResponse(c, log, ed.ErrTrace(err, ed.Trace()))
		return
	}

	c.Set("id_acc", id)
}
