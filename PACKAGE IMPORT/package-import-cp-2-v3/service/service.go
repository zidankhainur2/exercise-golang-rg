package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}

	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	// Check if product already in cart
	for i, item := range cartItems {
		if item.ProductName == productName {
			cartItems[i].Quantity += quantity
			return s.database.SaveCartItems(cartItems)
		}
	}

	// Add new product to cart
	cartItems = append(cartItems, entity.CartItem{
		ProductName: productName,
		Price:       product.Price,
		Quantity:    quantity,
	})

	return s.database.SaveCartItems(cartItems)
}

func (s *Service) RemoveCart(productName string) error {
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	for i, item := range cartItems {
		if item.ProductName == productName {
			cartItems = append(cartItems[:i], cartItems[i+1:]...)
			return s.database.SaveCartItems(cartItems)
		}
	}

	return errors.New("product not found")
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	err := s.database.SaveCartItems([]entity.CartItem{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	products := s.database.GetProductData()
	return products, nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	totalPrice := 0
	for _, item := range cartItems {
		totalPrice += item.Price * item.Quantity
	}

	if money < totalPrice {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	change := money - totalPrice

	paymentInfo := entity.PaymentInformation{
		ProductList: cartItems,
		TotalPrice:  totalPrice,
		MoneyPaid:   money,
		Change:      change,
	}

	// Reset cart after payment
	err = s.ResetCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	return paymentInfo, nil
}
