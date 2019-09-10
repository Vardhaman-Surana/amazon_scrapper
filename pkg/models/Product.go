package models

import "gopkg.in/gorp.v1"

const(
	ProductTableName="products"
)

type Product struct{
	BaseModel
	URL string
	Title string
	Price float32
	CompanyName string
	Status int
}

func(p *Product)PreInsert(s gorp.SqlExecutor) error{
	p.Deleted=0
	p.Status=0
	return nil
}
