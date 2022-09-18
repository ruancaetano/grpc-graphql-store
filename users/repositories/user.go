package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ruancaetano/grpc-graphql-store/users/pb"
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
	}, nil
}

func (repository *UserRepository) GetUserById(id string) (*pb.User, error) {
	var (
		name  string
		email string
	)

	err := repository.db.QueryRow("select name, email from users where id = $1;", id).Scan(&name, &email)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (repository *UserRepository) GetUserByEmail(email string) (*pb.User, error) {
	var (
		id       string
		name     string
		password string
	)

	err := repository.db.QueryRow("select id, name, password from users where email = $1;", email).Scan(&id, &name, &password)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
