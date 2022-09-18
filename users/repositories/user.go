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

func (repository *UserRepository) InsertUser(name string, email string) (*pb.User, error) {
	var id int

	err := repository.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", name, email).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    int32(id),
		Name:  name,
		Email: email,
	}, nil
}

func (repository *UserRepository) GetUserById(id int32) (*pb.User, error) {
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
