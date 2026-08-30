package orders

import (
	"ManogyaDahal/ecom/internal/json"
	"log"
	"net/http"
)

type handlers struct{ 
	orderService service
}

func NewHandler(service service) *handlers{ 
	return &handlers{ 
		orderService: service,
	}
}

// extract, validate data from request and write the data into response
func (h *handlers) PlaceOrder (w http.ResponseWriter, r *http.Request){ 
	// read the payload
	var tempOrder createOrderParams
	if err := json.Read(r, &tempOrder); err != nil { 
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// perform certain operations with the service and recieve data
	createdOrder, err := h.orderService.PlaceOrder(r.Context(), tempOrder)
	if err != nil { 
		log.Println(err)
		if err == ErrProductNotFound { 
			http.Error(w, err.Error(), http.StatusNotFound)
			return 
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// wirte the data to the user
	json.Write(w, http.StatusCreated, createdOrder)	
}
