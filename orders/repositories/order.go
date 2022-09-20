package repositories

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	pb "github.com/ruancaetano/grpc-graphql-store/orders/pborders"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db,
	}
}

func (repository *OrderRepository) CreateOrder(order *pb.CreateOrderRequest) (*pb.Order, error) {
	var id string

	err := repository.db.
		QueryRow(
			"INSERT INTO orders (user_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id",
			order.GetUser(),
			order.GetProduct(),
			order.GetQuantity(),
		).
		Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Id:        id,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
		User:      order.GetUser(),
		Product:   order.GetProduct(),
		Quantity:  order.GetQuantity(),
	}, nil
}

func (repository *OrderRepository) ListUserOrders(request *pb.ListUserOrdersRequest) (*pb.ListUserOrdersResponse, error) {
	query := `SELECT id, created_at, updated_at, user_id, product_id, quantity FROM orders where user_id = $1 ORDER BY "created_at" `

	rows, err := repository.db.Query(query, request.GetUser())
	if err != nil {
		return nil, err
	}

	orders := []*pb.Order{}

	for rows.Next() {
		var (
			id         string
			created_at string
			updated_at string
			user_id    string
			product_id string
			quantity   uint32
		)

		err = rows.Scan(&id, &created_at, &updated_at, &user_id, &product_id, &quantity)

		if err != nil {
			return nil, err
		}

		orders = append(orders, &pb.Order{
			Id:        id,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			User:      user_id,
			Product:   product_id,
			Quantity:  quantity,
		})

	}

	return &pb.ListUserOrdersResponse{
		Items: orders,
	}, nil
}
