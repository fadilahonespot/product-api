package handler

import (
	"product-api/model"
	"product-api/service"

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
	summary, err := h.transactionService.Summary()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get summary",
		})
	}
	return c.JSON(summary)
}