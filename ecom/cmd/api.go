package main

import (
	"log"
	"net/http"
	"time"

	product "ManogyaDahal/ecom/internal/products"
	repo "ManogyaDahal/ecom/internal/adapters/postgresql/sqlc"
	orders "ManogyaDahal/ecom/internal/orders"

	"github.com/jackc/pgx/v5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/* Global struct that my api going to have*/
type api struct {
	config *config
	db     *pgx.Conn
}

// mount
func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)              // important for rate limiting
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good everything's working"))
	})

	// routers
	productService := product.NewService(repo.New(app.db))
	productHandler := product.NewHandler(productService)
	r.Get("/product", productHandler.ListProductHandler)
	r.Get("/product/{id}", productHandler.GetProductByIDHandler)


	ordersService  := orders.NewService(repo.New(app.db), app.db)
	ordersHandlers := orders.NewHandler(ordersService)
	r.Post("/orders", ordersHandlers.PlaceOrder)

	return r
}

// run
func (app *api) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Server has starder at addr %s", srv.Addr)

	return srv.ListenAndServe()
}

/* defines the config */
type config struct {
	addr string
	db   dbConfig
}

// database configuration
type dbConfig struct {
	dsn string
}
