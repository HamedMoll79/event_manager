package entity

import "time"

type actionType int

const (
	create actionType = iota + 1
	update
	delete
)

func (action actionType) String() string {
	switch action {
	case create:
		return "create"
	case update:
		return "update"
	case delete:
		return "delete"
	}

	return ""
}

type Order struct {
	ID          string     `json:"id"`
	StoreID     string     `json:"store_id"`
	MetaData    string     `json:"meta_data"`
	Type        actionType `json:"action_type"`
	IsPublished bool       `json:"is_published"`
	Amount      float64    `json:"amount"`
	UserID      string     `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
