package hte

import (
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// OrderManager coordinates during negotitation/payment/delivery flows
type OrderManager interface {
	AddOrder(order types.Order) error
	GetOrder(orderUUID string) (*types.Order, error)
	OrderExists(orderUUID string) bool
	UpdateOrder(order types.Order) error
}
