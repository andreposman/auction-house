package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/andreposman/action-house-api/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	queries *pgstore.Queries
	pool    *pgxpool.Pool
}

var ErrDuplicateUniqueField = errors.New("username or email already exists")

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error hashing password: %w", err)
	}

	args := pgstore.CreateUserParams{
		UserName:     userName,
		PasswordHash: hash,
		Email:        email,
		Bio:          bio,
	}

	id, err := us.queries.CreateUser(ctx, args)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == "23505" {
			return uuid.UUID{}, ErrDuplicateUniqueField
		}
		return uuid.UUID{}, err
	}

	fmt.Println("User created sucessfully: ", id)

	return id, nil
}
