package types
import(
	"context"
	"ManogyaDahal/oms/services/common/genproto/orders"
) 
type OrderService interface { 
	CreateOrder(context.Context, *orders.Order) error
}
