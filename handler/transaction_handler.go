package handler

import (
	"product-api/model"
	"product-api/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionService service.TransactionServiceInterface
}

func NewTransactionHandler(transactionService service.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) Create(c *fiber.Ctx) error {
	var request model.CheckoutRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	transaction, err := h.transactionService.Checkout(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(transaction)
}

func (h *TransactionHandler) Summary(c *fiber.Ctx) error {
	timeNow := time.Now()
	fromDate := timeNow.Format("2006-01-02") + " 00:00:00"
	toDate := timeNow.Format("2006-01-02") + " 23:59:59"
	summary, err := h.transactionService.Summary(fromDate, toDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get summary",
		})
	}
	return c.JSON(summary)
}

func (h *TransactionHandler) SummaryByDate(c *fiber.Ctx) error {
	fromDate := c.Query("start_date") + " 00:00:00"
	toDate := c.Query("end_date") + " 23:59:59"
	summary, err := h.transactionService.Summary(fromDate, toDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get summary",
		})
	}
	return c.JSON(summary)
}
