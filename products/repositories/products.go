package repositories

import (
	"database/sql"
	"math"
	"time"

	_ "github.com/lib/pq"

	pb "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (repository *ProductRepository) CreateProduct(product *pb.CreateProductRequest) (*pb.Product, error) {
	var id string

	err := repository.db.
		QueryRow(
			"INSERT INTO products (title, description, thumb, availables) VALUES ($1, $2, $3, $4) RETURNING id",
			product.GetTitle(),
			product.GetDescription(),
			product.GetThumb(),
			product.GetAvailables(),
		).
		Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          id,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		Title:       product.GetTitle(),
		Description: product.GetDescription(),
		Thumb:       product.GetThumb(),
		Availables:  product.GetAvailables(),
	}, nil
}

func (repository *ProductRepository) UpdateProduct(product *pb.UpdateProductRequest) (*pb.Product, error) {
	var id string

	_, err := repository.db.
		Query(
			"UPDATE products set updated_at = now(), title = $1, description = $2, thumb = $3, where id = $4",
			product.GetTitle(),
			product.GetDescription(),
			product.GetThumb(),
			product.GetId(),
		)
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          id,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		Title:       product.GetTitle(),
		Description: product.GetDescription(),
		Thumb:       product.GetThumb(),
	}, nil
}

func (repository *ProductRepository) UpdateProductAvailablesValue(req *pb.UpdateProductAvailablesValueRequest) (*pb.Product, error) {
	_, err := repository.db.
		Query(
			"UPDATE products set availables = availables + $1 where id = $2",
			req.GetValueToAdd(),
			req.GetId(),
		)
	if err != nil {
		return nil, err
	}

	return repository.GetProductById(req.GetId())
}

func (repository *ProductRepository) DeleteProduct(product *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	_, err := repository.db.
		Query(
			"UPDATE products set updated_at = now(), is_active = false where id = $1",
			product.GetId(),
		)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Deleted: true,
	}, nil
}

func (repository *ProductRepository) ListProducts(pagination *pb.PaginationParams) (*pb.ProductListResponse, error) {
	limit := uint32(math.Min(float64(pagination.GetLimit()), 100))
	page := pagination.GetPage()

	offset := limit * (page - 1)

	query := `SELECT id, created_at, updated_at, title, thumb, description, availables FROM products ORDER BY "id" LIMIT $1 OFFSET $2`

	rows, err := repository.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	products := []*pb.Product{}

	for rows.Next() {
		var (
			id          string
			created_at  string
			updated_at  string
			title       string
			description string
			thumb       string
			availables  uint32
		)

		err = rows.Scan(&id, &created_at, &updated_at, &title, &thumb, &description, &availables)

		if err != nil {
			return nil, err
		}

		products = append(products, &pb.Product{
			Id:          id,
			CreatedAt:   created_at,
			UpdatedAt:   updated_at,
			Title:       title,
			Description: description,
			Thumb:       thumb,
			Availables:  availables,
		})

	}

	return &pb.ProductListResponse{
		Items: products,
	}, nil
}

func (repository *ProductRepository) GetProductById(id string) (*pb.Product, error) {
	query := `SELECT created_at, updated_at, title, thumb, description, availables FROM products where id = $1`

	var (
		created_at  string
		updated_at  string
		title       string
		description string
		thumb       string
		availables  uint32
	)

	err := repository.db.QueryRow(query, id).Scan(&created_at, &updated_at, &title, &thumb, &description, &availables)
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          id,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
		Title:       title,
		Description: description,
		Thumb:       thumb,
		Availables:  availables,
	}, nil
}
