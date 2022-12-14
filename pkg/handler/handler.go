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

	router.GET("/get-json", h.getJSON)
	router.GET("/get-csv-file", h.getCSVFile)
	router.POST("/upload-csv", h.uploadCSV)

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
func (h *Handler) uploadCSV(c *gin.Context) {
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

// @Summary getJSON
// @Tags Transactions
// @Descriptions get response in JSON format
// @ID get_json
// @Produce json
// @Param transaction_id query string false "Transaction ID"
// @Param terminal_id query string false "Terminal ID"
// @Param status query string false "Status"
// @Param payment_type query string false "Payment Type"
// @Param date_post_from query string false "Date Post From (Example: 2022-08-17)"
// @Param date_post_to query string false "Date Post To (Example: 2022-08-17)"
// @Param payment_narrative query string false "Payment Narrative (Example: '?????? ?????????????? ????????????')"
// @Success 200 {object} []traineeEVOFintech.Transaction
// @Failure 400 {string}  string
// @Router /get-json [get]
func (h *Handler) getJSON(c *gin.Context) {
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

// @Summary getCSVFile
// @Tags Transactions
// @Descriptions get generated CSV file in response
// @ID get_csv_file
// @Produce plain
// @Param transaction_id query string false "Transaction ID"
// @Param terminal_id query string false "Terminal ID"
// @Param status query string false "Status"
// @Param payment_type query string false "Payment Type"
// @Param date_post_from query string false "Date Post From (Example: 2022-08-17)"
// @Param date_post_to query string false "Date Post To (Example: 2022-08-17)"
// @Param payment_narrative query string false "Payment Narrative (Example: '?????? ?????????????? ????????????')"
// @Success 200 {string}  string
// @Failure 400 {string}  string
// @Router /get-csv-file [get]
func (h *Handler) getCSVFile(c *gin.Context) {
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

	transactions, err := h.services.Transaction.GetCSVFile(params)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: '%s'", err))
	} else {
		timestamp := time.Now().Format("02-01-06_15:04:05")
		os.Create(fmt.Sprintf("./created_csv/%s", timestamp+".csv"))
		file, err := os.OpenFile(fmt.Sprintf("./created_csv/%s", timestamp+".csv"), os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: '%s'", err))
		}

		defer file.Close()

		err = gocsv.MarshalFile(&transactions, file)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: '%s'", err))
		} else {
			c.FileAttachment(fmt.Sprintf("./created_csv/%s", timestamp+".csv"), "export.csv")
		}
	}
}

func csvToStruct(timestamp string) ([]traineeEVOFintech.Transaction, error) {
	file, err := os.Open(fmt.Sprintf("./uploads/%s", timestamp+".csv"))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	transactions := []traineeEVOFintech.Transaction{}

	// make channel
	c := make(chan traineeEVOFintech.Transaction)

	go func() { // start parsing the CSV file
		err = gocsv.UnmarshalToChan(file, c) // <---- here it is
		if err != nil {
			log.Fatal(err)
		}
	}()

	for r := range c {
		transactions = append(transactions, r)
	}

	//if err := gocsv.UnmarshalFile(file, &transactions); err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}

	return transactions, nil
}
