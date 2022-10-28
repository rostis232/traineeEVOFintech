package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/rostis232/traineeEVOFintech"
	"github.com/rostis232/traineeEVOFintech/pkg/service"
	"log"
	"net/http"
	"os"
	"time"
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

func (h *Handler) uploadCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}

	//TODO: filename validation
	timestamp := time.Now().Format("02-01-06_15:04:05")
	err = c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", timestamp+".csv"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("'%s' not uploaded!", file.Filename))
	} else {
		CSVToStruct(timestamp)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}

}

func (h *Handler) getJson(c *gin.Context) {}

func (h *Handler) getCsv(c *gin.Context) {}

// TODO: needs to return error?
func CSVToStruct(timestamp string) []*traineeEVOFintech.Transaction {
	file, err := os.Open(fmt.Sprintf("./uploads/%s", timestamp+".csv"))
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	transactions := []*traineeEVOFintech.Transaction{}

	if err := gocsv.UnmarshalFile(file, &transactions); err != nil {
		log.Fatal(err)
	}
	for _, transaction := range transactions {
		fmt.Println("Hello, ", transaction.DateInput)
	}
	return transactions
}
