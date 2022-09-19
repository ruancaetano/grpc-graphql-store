package services

import (
	"context"

	pb "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"

	"github.com/ruancaetano/grpc-graphql-store/products/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
	pb.UnimplementedProductServiceServer
}

func (service *ProductService) CreateProduct(contexts context.Context, request *pb.CreateProductRequest) (*pb.Product, error) {
	return service.repository.CreateProduct(request)
}

func (service *ProductService) ListProducts(contexts context.Context, request *pb.PaginationParams) (*pb.ProductListResponse, error) {
	return service.repository.ListProducts(request)
}

func (service *ProductService) UpdateProduct(contexts context.Context, request *pb.UpdateProductRequest) (*pb.Product, error) {
	return service.repository.UpdateProduct(request)
}

func (service *ProductService) DeleteProduct(contexts context.Context, request *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return service.repository.DeleteProduct(request)
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repository,
		pb.UnimplementedProductServiceServer{},
	}
}
