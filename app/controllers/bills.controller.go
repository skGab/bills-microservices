package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	DTOs "skGab/Bills-management-service/app/dtos"
	"skGab/Bills-management-service/app/usecases"
)

type BillsController struct {
	Usecases *usecases.BillsUsecases
}

func (b *BillsController) CreateBill(g *gin.Context) {
	// GET THE FORM PAYLOAD
	var formData *DTOs.CreateBillDTO

	if err := g.ShouldBindJSON(&formData); err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	// PASS IT TO THE USE CASE
	err := b.Usecases.CreateBill(formData)

	// RETURN RESPONSE
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	} else {
		g.JSON(http.StatusOK, struct {
			status  bool
			message string
		}{status: true, message: "Conta registrada"})
	}
}
