package structure

import (
	//"math/rand"
	"github.com/google/uuid"
)

const (
	OrderNew       = "NEW_ORDER"
	OrderPickedUp  = "PICKED_UP"
	OrderDelivered = "DELIVERED"
	OrderCancelled = "CANCELLED"
)

type Order struct {
	Id       uuid.UUID
	Status   string
	Receiver *Customer
	Sender   *Customer
}

func NewOrder(reciver, sender *Customer) Order {
	return Order{
		//Id:       rand.Int(), //We did this way but it might be dangerous
		Id:       uuid.New(),
		Status:   OrderNew,
		Receiver: reciver,
		Sender:   sender,
	}
}

// // Order ID
// func (order *Order) SetOrderStatus(status string) {
// 	order.Status = status
// }

// PickedUp Yeni gelmiş siparişlerin teslimat adımına geçmesini sağlar
func (order *Order) PickedUp() {
	order.Status = OrderPickedUp
}

// Delivered teslim edilme işlemini eğer sipariş yola çıkmışsa gerçekleştirir
func (order *Order) Delivered() {
	order.Status = OrderDelivered
}

// Cancelled iptal işlemini eğer teslim edilmediyse gerçekleştirir
func (order *Order) Cancelled() {
	order.Status = OrderCancelled
}
