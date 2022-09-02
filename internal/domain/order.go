package domain

import (
	"errors"
	"time"
)

type status int8

const (
	Unconfirmed status = iota + 1
	InProgress
	Ready
	InTransit
	Delivered
)

func (s status) String() string {
	switch s {
	case Unconfirmed:
		return "unconfirmed"
	case InProgress:
		return "in_progress"
	case Ready:
		return "ready"
	case InTransit:
		return "in_transit"
	case Delivered:
		return "delivered"
	}
	return "unknown"
}

type OrderModel struct {
	Id     string    `json:"id"`
	Price  uint64    `json:"price"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

func (o OrderModel) Validate() error {
	if err := o.ValidateStatus(); err != nil {
		return err
	}
	return nil
}

func (o OrderModel) ValidateStatus() error {
	switch o.Status {
	case Unconfirmed.String():
		return nil
	case InProgress.String():
		return nil
	case Ready.String():
		return nil
	case InTransit.String():
		return nil
	case Delivered.String():
		return nil
	default:
		return errors.New("invalid order status")
	}
}
