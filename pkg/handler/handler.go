package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/rostis232/traineeEVOFintech"
	_ "github.com/rostis232/traineeEVOFintech/docs"
	"github.com/rostis232/traineeEVOFintech/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"strings"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

// @Summary uploadCSV
// @Tags Transactions
// @Descriptions upload CSV
// @ID upload_csv
// @Accept mpfd
// @Produce plain
// @Param file formData file true "CSV file"
// @Success 200 {string}  string
// @Failure 400 {string}  string
// @Router /upload-csv [post]
func (h *Handler) uploadCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}

	//Filename verification
	if !strings.HasSuffix(file.Filename, ".csv") {
		c.String(http.StatusBadRequest, "File extension isn`t CSV")
		return
	}

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

// @Summary getJON
// @Tags Transactions
// @Descriptions get JSON
// @ID get_json
// @Produce json
// @Param transaction_id query string false "Transaction ID"
// @Success 200 {object} []traineeEVOFintech.Transaction
// @Failure 400 {string}  string
// @Router /upload-csv [get]
func (h *Handler) getJson(c *gin.Context) {
	var params = map[string]string{}
	if c.Query("transaction_id") != "" {
		params["transactionId"] = c.Query("transaction_id")
	}
	if c.Query("terminal_id") != "" {
		params["terminalId"] = c.Query("terminal_id") //Can be more than only one ID
	}
	if c.Query("status") != "" {
		params["status"] = c.Query("status")
	}
	if c.Query("payment_type") != "" {
		params["paymentType"] = c.Query("payment_type")
	}
	if c.Query("date_post_from") != "" {
		params["datePostFrom"] = c.Query("date_post_from")
	}
	if c.Query("date_post_to") != "" {
		params["datePostTo"] = c.Query("date_post_to")
	}
	if c.Query("payment_narrative") != "" {
		params["paymentNarrative"] = c.Query("payment_narrative")
	}

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
