package models

import (
	"time"
)

type RegistryDate struct {
	time.Time
}

type FinanceRegistry struct {
	ID             string       `json:"id,omitempty"`
	Date           RegistryDate `json:"date,omitempty"`
	Category       string       `json:"category,omitempty"`
	Title          string       `json:"title,omitempty"`
	Amount         float64      `json:"amount,omitempty"`
	MyCategory     string       `json:"my_category,omitempty"`
	FamilyCategory string       `json:"family_category,omitempty"`
	Accuracy 	   float64      `json:"accuracy,omitempty"`
	Approved	   bool			`json:"approved,default false"`
}

const layout = "2006-01-02"

func (t *RegistryDate) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == `null` {
		*t = RegistryDate{}
		return
	}
	t.Time, err = time.Parse(layout, string(b))
	return
}
