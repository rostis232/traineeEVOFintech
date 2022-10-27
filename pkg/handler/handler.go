package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rostis232/traineeEVOFintech/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/get-json", h.getJson)
	router.GET("/get-csv", h.getCsv)
	router.POST("/upload-csv", h.uploadCsv)

	return router
}

func (h *Handler) uploadCsv(c *gin.Context) {}

func (h *Handler) getJson(c *gin.Context) {}

func (h *Handler) getCsv(c *gin.Context) {}
