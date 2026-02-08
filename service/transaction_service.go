package service

import (
	"errors"
	"product-api/model"
	"product-api/repository"
	"time"
)

type TransactionServiceInterface interface {
	Checkout(checkoutRequest *model.CheckoutRequest) (model.Transaction, error)
	Summary() (model.SummaryResponse, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepositoryInterface
	productRepo     repository.ProductRepositoryInterface
}

func NewTransactionService(transactionRepo repository.TransactionRepositoryInterface, productRepo repository.ProductRepositoryInterface) TransactionServiceInterface {
	return &transactionService{
		transactionRepo: transactionRepo,
		productRepo:     productRepo,
	}
}

func (s *transactionService) Checkout(checkoutRequest *model.CheckoutRequest) (model.Transaction, error) {
	tx, _ := s.productRepo.BeginTrans()
	transaction := model.Transaction{}
	for _, item := range checkoutRequest.Items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			s.productRepo.RollbackTrans(tx)
			return model.Transaction{}, err
		}
		if product.Stock < item.Quantity {
			s.productRepo.RollbackTrans(tx)
			return model.Transaction{}, errors.New("product stock not enough")
		}
		product.Stock -= item.Quantity
		err = s.productRepo.Update(tx, product)
		if err != nil {
			s.productRepo.RollbackTrans(tx)
			return model.Transaction{}, err
		}
		transactionDetails := model.TransactionDetail{
			ProductID:   product.ID,
			Quantity:    item.Quantity,
			ProductName: product.Name,
			Subtotal:    product.Price * item.Quantity,
		}
		transaction.Details = append(transaction.Details, transactionDetails)
		transaction.TotalAmount += transactionDetails.Subtotal
	}

	err := s.transactionRepo.Create(tx, &transaction)
	if err != nil {
		s.productRepo.RollbackTrans(tx)
		return model.Transaction{}, err
	}
	s.productRepo.CommitTrans(tx)
	return transaction, nil
}

func (s *transactionService) Summary() (model.SummaryResponse, error) {
	timeNow := time.Now()
	fromDate := timeNow.Format("2006-01-02") + " 00:00:00"
	toDate := timeNow.Format("2006-01-02") + " 23:59:59"
	transactions, err := s.transactionRepo.GetAll(fromDate, toDate)
	if err != nil {
		return model.SummaryResponse{}, err
	}
	var summary model.SummaryResponse
	var productTerlaris = map[string]int{}
	for _, transaction := range transactions {
		summary.TotalRevenue += transaction.TotalAmount
		summary.TotalTransaction += 1

		for _, detail := range transaction.Details {
			product, err := s.productRepo.GetByID(detail.ProductID)
			if err != nil {
				return model.SummaryResponse{}, err
			}
			productTerlaris[product.Name] += detail.Quantity
		}
	}

	for productName, qtyTerjual := range productTerlaris {
		if qtyTerjual > summary.ProductTerlaris.QtyTerjual {
			summary.ProductTerlaris.Name = productName
			summary.ProductTerlaris.QtyTerjual = qtyTerjual
		}
	}
	return summary, nil
}
