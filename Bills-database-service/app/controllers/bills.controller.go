package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	DTOs "github.com/skGab/Bills-database-service/app/dtos"
	"github.com/skGab/Bills-database-service/app/usecases"
)

type BillsController struct {
	Usecases *usecases.BillsUsecases
}

// GET ALL BILLS
func (b *BillsController) GetBills(gin *gin.Context) {
	// GET USER ID FROM ROUTE PARAMS AND CONVERT IT TO INT
	clientID, err := strconv.Atoi(gin.Param("id"))

	fmt.Println(clientID)

	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	// PASSA THE ID TO THE USECASE
	bills, err := b.Usecases.GetAllBills(clientID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	}

	// RETURN BILLS
	gin.JSON(http.StatusOK, bills)
}

// CREATE BILL
func (b *BillsController) CreateBill(gin *gin.Context) {
	// GET THE FORM PAYLOAD
	var formData *DTOs.CreateBillDTO

	if err := gin.ShouldBindJSON(&formData); err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	}

	// PASS IT TO THE USE CASE
	err := b.Usecases.CreateBill(formData)

	// RETURN RESPONSE
	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	} else {
		gin.JSON(http.StatusOK, struct {
			status  bool
			message string
		}{status: true, message: "Conta registrada"})
	}
}

// UPDATE BILL
func (b *BillsController) UpdateBill(gin *gin.Context) {
	// GET BILL ID
	billID, err := strconv.Atoi(gin.Param("id"))

	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

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
	billID, err := strconv.Atoi(gin.Param("id"))

	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

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
	var billsIDs []int

	if err := gin.ShouldBindJSON(billsIDs); err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	status, err := b.Usecases.DeleteAllBills(billsIDs)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
	}

	gin.JSON(http.StatusOK, status)
}
