package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	DTOs "github.com/skGab/Bills-microservices/Bills-database-service/app/dtos"
	"github.com/skGab/Bills-microservices/Bills-database-service/app/usecases"
)

type BillsController struct {
	Usecases *usecases.BillsUsecases
	Validate *validator.Validate
}

// GET ALL BILLS
func (b *BillsController) GetBills(gin *gin.Context) {
	// GET USER ID FROM ROUTE PARAMS AND CONVERT IT TO INT
	clientID := gin.Param("id")

	// PASSA THE ID TO THE USECASE
	bills, err := b.Usecases.GetAllBills(clientID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// RETURN BILLS
	gin.JSON(http.StatusOK, bills)
}

// CREATE BILL
func (b *BillsController) CreateBill(gin *gin.Context) {
	// GET THE FORM PAYLOAD
	var formData *DTOs.CreateBillDTO

	if err := gin.BindJSON(&formData); err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// VALIDATE THE STRUCT
	err := b.Validate.Struct(formData)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// PASS IT TO THE USE CASE
	validationError := b.Usecases.CreateBill(formData)

	// RETURN RESPONSE
	if validationError != nil {
		gin.JSON(http.StatusInternalServerError, validationError.Error())
	} else {
		gin.JSON(http.StatusOK, struct {
			Status  bool
			Message string
		}{Status: true, Message: "Conta registrada"})
	}
}

// UPDATE BILL
func (b *BillsController) UpdateBill(gin *gin.Context) {
	// GET BILL ID
	billID := gin.Param("id")

	// DATA
	var data interface{}

	if err := gin.ShouldBindJSON(data); err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	status, err := b.Usecases.UpdateBill(billID, data)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	}

	gin.JSON(http.StatusOK, status)
}

// DELETE BILL
func (b *BillsController) DeleteBill(gin *gin.Context) {
	//GET BILL ID
	billID := gin.Param("id")

	//PASS THE ID TO THE USECASE LOGIC
	status, err := b.Usecases.DeleteBill(billID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	}

	// RETURN RESPONSE
	gin.JSON(http.StatusOK, status)
}

// DELETE ALL BILLS
func (b *BillsController) DeleteAllBills(gin *gin.Context) {
	var billsIDs []string

	if err := gin.ShouldBindJSON(billsIDs); err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	status, err := b.Usecases.DeleteAllBills(billsIDs)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	}

	gin.JSON(http.StatusOK, status)
}
