package main

var ValidCylinderTypes = []string{"5kg", "15kg", "50kg"}

var ValidOwnerTypes = []string{"自有", "客户自带"}

const (
	StatusPending = "待充"
	StatusFull    = "满瓶"
	StatusInUse   = "使用中"
)

type Cylinder struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Owner     string `json:"owner"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type FillRecord struct {
	CylinderID string `json:"cylinder_id"`
	Date       string `json:"date"`
	Operator   string `json:"operator"`
}

type LendRecord struct {
	CylinderID string `json:"cylinder_id"`
	Customer   string `json:"customer"`
	Phone      string `json:"phone"`
	Date       string `json:"date"`
}

type ReturnRecord struct {
	CylinderID string `json:"cylinder_id"`
	Date       string `json:"date"`
}

type DataStore struct {
	Cylinders     []Cylinder     `json:"cylinders"`
	FillRecords   []FillRecord   `json:"fill_records"`
	LendRecords   []LendRecord   `json:"lend_records"`
	ReturnRecords []ReturnRecord `json:"return_records"`
}
