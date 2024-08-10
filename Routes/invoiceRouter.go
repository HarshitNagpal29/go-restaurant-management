package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", Controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", Controllers.GetInvoice())
	incomingRoutes.POST("/invoices", Controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", Controllers.UpdateInvoice())
	
}