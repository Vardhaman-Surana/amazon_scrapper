package models

type Product struct{
	BaseModel
	LinkID int64
	ProductTitle string
	Price float32
	CompanyName string
}
