package model

import "time"

type OrderData struct {
	OrderID  string  `json:"order_id,omitempty"`
	UserID   string  `json:"user_id,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type OrderCreatedEvent struct {
	EventID    string    `json:"event_id,omitempty"`
	EventType  string    `json:"event_type,omitempty"`
	OccurredAt time.Time `json:"occurred_at,omitempty"`
	Data       OrderData `json:"data,omitempty"`
}
