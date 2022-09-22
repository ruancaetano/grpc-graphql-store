package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	pb "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (repository *UserRepository) InsertUser(name string, email string, password string) (*pb.User, error) {
	var id string

	err := repository.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", name, email, password).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    id,
		Name:  name,
		Email: email,
		Role:  "user",
	}, nil
}

func (repository *UserRepository) GetUserById(id string) (*pb.User, error) {
	var (
		name       string
		email      string
		role       string
		created_at string
		updated_at string
	)

	err := repository.db.QueryRow(
		"select name, email, role, created_at, updated_at from users where id = $1 and is_active = true;", id,
	).Scan(
		&name, &email, &role, &created_at, &updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}

func (repository *UserRepository) GetUserByEmail(email string) (*pb.User, error) {
	var (
		id         string
		name       string
		role       string
		password   string
		created_at string
		updated_at string
	)

	err := repository.db.QueryRow(
		"select id, name, password, role, created_at, updated_at from users where email = $1 and is_active = true;", email,
	).Scan(
		&id, &name, &password, &role, &created_at, &updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      role,
		Password:  password,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}
