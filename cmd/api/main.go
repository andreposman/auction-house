package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/andreposman/auction-house-api/internal/api"
	"github.com/andreposman/auction-house-api/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Auction House API")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	connString := buildDatabaseConnString()
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode //? restringe o envio entre reqs de outras origens

	api := api.API{
		Router:      chi.NewMux(),
		UserService: services.NewUserService(pool),
	}

	api.BindRoutes()

	fmt.Println("Postgres conn string: ", connString)
	fmt.Println("Starting server on port :3080")
	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}

func buildDatabaseConnString() string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("ACTION_HOUSE_DB_USER"),
		os.Getenv("ACTION_HOUSE_DB_PASSWORD"),
		os.Getenv("ACTION_HOUSE_DB_HOST"),
		os.Getenv("ACTION_HOUSE_DB_PORT"),
		os.Getenv("ACTION_HOUSE_DB_NAME"),
	)
	return connString
}
