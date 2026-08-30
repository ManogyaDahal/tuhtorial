package product

import (
	"ManogyaDahal/ecom/internal/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type handler struct { 
	service Service
}

func NewHandler(service Service) *handler { 
	return &handler { 
		service: service,
	}
}

func (h *handler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := h.service.ListProducts(r.Context()) 
	if err != nil { 
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}

func (h *handler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request){ 
	// extraction and validation of data
	productId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { 
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err :=  h.service.GetProductByID(r.Context(), productId)
	if err != nil { 
		if errors.Is(err, pgx.ErrNoRows) { 
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
