package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/21toffy/relational-restaurant/database"
	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   *string
	Order_id         string
	Payment_status   *string
	Payment_due      time.Time
	Payment_due_date time.Time
	Order_details    OrderDisplay
}

type InvoiceCreate struct {
	Invoice_id       string
	Order_id         string
	Payment_method   string
	Payment_status   string
	Payment_due      time.Time
	Table_number     int64
	Payment_due_date time.Time `json:"payment_due_date,omitempty"`
	Created_at       time.Time
	Updated_at       time.Time
}

type InvoiceUpdate struct {
	Order_id         string
	Payment_method   string
	Payment_status   string
	Payment_due      time.Time
	Payment_due_date time.Time `json:"payment_due_date,omitempty"`
	Updated_at       time.Time
}

type InvoiceListViewFormat struct {
	Invoice_id       string
	Order_id         string
	Payment_method   *string
	Payment_status   *string
	Payment_due      time.Time
	Table_number     int64
	Payment_due_date time.Time
}

func GetAllInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice []models.Invoice
		err := models.GetAllInvoices(&invoice)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"data": invoice})
		}
	}
}

func GetInvoiceByStatusOrMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		filter_parameter := c.Param("get_by_status_or_method")
		var invoice []models.Invoice
		var err error
		if strings.ToLower(filter_parameter) == "cash" || strings.ToLower(filter_parameter) == "card" {
			err = models.GetInvoiceByPaymentMethod(strings.ToUpper(filter_parameter), &invoice)

		} else if strings.ToLower(filter_parameter) == "pending" || strings.ToLower(filter_parameter) == "paid" {
			err = models.GetInvoiceByPaymentStatus(strings.ToUpper(filter_parameter), &invoice)

		} else {
			msg := fmt.Sprintf("%s is an invalid filter parameter", filter_parameter)
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)

		} else {
			c.JSON(http.StatusOK, gin.H{"data": invoice, "count": len(invoice)})
		}
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input InvoiceCreate
		c.BindJSON(&input)
		if input.Order_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "OrderID can not be empty"})
			return
		}
		if input.Payment_method == "" || input.Payment_status == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "payment method or payment status can not be empty"})
			return
		}
		if input.Table_number < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Table number needs to be passed"})
			return
		}
		invoice := models.Invoice{
			Invoice_id:       helpers.GenerateUUID(),
			Order_id:         input.Order_id,
			Payment_method:   &input.Payment_method,
			Payment_status:   &input.Payment_status,
			Payment_due_date: input.Payment_due_date,
		}
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		err := models.CreateInvoice(&invoice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully", "data": input})
			return
		}
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceId, _ := strconv.Atoi(c.Param("invoice_id"))
		orderId := c.Param("order_id")

		var invoice models.Invoice
		invoice, err := models.GetInvoiceByID(invoiceId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderData, err := models.GetOrderByUID(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tableData, err := models.GetTableByUID(*orderData.Table_id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var invoice_details InvoiceViewFormat
		invoice_details.Invoice_id = invoice.Invoice_id
		invoice_details.Payment_method = invoice.Payment_method
		invoice_details.Order_id = orderId
		invoice_details.Payment_status = invoice.Payment_status
		invoice_details.Payment_due = invoice.Payment_due_date
		invoice_details.Payment_due_date = invoice.Payment_due_date
		invoice_details.Order_details.Table.Number_of_guests = tableData.Number_of_guests
		invoice_details.Order_details.Table.Table_number = tableData.Table_number
		invoice_details.Order_details.Order_Date = orderData.Order_Date

		c.JSON(http.StatusOK, invoice_details)
		return
	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceId, _ := strconv.Atoi(c.Param("invoice_id"))
		var invoice models.Invoice
		invoice, err := models.GetInvoiceByID(invoiceId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := database.DB.WithContext(ctx).Model(invoice).Updates(models.Invoice{Deleted: true}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
		return

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		invoiceId, _ := strconv.Atoi(c.Param("invoice_id"))
		var input InvoiceUpdate
		invoiceData, err := models.GetInvoiceByID(invoiceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Validate input
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if input.Order_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": " Order ID can not be empty"})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		if err := database.DB.WithContext(ctx).Model(&invoiceData).Updates(models.Invoice{
			Order_id:         input.Order_id,
			Payment_method:   &input.Payment_method,
			Payment_status:   &input.Payment_status,
			Payment_due_date: input.Payment_due_date,
			Updated_at:       helpers.GetCurrentTime(),
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": input})
		return

	}
}
