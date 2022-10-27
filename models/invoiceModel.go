package models

import (
	"context"
	"errors"
	"time"

	"github.com/21toffy/relational-restaurant/database"
)

type Invoice struct {
	Id               uint      `json:"id" gorm:"primary_key"`
	Invoice_id       string    `json:"invoice_id"`
	Order_id         string    `json:"order_id"`
	Payment_method   *string   `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string   `json:"payment_status" validate:"required, eq=PENDING|eq=PAID"`
	Payment_due_date time.Time `json:"payment_due_date,omitempty"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
}

func (b *Invoice) TableName() string {
	return "invoice"
}

func CreateInvoice(invoice *Invoice) (err error) {
	if err = database.DB.Create(invoice).Error; err != nil {
		return err
	}
	return nil
}

func GetAllInvoices(invoice *[]Invoice) (err error) {
	if err = database.DB.Find(invoice).Error; err != nil {
		return err

	}
	return nil
}

func GetInvoiceByID(uid int) (Invoice, error) {
	var invoice Invoice
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := database.DB.WithContext(ctx).Model(Invoice{}).Where("id = ?", uid).Take(&invoice).Error; err != nil {
		return invoice, errors.New("Invoice not found!")
	}
	return invoice, nil
}

func GetInvoiceByPaymentMethod(Payment_method string, invoice *[]Invoice) (err error) {
	if err = database.DB.Where("Payment_method = ?", Payment_method).Find(&invoice).Error; err != nil {
		return err
	}
	return nil
}

func GetInvoiceByPaymentStatus(Payment_status string, invoice *[]Invoice) (err error) {
	if err = database.DB.Where("Payment_status = ?", Payment_status).Find(&invoice).Error; err != nil {
		return err

	}
	return nil
}

// func GetInvoiceByPaymentMethod(Payment_method string) (Invoice, error) {
// 	var invoice Invoice
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	if err := database.DB.WithContext(ctx).Model(Invoice{}).Where("Payment_method = ?", Payment_method).Find(&invoice).Error; err != nil {
// 		return invoice, errors.New("Invoice not found!")
// 	}
// 	return invoice, nil
// }

// func GetInvoiceByPaymentStatus(Payment_status string) (Invoice, error) {
// 	var invoice Invoice
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	if err := database.DB.WithContext(ctx).Model(Invoice{}).Where("Payment_status = ?", Payment_status).Find(&invoice).Error; err != nil {
// 		return invoice, errors.New("Invoice not found!")
// 	}
// 	return invoice, nil
// }
