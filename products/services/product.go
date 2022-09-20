package services

import (
	"context"
	"errors"

	pb "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"

	"github.com/ruancaetano/grpc-graphql-store/products/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
	pb.UnimplementedProductServiceServer
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repository,
		pb.UnimplementedProductServiceServer{},
	}
}

func (service *ProductService) ListProducts(contexts context.Context, request *pb.PaginationParams) (*pb.ProductListResponse, error) {
	return service.repository.ListProducts(request)
}

func (service *ProductService) GetProductById(contexts context.Context, request *pb.GetProductByIdRequest) (*pb.Product, error) {
	return service.repository.GetProductById(request.GetId())
}

func (service *ProductService) ValidateProductAvailability(contexts context.Context, request *pb.ValidateProductAvailabilityRequest) (*pb.ValidateProductAvailabilityResponse, error) {
	product, err := service.repository.GetProductById(request.GetId())
	if err != nil {
		return nil, err
	}

	if request.GetQuantity() > product.GetAvailables() {
		return nil, errors.New("Product quantity unavailable")
	}

	return &pb.ValidateProductAvailabilityResponse{
		Available: true,
	}, nil
}

func (service *ProductService) UpdateProductAvailablesValue(contexts context.Context, request *pb.UpdateProductAvailablesValueRequest) (*pb.Product, error) {
	return service.repository.UpdateProductAvailablesValue(request)
}

func (service *ProductService) CreateProduct(contexts context.Context, request *pb.CreateProductRequest) (*pb.Product, error) {
	return service.repository.CreateProduct(request)
}

func (service *ProductService) UpdateProduct(contexts context.Context, request *pb.UpdateProductRequest) (*pb.Product, error) {
	return service.repository.UpdateProduct(request)
}

func (service *ProductService) DeleteProduct(contexts context.Context, request *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return service.repository.DeleteProduct(request)
}
