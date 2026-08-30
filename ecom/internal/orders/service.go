package orders

import (
	"context"
	"errors"
	"fmt"

	repo "ManogyaDahal/ecom/internal/adapters/postgresql/sqlc"

	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrInsufficientInStock = errors.New("Insufficient in stock")
)

type service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}

// service struct
type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) *svc {
	return &svc{
		repo: repo,
		db:   db,
	}
}

// business logic
func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	// validation NOTE: this part is usually done on handler's side
	if tempOrder.CustomerId == 0 {
		return repo.Order{}, fmt.Errorf("Customer id is required")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("Empty item list")
	}

	// transaction of db
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)
	// logic creating order

	// create order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerId)
	if err != nil {
		return repo.Order{}, err
	}

	// check if order exists
	for _, item := range tempOrder.Items {
		product, err := qtx.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrInsufficientInStock
		}

		// create order item
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  product.ID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, err
		}
	}

	return order, nil
}
