package controllers

import (
	"borrow-service/config"
	"borrow-service/models"
	"borrow-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BorrowRequest struct {
	MemberID uint `json:"member_id"`
	BookID   uint `json:"book_id"`
}

type BorrowController struct {
	BookClient   *services.BookClient
	MemberClient *services.MemberClient
}

func NewBorrowController(bookClient *services.BookClient, memberClient *services.MemberClient) *BorrowController {
	return &BorrowController{
		BookClient:   bookClient,
		MemberClient: memberClient,
	}
}

func (bc *BorrowController) CreateBorrow(c *gin.Context) {
	var request BorrowRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	member, err := bc.MemberClient.GetMember(request.MemberID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	book, err := bc.BookClient.GetBook(request.BookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	// Check availability
	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no available copies"})
		return
	}

	//  Check if member already borrowed before
	var existingBorrow models.Borrow
	result := config.DB.Where(
		"member_id = ? AND book_id = ? AND status = ?",
		member.ID, book.ID, "borrowed",
	).First(&existingBorrow)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member already borrowed this book"})
		return
	}

	// Create borrow record
	borrow := models.Borrow{
		MemberID: member.ID,
		BookID:   book.ID,
		Status:   "borrowed",
	}

	config.DB.Create(&borrow)

	// Decrease book available copies
	if err := bc.BookClient.DecreaseAvailable(book.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update book copies"})
		return
	}

	c.JSON(http.StatusCreated, borrow)
}
