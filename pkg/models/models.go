package models

import (
	"time"
)

type Event struct {
	ProductID string    `json:"product_id"`
	Timestamp time.Time `json:"timestamp`
}

type ManufacturingEvent struct {
	Event
	Manufacturer string `json:"manufacturer`
	Facility     string `json:"facility"`
}

type ShippingEvent struct {
	Event
	Transporter  string `json:"transporter`
	PrevFacility string `json:"prev_facility"`
	CurrFacility string `json:"curr_facility"`
}

type DeliveryEvent struct {
	Event
	Recipient string `json:"recipient"`
	Address   string `json:"address"`
}

type Block struct {
	PrevHash  string
	Position  string
	Data      interface{}
	Timestamp time.Time
	Hash      string
}

type Blockchain struct {
	blocks []*Block
}
