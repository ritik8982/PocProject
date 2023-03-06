package models

// structure for Lead
type Lead struct {
	UniqueId    int    `json:"unique_id" bson:"unique_id" validate:"omitempty"` //ye field me user se nhi lena chahte user ko dena chahta hu,kud se like auto increment se unique id,but ye field user ko dekh ki uski id ye hai, so wo update ke time bata paye
	FirstName   string `json:"first_name" bson:"first_name" validate:"required"`
	LastName    string `json:"last_name" bson:"last_name"`
	Email       string `json:"email" bson:"email" validate:"required,email"`
	PhoneNo     int    `json:"phone_no" bson:"phone_no" validate:"required,gt=999999999"`
	CompanyName string `json:"company_name" bson:"company_name"`
	Country     string `json:"country" bson:"country" validate:"required"`
}
