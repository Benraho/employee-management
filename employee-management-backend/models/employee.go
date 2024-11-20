package models

type Employee struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Position    string `json:"position"`
	Department  string `json:"department"`
	DateOfHire  string `json:"dateOfHire"`
}
