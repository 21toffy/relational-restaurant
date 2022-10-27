package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/21toffy/relational-restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {

	public := incommingRoutes.Group("/api")

	private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	public.GET("/invoices", controller.GetAllInvoice())
	public.GET("/invoices/:invoice_id/:order_id", controller.GetInvoice())
	public.GET("/invoices/filter/:get_by_status_or_method", controller.GetInvoiceByStatusOrMethod())
	public.POST("/invoices", controller.CreateInvoice())
	private.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
	public.DELETE("/invoice/delete/:invoice_id", controller.DeleteInvoice())

}
