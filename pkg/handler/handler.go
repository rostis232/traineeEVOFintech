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

	//TODO: filename validation?
	timestamp := time.Now().Format("02-01-06_15:04:05")
	err = c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", timestamp+".csv"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("'%s' not uploaded!", file.Filename))
	} else {
		transactions, err := csvToStruct(timestamp)

		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: '%s'", err))
		}

		err = h.services.Transaction.InsertToDB(transactions)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: '%s'", err))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded and successful added to DB!", file.Filename))
		}

	}

}

func (h *Handler) getJson(c *gin.Context) {
	var params = map[string]string{}
	params["transactionId"] = c.Query("transaction_id")
	params["terminalId"] = c.Query("terminal_id") //Can be more than only one ID
	params["status"] = c.Query("status")
	params["paymentType"] = c.Query("payment_type")
	params["datePostFrom"] = c.Query("date_post_from")
	params["datePostTo"] = c.Query("date_post_to")
	params["paymentNarrative"] = c.Query("payment_narrative")

	transactions, err := h.services.Transaction.GetJSON(params)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: '%s'", err))
	} else {
		c.JSON(http.StatusOK, transactions)
	}
}

func (h *Handler) getCsv(c *gin.Context) {}

// TODO: needs to return an error?
func csvToStruct(timestamp string) ([]traineeEVOFintech.Transaction, error) {
	file, err := os.Open(fmt.Sprintf("./uploads/%s", timestamp+".csv"))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	transactions := []traineeEVOFintech.Transaction{}

	if err := gocsv.UnmarshalFile(file, &transactions); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return transactions, nil
}
