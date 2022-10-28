package handler

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rostis232/traineeEVOFintech/pkg/service"
	"io"
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
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

		//TODO: needs to return error?
		func(timestamp string) {
			f, err := os.Open(fmt.Sprintf("./uploads/%s", timestamp+".csv"))
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			csvReader := csv.NewReader(f)
			for {
				rec, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				//TODO: Adding new values to DB
				for _, v := range rec {
					fmt.Printf("%s\n", v)
				}

			}
		}(timestamp)
	}

}

func (h *Handler) getJson(c *gin.Context) {}

func (h *Handler) getCsv(c *gin.Context) {}
